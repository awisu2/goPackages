package files


import(
  "path/filepath"
  "os"
  "strings"
  "io/ioutil"
)

/**
 * ディレクトリを再帰的に取得するためのConfig
 */
type WalkDirsConfig struct {
  MaxDeep int // 最大探索震度(0: infinity)
  SkipDirs []string // no search directory names
  NoFile bool // ファイルを対象外とする
  NoDir bool // ディレクトリを対象外とする
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
      // 設定がある場合を対象外とする(ディレクトリ)
      if (config.NoDir) {
        return nil
      }

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
        deep = strings.Count(rel, separator)

        // スキップ処理はするが、同階層のディレクトリ情報は取得
        dirs = append(dirs, path)
      }

      // check deep (0: infinity)
      if (config.MaxDeep > 0 && deep >= config.MaxDeep) { return filepath.SkipDir }

      return nil
    }

    // 設定がある場合を対象外とする(ファイル)
    if (config.NoFile) {
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

/**
 * ディレクトリの直下にあるディレクトリをリスト取得
 * filepath.Walkはどうも遅いので、通常はこちらを利用する
 * TODO: 再帰的な取得
 * @param {[type]} dirname string) (files []string, err error [description]
 */
func GetCurrentDirs(dirname string) (dirs []string, err error) {
  fs, err := ioutil.ReadDir(dirname)
  if err != nil {
    return
  }

  for _, f := range fs {
    if (f.IsDir()) {
      path := filepath.Join(dirname, f.Name())
      dirs = append(dirs, path)
    }
  }
  return
}

/**
 * ディレクトリの直下にあるファイルをリスト取得
 * filepath.Walkはどうも遅いので、通常はこちらを利用する
 * TODO: 再帰的な取得
 * @param {[type]} dirname string) (files []string, err error [description]
 */
func GetCurrentFiles(dirname string) (files []string, err error) {
  fs, err := ioutil.ReadDir(dirname)
  if err != nil {
    return
  }

  for _, f := range fs {
    if (!f.IsDir()) {
      path := filepath.Join(dirname, f.Name())
      files = append(files, path)
    }
  }
  return
}