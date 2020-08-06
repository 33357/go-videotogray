package main

import (
	"./lib"
	"./run"
	"fmt"
)

func main() {
	config,err:= lib.GetConfig()
	if err != nil {
		fmt.Printf(err.Error())
	}
	sourcePath,gifPath:=run.VideoToGif(config)
	_,err=run.GifToGp(sourcePath,gifPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	//gvPath,err:=run.GpToGv(gpPath,config)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	//err=run.GvToZip(gvPath,config)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
}


