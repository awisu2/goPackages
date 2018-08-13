package ffmpegdecode

import (
	"encoding/json"
	"os/exec"
)

type Stream struct {
	Index           int
	Codec_type      string
	Codec_name      string
	Codec_long_name string
	Width           int
	Height          int
	Coded_width     int
	Coded_height    int
	Bit_rate        string
}

type Streams struct {
	Streams []Stream
}

const (
	CODEC_TYPE_VIDEO = "video"
)

func Decode(src string) (streams Streams, err error) {
	// get stream info
	out, err := exec.Command("ffprobe", "-i", src, "-show_streams", "-print_format", "json").Output()
	if err != nil {
		return streams, err
	}

	if err := json.Unmarshal(out, &streams); err != nil {
		return streams, err
	}

	return
}
