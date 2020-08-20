package run

import (
	"../lib"
	"fmt"
	"io/ioutil"
	"os"
)

func BinToGbp(binFolderPath string,gbpFolderPath string,reBinFolderPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gbpFolderPath)
	if err != nil {
		err:=os.MkdirAll(gbpFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	for i:=1;;i+=config.MaxBPageNum+1 {
		reBinIndex1 := i
		reBinIndex2 := reBinIndex1 + config.MaxBPageNum+1
		var gbpPath =fmt.Sprintf("%s/%d-%d.gbp", gbpFolderPath,reBinIndex1+1,reBinIndex2-1)
		_, err = os.Stat(gbpPath)
		if err == nil {
			continue
		}
		var bPageGrayArrays [] [] []uint8

		reBinPath1 := fmt.Sprintf("%s/%d.bin", reBinFolderPath, reBinIndex1)
		reBinPath2 := fmt.Sprintf("%s/%d.bin", reBinFolderPath, reBinIndex2)
		var binPaths [] string
		for j:=reBinIndex1+1;j<reBinIndex2;j++{
			binPaths=append(binPaths,fmt.Sprintf("%s/%d.bin", binFolderPath,j) )
		}

		reBinByteArray1, err := ioutil.ReadFile(reBinPath1)
		if err != nil {
			break
		}

		reBinGrayArray1:=lib.ByteArrayToGrayArray(reBinByteArray1,config)

		reBinByteArray2, err := ioutil.ReadFile(reBinPath2)
		if err != nil {
			break
		}

		reBinGrayArrays2:=lib.ByteArrayToGrayArray(reBinByteArray2,config)

		for _,path :=range binPaths{
			byteArray, err := ioutil.ReadFile(path)
			if err != nil {
				break
			}
			grayArrays:=lib.ByteArrayToGrayArray(byteArray,config)
			bPageGrayArrays=append(bPageGrayArrays,grayArrays)
		}

		byteArray,reGrayArrays := lib.TranscodeGbp(reBinGrayArray1, reBinGrayArrays2, bPageGrayArrays,config)
		err = lib.ArraySaveAsBufferFile(byteArray, gbpPath)
		if err != nil {
			return err
		}
		for i,reGrayArray :=range reGrayArrays{
			reBinPath:=fmt.Sprintf("%s/%d.bin",reBinFolderPath,reBinIndex1+i+1)
			byteArray:=lib.GrayArrayToByteArray(reGrayArray,config)
			err = lib.ArraySaveAsBufferFile(byteArray, reBinPath)
			if err != nil {
				return err
			}
		}
	}

	fmt.Println("BinToGpp Success")
	return nil
}
