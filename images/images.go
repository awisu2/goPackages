package images

import(
  "os"
  "image"
  "image/jpeg"
  "image/gif"
  "image/png"
  "log"
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
 * Decodeは同じストリーム内で行うとエラーになるため関数に分ける
 */
func DecodeConfig(src string) (config image.Config, format Format, err error) {
  file, err := os.Open(src)
  if err != nil { return }
  defer file.Close() // latest close

  config, formatDecode, err := image.DecodeConfig(file)
  if err != nil { return }
  format = GetFormatByWord(formatDecode)

  return
}

/**
 * 拡張子を確認し返却
 * 頭に.付きでチェック
 */
func CanConvertExt (ext string) bool {
  for _, format := range FormatList() {
    if ext == "." + format.String() {
      return true
    }
  }
  return false
}

type ConvertOption struct {
  MaxHeight int
  Format Format
}

func Converte(src, dist string, option ConvertOption) (err error) {
  config, format, err := DecodeConfig(src)
  if err != nil { return }

  // 指定がフォーマットを取得
  if option.Format != Unknown {
   format = option.Format
  }

  // get file
  file, err := os.Open(src)
  if err != nil { return }
  defer file.Close() // latest close

  // Decode
  img, _, err := image.Decode(file)
  if err != nil { return }

  // Resize
  if (option.MaxHeight > 0 && config.Height >= option.MaxHeight) {
   // calc size
    rate := float64(option.MaxHeight) / float64(config.Height)
    width := uint(float64(config.Width) * rate)
    height := uint(float64(config.Height) * rate)

    // set image
    img = resize.Resize(width, height, img, resize.Lanczos3)
  }

  err = Save(dist, format, img)
  if err != nil { return }

  return
}

func Save(dist string, format Format, img image.Image) (err error) {
  // create file
  out, err := os.Create(dist)
  if err != nil { return }
  defer out.Close()

  // save
  switch format {
  case Jpeg, Jpg:
    opts := &jpeg.Options{Quality: 70}
    jpeg.Encode(out, img, opts)
  case Png:
    png.Encode(out, img)
  case Gif:
    // TODO: 未完成
    opts := &gif.Options{NumColors: 256}
    gif.Encode(out, img, opts)
  default:
    log.Print("no format")
  }

  return
}