package words

import (
  "strings"
  "regexp"
  "strconv"
  "fmt"
)

/**
 * 全部を同じ長さの0フィルにする
 * 一致した文字列を変換する関数はなく
 * 同じ文字が連続した場合などに対処するため直指定での変換を採用
 */
func ZeroFillInStr(str string, length int) string {
  // 文字列から数値を抜粋
  reg := regexp.MustCompile(`[0-9]+`)
  finds := reg.FindAllString(str, -1)

  // Replacer用に変換配列を作成
  replaces := []string{}
  format := `%0` + strconv.Itoa(length) + `d`
  for _, find := range finds {
    replaces = append(replaces, find)
    i, _ := strconv.Atoi(find)
    replaces = append(replaces, fmt.Sprintf(format, i))
  }
  r := strings.NewReplacer(replaces...)

  // 変換
  return r.Replace(str)
}

func MaxLengthNumberByStrs(strs []string) (length int) {
  for _, str := range strs {
    tmpLength := MaxLengthNumberByStr(str)
    if length < tmpLength {
      length = tmpLength
    }
  }
  return
}

func MaxLengthNumberByStr(str string) (length int) {
  reg := regexp.MustCompile(`[1-9][0-9]+`)
  for _, find := range reg.FindAllString(str, -1) {
    if length < len(find) {
      length = len(find)
    }
  }
  return
}