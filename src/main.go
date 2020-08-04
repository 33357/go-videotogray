package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	source := "./image/yjm.gif" //输入图片
	target := "./image/vb/"                   //输出图片

	ff, _ := ioutil.ReadFile(source) //读取文件
	bbb := bytes.NewBuffer(ff)
	SplitAnimatedGIF(bbb,target)
}

func SplitAnimatedGIF(reader io.Reader,outPath string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error while decoding: %s", r)
		}
	}()
	gif, err := gif.DecodeAll(reader)
	if err != nil {
		return err
	}

	imgWidth, imgHeight := getGifDimensions(gif)
	overpaintImage := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(overpaintImage, overpaintImage.Bounds(), gif.Image[0], image.Point{}, draw.Src)

	for i, srcImg := range gif.Image {
			draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.Point{}, draw.Over)
			array := hdImage(overpaintImage)
			//fmt.Println(len(array))
			//for j:=0;j<len(array);j++{
			//	if array[j]!=0 {
			//		fmt.Println(j,array[j])
			//	}
			//}
			//fmt.Println(array[19403])

			file, err := os.Create(fmt.Sprintf("%s%d%s", outPath, i, ".vb"))
			if err != nil {
				return err
			}
			buf := ArraytoBuffer(array)
			_, err = file.Write(buf.Bytes())

			if err != nil {
				return err
			}
			file.Close()

			//file, err = os.Create(fmt.Sprintf("%s%d%s", outPath, i, ".png"))
			//if err != nil {
			//	return err
			//}
			//
			//err = png.Encode(file, overpaintImage)
			//if err != nil {
			//	return err
			//}
			println(i)
	}
	return nil
}

func getGifDimensions(gif *gif.GIF) (x, y int) {
	var lowestX int
	var lowestY int
	var highestX int
	var highestY int

	for _, img := range gif.Image {
		if img.Rect.Min.X < lowestX {
			lowestX = img.Rect.Min.X
		}
		if img.Rect.Min.Y < lowestY {
			lowestY = img.Rect.Min.Y
		}
		if img.Rect.Max.X > highestX {
			highestX = img.Rect.Max.X
		}
		if img.Rect.Max.Y > highestY {
			highestY = img.Rect.Max.Y
		}
	}

	return highestX - lowestX, highestY - lowestY
}
//图片灰化处理
func hdImage(m image.Image) [] uint32 {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	array := make([] uint32, dx*dy)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			_r, _g, _b, _ := colorRgb.RGBA()
			r := _r >> 8
			g := _g >> 8
			b := _b >> 8
			hd:=r * 299/1000 +g * 587/1000+ b * 114/1000
			array[i+j*dx]=changeColorSize(hd, 25)
			//if array[i+j*dx] != 0 {
			//	fmt.Println(i,j,i+j*dx,array[i+j*dx])
			//}
		}
	}
	return array
}

func ArraytoBuffer(arr [] uint32) *bytes.Buffer{
	buf := new(bytes.Buffer)
	for _,value := range arr{
		err := binary.Write(buf, binary.LittleEndian, uint8(value))
		if err != nil {
			fmt.Println("binary.Read failed:", err)
		}
	}
	return buf
}

func changeColorSize(gray uint32, size uint32) uint32 {
	if gray<255 {
		var i uint32 = 0
		g := gray*size/255
		for ;i<size;i++{
			if g>=i && g<i+1 {
				return i
			}
		}
		return 0
	}else{
		return size-1
	}
}