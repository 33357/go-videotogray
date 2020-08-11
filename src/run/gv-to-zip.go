package run

import (
	"../lib"
	"fmt"
	"os"
)

func GvToZip(gvPath string,zipPath string,config *lib.ConfigInfo) error{
	var files [] *os.File
	for i:=0;i<11*5;i++{
		path:=fmt.Sprintf("%s/%d.gv",gvPath,i)
		file, err := os.Open(path)
		if err == nil {
			files=append(files, file)
		}
	}
	_, err := os.Stat(zipPath)
	if err != nil {
		err:=os.MkdirAll(zipPath,os.ModePerm)
		if err!=nil{
			return err
		}
	}
	lib.CompressFiles(files,zipPath+"/1.zip")
	return nil
}