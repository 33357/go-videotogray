package run

import (
	"../lib"
	"fmt"
	"io/ioutil"
	"os"
)

func BinToGpp(binFolderPath string,gppFolderPath string,reBinFolderPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gppFolderPath)
	if err != nil {
		err:=os.MkdirAll(gppFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	for i:=1;;i+=config.MaxBPageNum+1 {
		reBinIndex := i
		binIndex := reBinIndex + config.MaxBPageNum+1
		gppPath := fmt.Sprintf("%s/%d.gpp", gppFolderPath, binIndex)
		_, err = os.Stat(gppPath)
		if err == nil {
			continue
		}
		reBinPath := fmt.Sprintf("%s/%d.bin", reBinFolderPath, reBinIndex)
		binPath := fmt.Sprintf("%s/%d.bin", binFolderPath, binIndex)
		reBinPath2 := fmt.Sprintf("%s/%d.bin", reBinFolderPath, binIndex)
		reBinByteArray, err := ioutil.ReadFile(reBinPath)
		if err != nil {
			break
		}
		reBinGrayArray:=lib.ByteArrayToGrayArray(reBinByteArray,config)

		binByteArray, err := ioutil.ReadFile(binPath)
		if err != nil {
			break
		}
		binGrayArray:=lib.ByteArrayToGrayArray(binByteArray,config)

		byteArray,reGrayArray := lib.TranscodeGpp(reBinGrayArray, binGrayArray, config)
		err = lib.ArraySaveAsBufferFile(byteArray, gppPath)
		if err != nil {
			return err
		}
		reBinArray:=lib.GrayArrayToByteArray(reGrayArray,config)
		err=lib.ArraySaveAsBufferFile(reBinArray,reBinPath2)
		if err != nil {
			return err
		}
	}

	fmt.Println("BinToGpp Success")
	return nil
}
