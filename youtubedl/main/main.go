package main

import (
	_ "fmt"
	_ "go/youtubedl/youtubedlpkg"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {

	videoID := "https://www.youtube.com/watch?v=ocBhK7w0dRY"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create("video2.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
