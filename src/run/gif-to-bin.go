package run

import (
	"../lib"
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"io/ioutil"
	"os"
)

func GifToBin(gifPath string,binFolderPath string,config *lib.ConfigInfo) error {
	f, err := ioutil.ReadFile(gifPath) //读取文件
	if err != nil {
		err = fmt.Errorf(err.Error())
	}
	b := bytes.NewBuffer(f)

	_, err = os.Stat(binFolderPath)
	if err != nil {
		err:=os.MkdirAll(binFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}

	gif, err := gif.DecodeAll(b)
	if err != nil {
		return err
	}
	fmt.Println("Decode GIF Success")
	overpaintImage := image.NewRGBA(image.Rect(0, 0, config.SourceWidth, config.SourceHeight))
	draw.Draw(overpaintImage, overpaintImage.Bounds(), gif.Image[0], image.Point{}, draw.Src)

	for i, srcImg := range gif.Image {
		binPath:=fmt.Sprintf("%s/%d.bin",binFolderPath,i+1)
		_, err = os.Stat(binPath)
		if err == nil {
			continue
		}
		draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.Point{}, draw.Over)
		grayArrays := lib.GrayImage(overpaintImage,config)
		var array []uint8
		for _, arr := range grayArrays {
			array=append(array,arr...)
		}
		lib.ArraySaveAsBufferFile(array,binPath)
	}
	fmt.Println("GifToBin Success")
	return nil
}

//图片灰化处理


