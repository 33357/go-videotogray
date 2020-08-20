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
	index:=1
	for i:=1;;i++{
		path:=fmt.Sprintf("%s/%d.bin",binFolderPath,i)
		byteArray, err := ioutil.ReadFile(path)
		if err != nil {
			break
		}
		var grayArray=lib.ByteArrayToGrayArray(byteArray,config)
		arrs= append(arrs, grayArray)
		if i%(config.GvSeconds*config.OutFrame)==0 {
			array:=lib.TranscodeGV(arrs,config)
			lib.ArraySaveAsBufferFile(array,fmt.Sprintf("%s/%d.gv",gvFolderPath,index))
			index++
			arrs = [] [] [] uint8{}
		}
	}

	if len(arrs)!=0 {
		array:=lib.TranscodeGV(arrs,config)
		lib.ArraySaveAsBufferFile(array,fmt.Sprintf("%s/%d.gv",gvFolderPath,index))
	}
	return nil
}