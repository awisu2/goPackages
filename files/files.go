package files


import(
  "path/filepath"
  "os"
  "strings"
)


/**
 * 階層を潜り、ファイルとパスを取得
 * @param  {[type]} dir string)       ([]string, []string [description]
 * @return {[type]}     [description]
 */
func WalkWithOption(dir string, maxDeep int, skipDirs []string) (files []string, dirs []string) {
  // Walk内で変換処理を走らせたくないので事前取得
  separator := string(filepath.Separator)

  filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
    if err != nil { return err }

    rel, err := filepath.Rel(dir, path)
    if info.Mode().IsDir() {
      // check skipDirs
      for _, skipDir := range skipDirs {
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
      if (maxDeep > 0 && deep >= maxDeep) { return filepath.SkipDir }

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