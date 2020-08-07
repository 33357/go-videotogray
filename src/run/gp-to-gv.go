package run

import (
	"../lib"
	"fmt"
	"io/ioutil"
	"strings"
)

func GpToGv(gpPath string,config *lib.ConfigInfo) (string,error){
	dvPath:=strings.Replace(gpPath,"gp","gv",1)
	var arr [] [] uint8
	for i:=0;;i++{
		path:=fmt.Sprintf("%s/%d.gp",gpPath,i)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			break
		}
		arr= append(arr, data)
		if len(arr)>config.GvSeconds*config.OutFrame {
			lib.TranscodeGV(arr,config,dvPath)
			arr = [] [] uint8{}
		}
	}
	if len(arr)!=0 {
		lib.TranscodeGV(arr,config,dvPath)
	}
	return dvPath,nil
}