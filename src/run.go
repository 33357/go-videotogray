package main

import (
	"./lib"
	"./run"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	config,err:= lib.GetConfig()
	if err != nil {
		fmt.Printf(err.Error())
	}

	dir, _ := os.Getwd()
	sourceFolderPath:=path.Join(dir,fmt.Sprintf("%s/%s/source",config.OutPath,getVideoName(config)))
	gifPath:=fmt.Sprintf("%s/%s_w%d_h%d_f%d.gif",sourceFolderPath,getVideoName(config),config.GifWidth,config.GifHeight,config.GifFrame)
	mp3Path:=fmt.Sprintf("%s/%s_%s.mp3",sourceFolderPath,getVideoName(config),config.Mp3Bit)
	binFolderPath:=strings.Replace(sourceFolderPath,"source",fmt.Sprintf("w%d_h%d_f%d_s%d/bin",config.GifWidth,config.GifHeight,config.OutFrame,config.ColorSize),1)
	gpFolderPath:=strings.Replace(binFolderPath,"bin","gp",1)
	gvFolderPath:=strings.Replace(binFolderPath,"bin","gv",1)
	//zipFolderPath:=strings.Replace(binFolderPath,"bin","zip",1)

	err=run.VideoToGif(sourceFolderPath,gifPath,mp3Path,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.GifToBin(gifPath,binFolderPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.BinToGp(binFolderPath,gpFolderPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.BinToGv(binFolderPath,gvFolderPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	//err=run.GvToZip(gvFolderPath,zipFolderPath,config)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	fmt.Println("run success!")
}


func getVideoName(config *lib.ConfigInfo) string {
	arr1 :=strings.Split(config.VideoPath, "/")
	arr2 :=strings.Split(arr1[len(arr1)-1],".")
	return arr2[0]
}


