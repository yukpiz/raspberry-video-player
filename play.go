package main

import (
	"fmt"
	"github.com/yukpiz/omxplayer"
	"path/filepath"
	"os"
	"time"
)

func main() {
	fmt.Println("===> START MAIN")
	filename := "6UoatE5DXJw.mp4"
	dir, _ := filepath.Abs(filepath.Dir(filename))
	file := filepath.Join(dir, filename)
	omxplayer.SetUser("root", "/root")
	player, err := omxplayer.New(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println("===> WaitForReady()")
	//player.WaitForReady()
	//fmt.Println("===> PlayPause()")
	//err = player.PlayPause()

	//time.Sleep(5 * time.Second)
	//fmt.Println("===> ShowSubtitles()")
	//err = player.ShowSubtitles()

	//time.Sleep(5 * time.Second)
	//fmt.Println("===> Quit()")
	//err = player.Quit()
	fmt.Println(player.IsRunning())
	time.Sleep(60 * time.Second)
}
