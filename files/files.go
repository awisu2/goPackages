package files


import(
  "path/filepath"
  "os"
  "strings"
)

/**
 * ディレクトリを再帰的に取得するためのConfig
 */
type WalkDirsConfig struct {
  MaxDeep int
  SkipDirs []string
}

/**
 * 階層を潜り、ファイルとパスを取得
 * @param  {[type]} dir string)       ([]string, []string [description]
 * @return {[type]}     [description]
 */
func WalkDirs(dir string, config WalkDirsConfig) (files []string, dirs []string) {
  // Walk内で変換処理を走らせたくないので事前取得
  separator := string(filepath.Separator)

  filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
    if err != nil { return err }

    rel, err := filepath.Rel(dir, path)
    if info.Mode().IsDir() {
      // check skipDirs
      for _, skipDir := range config.SkipDirs {
        if (info.Name() == skipDir) {
          return filepath.SkipDir
        }
      }

      // get deep
      if err != nil { return err }
      deep := 0
      // . はTopディレクトリ
      if (rel != ".") {
        deep = strings.Count(rel, separator) + 1

        // スキップ処理はするが、同階層のディレクトリ情報は取得
        dirs = append(dirs, path)
      }

      // check deep (0: infinity)
      if (config.MaxDeep > 0 && deep >= config.MaxDeep) { return filepath.SkipDir }

      return nil
    }

    files = append(files, path)
    return nil
  })

  return
}

// ファイル名から拡張子を分離し返却
func SepalateExt(path string) (base, ext string){
  ext = filepath.Ext(path)
  base = path[0:len(path)-len(ext)]
  return
}