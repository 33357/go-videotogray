package run

import (
"../lib"
"fmt"
"io/ioutil"
"os"
)

func BinToGip(binFolderPath string,gipFolderPath string,reBinFolderPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gipFolderPath)
	if err != nil {
		err=os.MkdirAll(gipFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	_, err = os.Stat(reBinFolderPath)
	if err != nil {
		err=os.MkdirAll(reBinFolderPath,os.ModePerm)
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
		reBinPath:=fmt.Sprintf("%s/%d.bin",reBinFolderPath,i)
		grayArray:=lib.ByteArrayToGrayArray(byteArray,config)
		byteArray,reGrayArray:=lib.TranscodeGip(grayArray,config)
		reBinArray:=lib.GrayArrayToByteArray(reGrayArray,config)
		err=lib.ArraySaveAsBufferFile(byteArray,gipPath)
		if err != nil {
			return err
		}
		err=lib.ArraySaveAsBufferFile(reBinArray,reBinPath)
		if err != nil {
			return err
		}
	}
	fmt.Println("BinToGip Success")
	return nil
}
