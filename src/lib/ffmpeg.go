package lib

import (
	"fmt"
	"os/exec"
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
func VideoToMP3(ffmpegPath string,videoPath string,mp3Bit string ,mp3Path string){
	cmd := exec.Command(ffmpegPath,"-i", videoPath, "-vn", "-acodec","libmp3lame","-ac", "2", "-ab", mp3Bit ,"-ar" ,"48000", mp3Path)
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
