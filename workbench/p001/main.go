package main

import (
	"fmt"

	"github.com/qrtc/ffmpeg-dev-go"
)

func main() {
	fmt.Println(ffmpeg.AvVersionInfo())
}
