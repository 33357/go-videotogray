package run

import (
	"../lib"
	"fmt"
	"os"
	"path"
)

func VideoToGif(config *lib.ConfigInfo) (sourcePath,gifPath string){
	dir, _ := os.Getwd()
	sourcePath=path.Join(dir,fmt.Sprintf("%s/%s/source",config.OutPath,getVideoName(config)))
	gifPath=fmt.Sprintf("%s/%s_w%d_h%d_f%d.gif",sourcePath,getVideoName(config),config.GifWidth,config.GifHeight,config.GifFrame)
	_, err := os.Stat(sourcePath)
	if err != nil {
		err:=os.MkdirAll(sourcePath,os.ModePerm)
		if err!=nil{
			fmt.Errorf(err.Error())
		}
	}
	_, err = os.Stat(gifPath)
	if err != nil {
		lib.VideoToGIF(path.Join(dir,config.FFMPEGPath),path.Join(dir,config.VideoPath),config.GifWidth,config.GifHeight,config.GifFrame,gifPath)
	}
	fmt.Println("testGif Success")
	mp3Path:=fmt.Sprintf("%s/%s_%s.mp3",sourcePath,getVideoName(config),config.Mp3Bit)
	_, err = os.Stat(mp3Path)
	if err != nil {
		lib.VideoToMP3(path.Join(dir,config.FFMPEGPath),path.Join(dir,config.VideoPath),config.Mp3Bit,mp3Path)
	}
	fmt.Println("testMp3 Success")
	return sourcePath,gifPath
}
