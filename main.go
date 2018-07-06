package main

import (
	"fmt"
	. "github.com/KeluDiao/gotube/api"
	"log"
)

func main() {
	fmt.Println("===> START MAIN")
	//y := youtube.NewYoutube(true)
	//y.DecodeURL("https://www.youtube.com/watch?v=6UoatE5DXJw")
	//y.DecodeURL("https://www.youtube.com/watch?v=4jHNBzTe1iI")
	//y.StartDownload("/home/yukpiz/labo/repos/private/raspberry-video-player/hoge.mp4")

	//vid, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=6UoatE5DXJw")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//formatsFound := vid.Formats
	//search := []string{"720p"}
	//s := make([]interface{}, len(search))
	//for i, v := range search {
	//	s[i] = v
	//}
	//formats := formatsFound.Filter(ytdl.FormatResolutionKey, s)

	//file, _ := os.Create(vid.Title + ".mp4")
	//defer file.Close()
	//err = vid.Download(formats[0], file)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//idList := [...]string{"6UoatE5DXJw"}
	//rep := "Curry_highlights"
	//for _, id := range idList {
	//	vl, err := GetVideoListFromId(id)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	err = vl.Download(rep, "hoge.mp4", "720p", "video/mp4")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}
