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
	for i:=0;;i++{
		gpPath:=fmt.Sprintf("%s/%d.gip",gipFolderPath,i)
		_, err = os.Stat(gpPath)
		if err == nil {
			continue
		}

		binPath:=fmt.Sprintf("%s/%d.bin",binFolderPath,i)
		byteArray, err := ioutil.ReadFile(binPath)
		if err != nil {
			break
		}

		var grayArrays [][]uint8
		for j:=0;j<config.OutHeight;j++ {
			grayArrays=append(grayArrays,byteArray[j*config.OutWidth:(j+1)*config.OutWidth])
		}
		array:=lib.TranscodeGP(grayArrays,config)
		err=lib.ArraySaveAsBufferFile(array,gpPath)
		if err != nil {
			return err
		}
	}
	fmt.Println("BinToGip Success")
	return nil
}