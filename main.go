package main

import (
	"fmt"
	"github.com/jleight/omxplayer"
	"github.com/otium/ytdl"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	fmt.Println("===> START MAIN")
	url := "https://www.youtube.com/watch?v=6UoatE5DXJw"
	fmt.Println("Download: " + url)
	filename := download(url)

	dir, _ := filepath.Abs(filepath.Dir(filename))
	play(filepath.Join(dir, filename))
}

func download(url string) string {
	vid, _ := ytdl.GetVideoInfo(url)

	var format ytdl.Format
	var targetIndex int
	for i, f := range vid.Formats {
		if f.Extension == "mp4" {
			format = f
			targetIndex = i
			break
		}
		fmt.Printf("[%d, %s, %s, %s, %s]\n", i, f.Extension, f.AudioEncoding, f.Resolution, strconv.Itoa(f.AudioBitrate))
	}

	fmt.Printf("Target ===> [%d, %s, %s, %s, %s]\n", targetIndex, format.Extension, format.AudioEncoding, format.Resolution, strconv.Itoa(format.AudioBitrate))
	file, err := os.Create(vid.ID + "." + format.Extension)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	err = vid.Download(format, file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file.Name()
}

func play(file string) {
	fmt.Println("Playing: " + file)
	player, err := omxplayer.New(file, "--no-osd")
	fmt.Println(player)
	fmt.Println(err)
}
