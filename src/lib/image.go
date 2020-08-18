package lib

import (
	"image"
)

func GrayImage(m image.Image,config *ConfigInfo) [] [] uint8 {
	grayArrays := make([][]uint8, config.OutWidth)
	for i := range grayArrays {
		grayArrays[i] = make([]uint8,config.OutHeight)
	}
	for i := 0; i < config.OutWidth; i++ {
		for j := 0; j < config.OutHeight; j++ {
			colorRgb := m.At(i, j)
			_r, _g, _b, _ := colorRgb.RGBA()
			r := _r >> 8
			g := _g >> 8
			b := _b >> 8
			hd:=int(r * 299/1000 +g * 587/1000+ b * 114/1000)
			grayArrays[i][j]=changeColorSize(hd,config.ColorSize)
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

func GrayArrayToByteArray(grayArray [] [] uint8,config *ConfigInfo)[]uint8 {
	var byteArray []uint8
	for i := 0; i < config.OutWidth; i++ {
		byteArray = append(byteArray, grayArray[i]...)
	}
	return byteArray
}

func ByteArrayToGrayArray(byteArray [] byte,config *ConfigInfo)[][]uint8  {
	var grayArrays [][]uint8
	for i := 0; i < config.OutWidth; i++ {
		grayArrays = append(grayArrays, byteArray[i*config.OutHeight:(i+1)*config.OutHeight])
	}
	return grayArrays
}