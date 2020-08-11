package run

import (
	"../lib"
	"fmt"
	"io/ioutil"
	"os"
)

func BinToGp(binFolderPath string,gpFolderPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gpFolderPath)
	if err != nil {
		err:=os.MkdirAll(gpFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	for i:=0;;i++{
		gpPath:=fmt.Sprintf("%s/%d.gp",gpFolderPath,i)
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
	fmt.Println("BinToGp Success")
	return nil
}
