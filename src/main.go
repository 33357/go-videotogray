package main
import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

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
		array:=hdImage(overpaintImage)

		file, err := os.Create(fmt.Sprintf("%s%d%s", outPath, i, ".vb"))
		if err != nil {
			return err
		}
		str:=""
		for _,v := range array{
			str+=strconv.Itoa(int(v))+","
		}

		_, err = file.WriteString(str)

		if err != nil {
			return err
		}
		file.Close()
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

func main() {
	source := "./image/yjm_w1020_h300_f12.gif" //输入图片
	target := "./image/png/"                   //输出图片

	ff, _ := ioutil.ReadFile(source) //读取文件
	bbb := bytes.NewBuffer(ff)
	SplitAnimatedGIF(bbb,target)
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
			r, g, b, _ := colorRgb.RGBA()
			array[i*(j+1)]=r * 299/1000 + g * 587/1000+ b * 114/1000
		}
	}
	return array
}
