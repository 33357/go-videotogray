package run

import (
	"../lib"
	"fmt"
	"io/ioutil"
	"os"
)

func BinToGv(binFolderPath string,gvFolderPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gvFolderPath)
	if err != nil {
		err:=os.MkdirAll(gvFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}

	var arrs [] [] []uint8
	for i:=0;;i++{
		path:=fmt.Sprintf("%s/%d.bin",binFolderPath,i)
		byteArray, err := ioutil.ReadFile(path)
		if err != nil {
			break
		}
		var grayArrays [][]uint8
		for j:=0;j<config.OutHeight;j++ {
			grayArrays=append(grayArrays,byteArray[j*config.OutWidth:(j+1)*config.OutWidth])
		}
		arrs= append(arrs, grayArrays)
		if len(arrs)>config.GvSeconds*config.OutFrame {
			array:=lib.TranscodeGV(arrs,config)
			lib.ArraySaveAsBufferFile(array,fmt.Sprintf("%s/%d.gv",gvFolderPath,i))
			arrs = [] [] [] uint8{}
		}
	}

	if len(arrs)!=0 {
		lib.TranscodeGV(arrs,config)
	}
	return nil
}