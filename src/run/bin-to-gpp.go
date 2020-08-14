package run

import (
	"../lib"
	"fmt"
	"io/ioutil"
	"os"
)

func BinToGpp(binFolderPath string,gppFolderPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gppFolderPath)
	if err != nil {
		err:=os.MkdirAll(gppFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	for i:=1;;i++ {
		indexb := i
		indexp := indexb + 1
		gppPath := fmt.Sprintf("%s/%d.gpp", gppFolderPath, indexp)
		bin1Path := fmt.Sprintf("%s/%d.bin", binFolderPath, indexb)
		bin2Path := fmt.Sprintf("%s/%d.bin", binFolderPath, indexp)
		byteArray1, err := ioutil.ReadFile(bin1Path)
		if err != nil {
			return err
		}

		var grayArrays1 [][]uint8
		for j := 0; j < config.OutHeight; j++ {
			grayArrays1 = append(grayArrays1, byteArray1[j*config.OutWidth:(j+1)*config.OutWidth])
		}

		byteArray2, err := ioutil.ReadFile(bin2Path)
		if err != nil {
			return err
		}

		var grayArrays2 [][]uint8
		for j := 0; j < config.OutHeight; j++ {
			grayArrays2 = append(grayArrays2, byteArray2[j*config.OutWidth:(j+1)*config.OutWidth])
		}

		array := lib.TranscodeGpp(grayArrays1, grayArrays2, config)
		err = lib.ArraySaveAsBufferFile(array, gppPath)
		if err != nil {
			return err
		}
	}

	fmt.Println("BinToGpp Success")
	return nil
}
