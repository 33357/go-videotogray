package main

import (
	"./lib"
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	config,err:= lib.GetConfig()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	sourcePath,gifPath:=getSource(config)
	err=splitAnimatedGIF(sourcePath,gifPath,config)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func getSource(config *lib.ConfigInfo) (sourcePath,gifPath string){
	dir, _ := os.Getwd()
	sourcePath=path.Join(dir,fmt.Sprintf("%s/%s/source",config.OutPath,getVideoName(config)))
	gifPath=fmt.Sprintf("%s/%s_w%d_h%d_f%d.gif",sourcePath,getVideoName(config),config.GifWidth,config.GifHeight,config.GifFrame)
	_, err := os.Stat(sourcePath)
	if err != nil {
		err:=os.MkdirAll(sourcePath,os.ModePerm)
		if err!=nil{
			fmt.Errorf(err.Error())
		}
	}
	_, err = os.Stat(gifPath)
	if err != nil {
		lib.VideoToGIF(path.Join(dir,config.FFMPEGPath),path.Join(dir,config.VideoPath),config.GifWidth,config.GifHeight,config.GifFrame,gifPath)
	}
	fmt.Println("testGif Success")
	mp3Path:=fmt.Sprintf("%s/%s_%s.mp3",sourcePath,getVideoName(config),config.Mp3Bit)
	_, err = os.Stat(mp3Path)
	if err != nil {
		lib.VideoToMP3(path.Join(dir,config.FFMPEGPath),path.Join(dir,config.VideoPath),config.Mp3Bit,mp3Path)
	}
	fmt.Println("testMp3 Success")
	return sourcePath,gifPath
}

func splitAnimatedGIF(sourcePath string,gifPath string,config *lib.ConfigInfo) (err error) {
	f, err := ioutil.ReadFile(gifPath) //读取文件
	if err != nil {
		err = fmt.Errorf(err.Error())
	}
	b := bytes.NewBuffer(f)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = fmt.Errorf("Error while decoding: %s", r)
		}
	}()

	gif, err := gif.DecodeAll(b)
	if err != nil {
		return err
	}
	fmt.Println("Decode GIF Success")

	overpaintImage := image.NewRGBA(image.Rect(0, 0, config.GifWidth, config.GifHeight))
	draw.Draw(overpaintImage, overpaintImage.Bounds(), gif.Image[0], image.Point{}, draw.Src)

	vbPath:=strings.Replace(sourcePath,"source",fmt.Sprintf("%d_%d/vb",config.GifWidth,config.GifHeight),1)

	_, err = os.Stat(vbPath)
	if err != nil {
		err:=os.MkdirAll(vbPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	for i, srcImg := range gif.Image {
			draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.Point{}, draw.Over)
			//if(config.GifHeight!=config.OutHeight||config.GifWidth!=config.OutWidth){
			//	_image:= resize.Resize(uint(config.OutWidth), uint(config.OutHeight), overpaintImage, resize.Lanczos3)
			//	overpaintImage=_image
			//}

			array := hdImage(overpaintImage)

			buf:=ArraytoBuffer1(array,config)
			path:=fmt.Sprintf("%s/%d.vb",vbPath,i)
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			_, err = file.Write(buf.Bytes())
			if err != nil {
				return err
			}
			file.Close()
				
			//break

			file, err = os.Create(fmt.Sprintf("%s/%d.png",vbPath,i))
			if err != nil {
				return err
			}

			err = png.Encode(file, overpaintImage)
			if err != nil {
				return err
			}
			file.Close()

			fmt.Println("save:"+path)
	}
	return nil
}

//图片灰化处理
func hdImage(m image.Image) [] int {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	array := make([] int, dx*dy)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			_r, _g, _b, _ := colorRgb.RGBA()
			r := _r >> 8
			g := _g >> 8
			b := _b >> 8
			hd:=int(r * 299/1000 +g * 587/1000+ b * 114/1000)
			array[i+j*dx]=changeColorSize(hd, 25)
		}
	}
	return array
}

func ArraytoBuffer2(array [] int,config *lib.ConfigInfo) *bytes.Buffer{
	code,KeyCodeMap:=huffman(array)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("width:%d,height:%d,",config.OutWidth,config.OutHeight))
	for key, value := range KeyCodeMap {
		buf.WriteString(fmt.Sprintf("%s:%d,",value,key))
	}
	buf.WriteString("//")
	length:= len(code)
	for i:=0;;i+=8 {
		if i+8<=length {
			bin,_:=strconv.ParseUint(code[i:i+8],2,8)
			binary.Write(&buf, binary.LittleEndian, uint8(bin))
		}else{
			if i<length {
				bin,_:=strconv.ParseUint(code[i:length],2,8)
				binary.Write(&buf, binary.LittleEndian, uint8(bin))
			}
			break
		}
	}
	return &buf
}

func ArraytoBuffer1(arr [] int,config *lib.ConfigInfo) *bytes.Buffer{
	buf := new(bytes.Buffer)
	for _,value := range arr{
		err := binary.Write(buf, binary.LittleEndian, uint8(value))
		if err != nil {
			fmt.Println("binary.Read failed:", err)
		}
	}
	return buf
}

func changeColorSize(gray int, size int) int {
	if gray<255 {
		i := 0
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

func getVideoName(config *lib.ConfigInfo) string {
	arr1 :=strings.Split(config.VideoPath, "/")
	arr2 :=strings.Split(arr1[len(arr1)-1],".")
	return arr2[0]
}

func huffman(array [] int) (string,map[int]string){
	KeyCodeMap:=lib.GetHuffmanMap(array)
	var buf bytes.Buffer
	for _,value := range array {
		buf.WriteString(KeyCodeMap[value])
	}
	return buf.String(),KeyCodeMap
}
