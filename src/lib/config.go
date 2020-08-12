package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

type ConfigInfo struct {
	SourceWidth		int
	SourceHeight	int
	SourceFrame		int
	OutWidth		int
	OutHeight		int
	OutFrame		int
	ColorSize		int
	GvSeconds 		int
	ZipSeconds 		int
	BPageNum		int
	BPointNum		int
	Thread			int
	Mp3Bit			string
	FFMPEGPath		string
	VideoPath		string
	OutPath			string
}

func GetConfig() (*ConfigInfo ,error){
	dir, _ := os.Getwd()
	filePtr, err := os.Open(path.Join(dir,"config.json"))
	if err != nil {
		return nil,err
	}
	defer filePtr.Close()

	var config ConfigInfo
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&config)
	if err != nil {
		return nil,errors.New(err.Error())
	} else {
		fmt.Println("GetConfig success")
		return &config,nil
	}
}

