package main
import (
	"fmt"
	"os/exec"
	"strconv"
)

//video转gif
func VideoToGIF(ffmpegPath string,videoPath string,width int,height int,frame int,gifPath string ){
	cmd := exec.Command(ffmpegPath, "-i",videoPath, "-s", fmt.Sprintf("%d*%d", width,height),"-r",strconv.Itoa(frame),gifPath)
	// 执行命令，并返回结果
	output,err := cmd.Output()
	if err != nil {
		panic(err)
	}
	// 因为结果是字节数组，需要转换成string
	fmt.Println(string(output))
}
//video转mp3
func VideoToMP3(ffmpegPath string,videoPath string,mp3Bit int ,mp3Path string){
	cmd := exec.Command(ffmpegPath,"-i", videoPath, "-vn", "-acodec","libmp3lame","-ac", "2", "-ab", strconv.Itoa(mp3Bit) ,"-ar" ,"48000", mp3Path)
	// 执行命令，并返回结果
	output,err := cmd.Output()
	if err != nil {
		panic(err)
	}
	// 因为结果是字节数组，需要转换成string
	fmt.Println(string(output))
}
