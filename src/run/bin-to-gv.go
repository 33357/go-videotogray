package run

import (
	"../lib"
	"fmt"
	"io/ioutil"
	"os"
)

func BinToGv(binPath string,gvPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gvPath)
	if err != nil {
		err:=os.MkdirAll(gvPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}

	var arr [] [] uint8
	for i:=0;;i++{
		path:=fmt.Sprintf("%s/%d.bin",binPath,i)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			break
		}
		arr= append(arr, data)
		if len(arr)>config.GvSeconds*config.OutFrame {
			array:=lib.TranscodeGV(arr,gvPath,config)
			lib.ArraySaveAsBufferFile(array,fmt.Sprintf("%s/%d.gv",gvPath,i))
			arr = [] [] uint8{}
		}
	}

	if len(arr)!=0 {
		lib.TranscodeGV(arr,gvPath,config)
	}
	return nil
}