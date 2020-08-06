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
	err=run.GifToGp(sourcePath,gifPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.GpToGv(config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.GvToZip(sourcePath)
	if err != nil {
		fmt.Printf(err.Error())
	}
}


