package lib

import (
	"image"
)

func GrayImage(m image.Image,config *ConfigInfo) [] [] uint8 {
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
