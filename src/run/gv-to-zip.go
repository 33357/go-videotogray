package run

import (
	"../lib"
	"fmt"
	"os"
)

func GvToZip(gvFolderPath string,zipFolderPath string,config *lib.ConfigInfo) error{
	var files [] *os.File
	for i:=0;i<11*5;i++{
		path:=fmt.Sprintf("%s/%d.gv",gvFolderPath,i)
		file, err := os.Open(path)
		if err == nil {
			files=append(files, file)
		}
	}
	_, err := os.Stat(zipFolderPath)
	if err != nil {
		err:=os.MkdirAll(zipFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	lib.CompressFiles(files,zipFolderPath+"/1.zip")
	return nil
}