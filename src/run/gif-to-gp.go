package run

import (
	"../lib"
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

func GifToGp(sourcePath string,gifPath string,config *lib.ConfigInfo) (string,error) {
	f, err := ioutil.ReadFile(gifPath) //读取文件
	if err != nil {
		err = fmt.Errorf(err.Error())
	}
	b := bytes.NewBuffer(f)

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println(r)
	//		err = fmt.Errorf("Error while decoding: %s", r)
	//	}
	//}()

	gpPath:=strings.Replace(sourcePath,"source",fmt.Sprintf("%d_%d/gp",config.GifWidth,config.GifHeight),1)
	_, err = os.Stat(gpPath)
	if err != nil {
		err:=os.MkdirAll(gpPath,os.ModePerm)
		if err!=nil{
			return "",err
		}
	}

	gif, err := gif.DecodeAll(b)
	if err != nil {
		return "",err
	}
	fmt.Println("Decode GIF Success")
	overpaintImage := image.NewRGBA(image.Rect(0, 0, config.GifWidth, config.GifHeight))
	draw.Draw(overpaintImage, overpaintImage.Bounds(), gif.Image[0], image.Point{}, draw.Src)

	for i, srcImg := range gif.Image {
		draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.Point{}, draw.Over)
		//if(config.GifHeight!=config.OutHeight||config.GifWidth!=config.OutWidth){
		//	_image:= resize.Resize(uint(config.OutWidth), uint(config.OutHeight), overpaintImage, resize.Lanczos3)
		//	overpaintImage=_image
		//}

		grayArrays := hdImage(overpaintImage,config)
		_array:=lib.TranscodeGP(grayArrays,config)

		buf:=ArraytoBuffer(_array)

		path:=fmt.Sprintf("%s/%d.gp",gpPath,i)
		file, err := os.Create(path)
		if err != nil {
			return "",err
		}
		_, err = file.Write(buf.Bytes())
		if err != nil {
			return "",err
		}
		file.Close()

		file, err = os.Create(fmt.Sprintf("%s/%d.png",gpPath,i))
		if err != nil {
			return "",err
		}

		err = png.Encode(file, overpaintImage)
		if err != nil {
			return "",err
		}
		file.Close()

		fmt.Println("save:"+path)
		//return "",nil
	}
	fmt.Println("GifToGp Success")
	return gpPath,nil
}

//图片灰化处理
func hdImage(m image.Image,config *lib.ConfigInfo) [] [] uint8 {
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
			grayArrays[i][j]=changeColorSize(hd, config.ColorSize)
		}
	}
	return grayArrays
}

func ArraytoBuffer(arr [] uint8) *bytes.Buffer{
	buf := new(bytes.Buffer)
	for _,value := range arr{
		err := binary.Write(buf, binary.LittleEndian,value)
		if err != nil {
			fmt.Println("binary.Read failed:", err)
		}
	}
	return buf
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

func getVideoName(config *lib.ConfigInfo) string {
	arr1 :=strings.Split(config.VideoPath, "/")
	arr2 :=strings.Split(arr1[len(arr1)-1],".")
	return arr2[0]
}


//func ArraytoBuffer2(array [] int,config *lib.ConfigInfo) *bytes.Buffer{
//	code,KeyCodeMap:=huffman(array)
//	var buf bytes.Buffer
//	buf.WriteString(fmt.Sprintf("width:%d,height:%d,",config.OutWidth,config.OutHeight))
//	for key, value := range KeyCodeMap {
//		buf.WriteString(fmt.Sprintf("%s:%d,",value,key))
//	}
//	buf.WriteString("//")
//	length:= len(code)
//	for i:=0;;i+=8 {
//		if i+8<=length {
//			bin,_:=strconv.ParseUint(code[i:i+8],2,8)
//			binary.Write(&buf, binary.LittleEndian, uint8(bin))
//		}else{
//			if i<length {
//				bin,_:=strconv.ParseUint(code[i:length],2,8)
//				binary.Write(&buf, binary.LittleEndian, uint8(bin))
//			}
//			break
//		}
//	}
//	return &buf
//}

//
//func huffman(array [] int) (string,map[int]string){
//	KeyCodeMap:=lib.GetHuffmanMap(array)
//	var buf bytes.Buffer
//	for _,value := range array {
//		buf.WriteString(KeyCodeMap[value])
//	}
//	return buf.String(),KeyCodeMap
//}
