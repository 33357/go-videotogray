package run

import (
	"../lib"
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image/png"
	"io/ioutil"
	"os"
)

func PngToBin(pngFolderPath string,binFolderPath string,config *lib.ConfigInfo) error {
	_, err := os.Stat(binFolderPath)
	if err != nil {
		err:=os.MkdirAll(binFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	for i:=1;;i++{
		binPath:=fmt.Sprintf("%s/%d.bin",binFolderPath,i)
		_, err = os.Stat(binPath)
		if err == nil {
			continue
		}
		pngPath:=fmt.Sprintf("%s/%d.png",pngFolderPath,i)
		f, err := ioutil.ReadFile(pngPath) //读取文件
		if err != nil {
			break
		}
		buf := bytes.NewBuffer(f)
		image, err := png.Decode(buf)
		if err != nil {
			return err
		}
		if config.SourceHeight!=config.OutHeight||config.SourceWidth!=config.OutWidth {
			image= resize.Resize(uint(config.OutWidth), uint(config.OutHeight), image, resize.Lanczos3)
		}
		grayArray := lib.GrayImage(image,config)
		byteArray:=lib.GrayArrayToByteArray(grayArray,config)
		lib.ArraySaveAsBufferFile(byteArray,binPath)
		if err != nil {
			return err
		}
	}
	fmt.Println("PngToBin Success")
	return nil
}

