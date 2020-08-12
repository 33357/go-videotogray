package run

import (
	"../lib"
	"fmt"
	"os"
	"path"
)

func VideoToGif(sourceFolderPath string,gifPath string,mp3Path string,pngFolderPath string,config *lib.ConfigInfo) error{
	dir, _ := os.Getwd()
	_, err := os.Stat(sourceFolderPath)
	if err != nil {
		err=os.MkdirAll(sourceFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}

	_, err = os.Stat(gifPath)
	if err != nil {
		lib.VideoToGIF(path.Join(dir,config.FFMPEGPath),path.Join(dir,config.VideoPath),config.GifWidth,config.GifHeight,config.GifFrame,gifPath)
	}
	fmt.Println("testGif Success")

	_, err = os.Stat(mp3Path)
	if err != nil {
		lib.VideoToMP3(path.Join(dir,config.FFMPEGPath),path.Join(dir,config.VideoPath),config.Mp3Bit,mp3Path)
	}
	fmt.Println("testMp3 Success")

	_, err = os.Stat(pngFolderPath)
	if err != nil {
		err=os.MkdirAll(pngFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
		lib.VideoToPNG(path.Join(dir,config.FFMPEGPath),path.Join(dir,config.VideoPath),config.GifWidth,config.GifHeight,config.GifFrame,pngFolderPath)
	}

	fmt.Println("testPng Success")

	fmt.Println("VideoToGif Success")
	return nil
}
