package run

import (
	"../lib"
	"fmt"
	"io/ioutil"
	"os"
)

func BinToGbp(binFolderPath string,gbpFolderPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(gbpFolderPath)
	if err != nil {
		err:=os.MkdirAll(gbpFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	for i:=1;;i+=config.MaxBPageNum+1 {
		indexb1 := i
		indexp2 := indexb1 + config.MaxBPageNum+1
		var gbpPaths [] string
		var bPageGrayArrays [] [] []uint8
		for j:=indexb1+1;j<indexp2;j++{
			gbpPaths=append(gbpPaths,fmt.Sprintf("%s/%d.gbp", gbpFolderPath,j) )
		}

		bin1Path := fmt.Sprintf("%s/%d.bin", binFolderPath, indexb1)
		bin2Path := fmt.Sprintf("%s/%d.bin", binFolderPath, indexp2)
		var binPaths [] string
		for j:=indexb1+1;j<indexp2;j++{
			binPaths=append(binPaths,fmt.Sprintf("%s/%d.bin", binFolderPath,j) )
		}

		byteArray1, err := ioutil.ReadFile(bin1Path)
		if err != nil {
			return err
		}

		grayArrays1:=lib.ByteArrayToGrayArray(byteArray1,config)

		byteArray2, err := ioutil.ReadFile(bin2Path)
		if err != nil {
			return err
		}

		grayArrays2:=lib.ByteArrayToGrayArray(byteArray2,config)

		for _,path :=range binPaths{
			byteArray, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			grayArrays:=lib.ByteArrayToGrayArray(byteArray,config)
			bPageGrayArrays=append(bPageGrayArrays,grayArrays)
		}

		arrays := lib.TranscodeGbp(grayArrays1, grayArrays2, bPageGrayArrays,config)

		for i,_ :=range gbpPaths{
			err = lib.ArraySaveAsBufferFile(arrays[i], gbpPaths[i])
			if err != nil {
				return err
			}
		}

	}

	fmt.Println("BinToGpp Success")
	return nil
}
