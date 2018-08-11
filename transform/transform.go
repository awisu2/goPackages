package transform

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

/**
 * Eucjp2Utf8
 */
func Eucjp2Utf8(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.EUCJP.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}

// UTF-8 から ShiftJIS
func Utf82Sjis(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewEncoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}

// ShiftJIS から UTF-8
func Sjis2Utf8(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}

// UTF-8 から EUC-JP
func Utf82Eucjp(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.EUCJP.NewEncoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}

// type WCODE int

// const (
// 	WCODE_NON WCODE = iota
// 	WCODE_SHIFTJIS
// 	WCODE_UTF8
// )

// func GetEncoding(str string) (WCODE, error) {
// 	body := []byte(str)

// 	// チェック対象のコード
// 	code := WCODE_NON
// 	encodings := []string{"utf-8", "sjis"}
// 	for _, enc := range encodings {
// 		if enc != "" {
// 			ee, _ := charset.Lookup(enc)
// 			if ee == nil {
// 				continue
// 			}

// 			// ここでエラーが起きる場合は、エンコードがあっていない
// 			var buf bytes.Buffer
// 			ic := transform.NewWriter(&buf, ee.NewDecoder())
// 			if _, err := ic.Write(body); err != nil {
// 				continue
// 			}

// 			// クローズ
// 			if err := ic.Close(); err != nil {
// 				continue
// 			}

// 			switch enc {
// 			case "sjis":
// 				code = WCODE_SHIFTJIS
// 			case "utf-8":
// 				code = WCODE_UTF8
// 			}
// 			break
// 		}
// 	}
// 	return code, nil
// }
