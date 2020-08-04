package main

import (
	"github.com/giorgisio/goav/avformat"
	"log"
	"os"
	"path"
)

func main() {
	dir, _ := os.Getwd()
	filename := path.Join(dir,"/video/yjm.mp4")

	// Register all formats and codecs
	avformat.AvRegisterAll()

	// Open video file
	if avformat.AvformatOpenInput(&ctxtFormat, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctxtFormat.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")
		return
	}

	//...

}
