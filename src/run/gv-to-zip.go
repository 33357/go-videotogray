package run

import (
	"../lib"
	"fmt"
	"os"
)

func GvToZip(gvFolderPath string,zipFolderPath string,config *lib.ConfigInfo) error{
	_, err := os.Stat(zipFolderPath)
	if err != nil {
		err:=os.MkdirAll(zipFolderPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	var files [] *os.File
	index:=1
	for ;;index++{
		path:=fmt.Sprintf("%s/%d.gv" ,gvFolderPath,index)
		file, err := os.Open(path)
		if err == nil {
			files=append(files, file)
		}
		if (index+1)%(config.ZipSeconds/config.GvSeconds)==0 {
			lib.CompressFiles(files,fmt.Sprintf("%s/%d.zip",zipFolderPath,index))
		}
	}
	if len(files)!=0{
		lib.CompressFiles(files,fmt.Sprintf("%s/%d.zip",zipFolderPath,index))
	}
	return nil
}