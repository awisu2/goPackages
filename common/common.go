package common

import (
  "log"
  "bufio"
  "fmt"
  "os"
)

/**
 * エラーの場合、ログを吐いてExit
 * @param  {[type]} err error         [description]
 * @return {[type]}     [description]
 */
func LogError(err error) bool{
  if err != nil {
    log.Print(err)
    return true
  }
  return false
}

func Log(v ...interface{}) {
  for _, value := range v {
    log.Printf("Type:%T, Value:%v", value, value)
  }
}

/**
 * 適当に入力待機する関数
 * 終わる処理がないので、テスト用に
 * @return {[type]} [description]
 */
func Wait () {
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      fmt.Println(scanner.Text())
  }
  if scanner.Err() != nil {
      fmt.Println(scanner.Text())
  }
}