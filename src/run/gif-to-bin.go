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
	overpaintImage := image.NewRGBA(image.Rect(0, 0, config.GifWidth, config.GifHeight))
	draw.Draw(overpaintImage, overpaintImage.Bounds(), gif.Image[0], image.Point{}, draw.Src)

	for i, srcImg := range gif.Image {
		binPath:=fmt.Sprintf("%s/%d.bin",binFolderPath,i)
		_, err = os.Stat(binPath)
		if err == nil {
			continue
		}
		draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.Point{}, draw.Over)
		//if(config.GifHeight!=config.OutHeight||config.GifWidth!=config.OutWidth){
		//	_image:= resize.Resize(uint(config.OutWidth), uint(config.OutHeight), overpaintImage, resize.Lanczos3)
		//	overpaintImage=_image
		//}
		grayArrays := hdImage(overpaintImage,config)
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
func hdImage(m image.Image,config *lib.ConfigInfo) [] [] uint8 {
	grayArrays := make([][]uint8, config.OutHeight)
	for i := range grayArrays {
		grayArrays[i] = make([]uint8,config.OutWidth)
	}
	for i := 0; i < config.OutHeight; i++ {
		for j := 0; j < config.OutWidth; j++ {
			colorRgb := m.At(j, i)
			_r, _g, _b, _ := colorRgb.RGBA()
			r := _r >> 8
			g := _g >> 8
			b := _b >> 8
			hd:=int(r * 299/1000 +g * 587/1000+ b * 114/1000)
			grayArrays[i][j]=changeColorSize(hd, config.ColorSize)
		}
	}
	return grayArrays
}

func changeColorSize(gray int, size int) uint8 {
	if gray<255 {
		i := 0
		g := gray*size/255
		for ;i<size;i++{
			if g>=i && g<i+1 {
				return uint8(i)
			}
		}
		return 0
	}else{
		return uint8(size-1)
	}
}

