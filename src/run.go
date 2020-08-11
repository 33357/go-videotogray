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
	sourcePath:=path.Join(dir,fmt.Sprintf("%s/%s/source",config.OutPath,getVideoName(config)))
	gifPath:=fmt.Sprintf("%s/%s_w%d_h%d_f%d.gif",sourcePath,getVideoName(config),config.GifWidth,config.GifHeight,config.GifFrame)
	mp3Path:=fmt.Sprintf("%s/%s_%s.mp3",sourcePath,getVideoName(config),config.Mp3Bit)
	binPath:=strings.Replace(sourcePath,"source",fmt.Sprintf("w%d_h%d_f%d_s%d/bin",config.GifWidth,config.GifHeight,config.OutFrame,config.ColorSize),1)
	gpPath:=strings.Replace(binPath,"bin","gp",1)
	//gvPath:=strings.Replace(gpPath,"bin","gv",1)
	//zipPath:=strings.Replace(gpPath,"bin","zip",1)

	err=run.VideoToGif(sourcePath,gifPath,mp3Path,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.GifToBin(gifPath,binPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.BinToGp(binPath,gpPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	//err=run.BinToGv(binPath,gvPath,config)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	//err=run.GvToZip(gvPath,zipPath,config)
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


