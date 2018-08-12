package images

import (
	"bufio"
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"

	"github.com/nfnt/resize"
)

type Format int

const (
	Unknown Format = iota
	Jpeg
	Jpg
	Png
	Gif
)

func (f Format) String() string {
	switch f {
	case Jpeg:
		return "jpeg"
	case Jpg:
		return "jpg"
	case Png:
		return "png"
	case Gif:
		return "gif"
	default:
		return "Unknown"
	}
}

func FormatList() (formats []Format) {
	return []Format{
		Jpeg,
		Jpg,
		Png,
		Gif}
}

func GetFormatByWord(word string) Format {
	for _, format := range FormatList() {
		if word == format.String() {
			return format
		}
	}
	return Unknown
}

/**
 * 対象ファイルのimage.DecodeConfigの値を返却
 * Decodeは同じストリーム内で行うとエラーになるため関数に分ける
 */
func DecodeConfig(src string) (config image.Config, format Format, err error) {
	file, err := os.Open(src)
	if err != nil {
		return
	}
	defer file.Close() // latest close

	config, formatDecode, err := image.DecodeConfig(file)
	if err != nil {
		return
	}
	format = GetFormatByWord(formatDecode)

	return
}

func DecodeByFile(src string) (image.Image, string, error) {
	// get file
	file, err := os.Open(src)
	if err != nil {
		return nil, "", err
	}
	defer file.Close() // latest close

	// Decode(get image)
	img, name, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}

	return img, name, nil
}

/**
 * 拡張子を確認し返却
 * 頭に.付きでチェック
 */
func CanConvertExt(ext string) bool {
	for _, format := range FormatList() {
		if ext == "."+format.String() {
			return true
		}
	}
	return false
}

type ConvertOption struct {
	MaxHeight int
	Format    Format
}

func Converte(src string, option ConvertOption) (img image.Image, err error) {
	// 処理可能なフォーマットか確認
	_, format, err := DecodeConfig(src)
	if err != nil {
		return
	}
	if option.Format != Unknown {
		format = option.Format
	}

	// get file
	img, _, err = DecodeByFile(src)
	if err != nil {
		return
	}

	// Resize
	size := GetSize(img)
	if option.MaxHeight > 0 && size.Y >= option.MaxHeight {
		img = ResizeByHeight(img, option.MaxHeight)
	}

	// change format
	// save
	b := new(bytes.Buffer)
	writer := bufio.NewWriter(b)
	err = ChangeFormat(writer, format, img)
	if err != nil {
		return
	}

	// 画像ファイルに変換
	reader := bufio.NewReader(b)
	img, _, err = image.Decode(reader)
	if err != nil {
		return
	}

	return
}

func ResizeByHeight(img image.Image, height int) image.Image {
	size := GetSize(img)
	log.Printf("%v, %v", img.Bounds().Min, img.Bounds().Max)
	rate := float64(height) / float64(size.Y)
	_width := uint(float64(size.X) * rate)
	_height := uint(float64(size.Y) * rate)

	// set image
	img = resize.Resize(_width, _height, img, resize.Lanczos3)
	return img
}

// IMAGEオブジェクトからサイズを取得
func GetSize(img image.Image) image.Point {
	point := image.Point{
		X: img.Bounds().Max.X - img.Bounds().Min.X,
		Y: img.Bounds().Max.Y - img.Bounds().Min.Y,
	}
	return point
}

func ChangeFormat(writer io.Writer, format Format, img image.Image) error {
	switch format {
	case Jpeg, Jpg:
		opts := &jpeg.Options{Quality: 70}
		jpeg.Encode(writer, img, opts)
	case Png:
		png.Encode(writer, img)
	case Gif:
		// TODO: 未完成
		opts := &gif.Options{NumColors: 256}
		gif.Encode(writer, img, opts)
	default:
		log.Print("no format")
	}

	return nil
}

func SaveJpegToFile(img image.Image, dest string) error {
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	err = ChangeFormat(f, Jpeg, img)
	if err != nil {
		return err
	}

	return nil
}
