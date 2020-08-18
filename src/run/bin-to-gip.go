package run

import (
"../lib"
"fmt"
"io/ioutil"
"os"
)

func BinToGip(binFolderPath string,gipFolderPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gipFolderPath)
	if err != nil {
		err:=os.MkdirAll(gipFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	for i:=1;;i++{
		gipPath:=fmt.Sprintf("%s/%d.gip",gipFolderPath,i)
		_, err = os.Stat(gipPath)
		if err == nil {
			continue
		}
		binPath:=fmt.Sprintf("%s/%d.bin",binFolderPath,i)
		byteArray, err := ioutil.ReadFile(binPath)
		if err != nil {
			break
		}
		grayArray:=lib.ByteArrayToGrayArray(byteArray,config)
		array:=lib.TranscodeGip(grayArray,config)
		err=lib.ArraySaveAsBufferFile(array,gipPath)
		if err != nil {
			return err
		}
	}
	fmt.Println("BinToGip Success")
	return nil
}
