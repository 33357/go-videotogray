package run

import (
	"../lib"
	"fmt"
	"io/ioutil"
	"os"
)

func BinToGp(binPath string,gpPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gpPath)
	if err != nil {
		err:=os.MkdirAll(gpPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	for i:=0;;i++{
		path:=fmt.Sprintf("%s/%d.bin",binPath,i)
		byteArray, err := ioutil.ReadFile(path)
		if err != nil {
			break
		}

		var grayArrays [][]uint8
		for j:=0;j<config.OutHeight;j++ {
			grayArrays=append(grayArrays,byteArray[j*config.OutWidth:(j+1)*config.OutWidth])
		}
		array:=lib.TranscodeGP(grayArrays,config)
		err=lib.ArraySaveAsBufferFile(array,fmt.Sprintf("%s/%d.gp",gpPath,i))
		if err != nil {
			return err
		}
	}
	fmt.Println("BinToGp Success")
	return nil
}
