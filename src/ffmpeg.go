package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
)

//video转gif
func VideoToGIF(ffmpegPath string,videoPath string,width int,height int,frame int,gifPath string ){
	cmd := exec.Command(ffmpegPath, "-i",videoPath, "-s", fmt.Sprintf("%d*%d", width,height),"-r",strconv.Itoa(frame),gifPath)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		fmt.Print(err)
	}

	if err = cmd.Start(); err != nil {
		fmt.Print(err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		fmt.Print(err)
	}
}
//video转mp3
func VideoToMP3(ffmpegPath string,videoPath string,mp3Bit int ,mp3Path string){
	cmd := exec.Command(ffmpegPath,"select=(gte(t\\,+0))*(isnan(prev_selected_t)+gte(t-prev_selected_t\\,120))","-i", videoPath, "-vn", "-acodec","libmp3lame","-ac", "2", "-ab", strconv.Itoa(mp3Bit) ,"-ar" ,"48000", mp3Path)
	// 执行命令，并返回结果
	output,err := cmd.Output()
	if err != nil {
		panic(err)
	}
	// 因为结果是字节数组，需要转换成string
	fmt.Println(string(output))
}

func main() {
	dir, _ := os.Getwd()
	VideoToGIF(path.Join(dir,"/exe/ffmpeg.exe"),path.Join(dir,"/video/yjm.mp4"),1280,720,12,path.Join(dir,"/image/yjm.gif"))
}
