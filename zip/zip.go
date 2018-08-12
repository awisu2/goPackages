package zip

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/saintfish/chardet"

	"github.com/awisu2/goPackages/transform"
)

func Unzip(src, dest string) ([]string, error) {
	files := []string{}

	r, err := zip.OpenReader(src)
	if err != nil {
		return files, err
	}
	defer r.Close()

	det := chardet.NewTextDetector()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return files, err
		}
		defer rc.Close()

		// ShiftJISの場合,encodeをutf8にする
		fName := f.Name
		res, err := det.DetectBest([]byte(fName))
		switch res.Charset {
		case "Shift_JIS", "windows-1252":
			fName, err = transform.Sjis2Utf8(fName)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		case "EUC-KR", "EUC-JP":
			fName, err = transform.Eucjp2Utf8(fName)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		case "UTF-8":
		// 知らないフォーマットの場合はとりあえずshiftjisから変換させておく
		case "windows-1254", "ISO-8859-6", "Big5", "GB-18030", "ISO-8859-1", "ISO-8859-9", "IBM420_ltr", "IBM424_ltr", "IBM420_rtl":
			fName, err = transform.Sjis2Utf8(fName)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		default:
			fName, err = transform.Sjis2Utf8(fName)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			log.Printf("unknown charset: %s, char: %s, conv: %s", res.Charset, f.Name, fName)
		}

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, fName)
		files = append(files, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return files, err
			}

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return files, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return files, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return files, err
			}

		}
	}

	return files, err
}
