package run

import (
	"../lib"
	"fmt"
	"os"
	"strings"
)

func GvToZip(vbPath string) error{
	var files [] *os.File
	for i:=0;i<12*5;i++{
		path:=fmt.Sprintf("%s/%d.vb",vbPath,i)
		file, err := os.Open(path)
		if err == nil {
			files=append(files, file)
		}
	}
	zipPath:=strings.Replace(vbPath,"vb","zip",1)
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