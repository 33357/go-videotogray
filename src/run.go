package main

import (
	"./lib"
	"./run"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	config,err:= lib.GetConfig()
	if err != nil {
		fmt.Printf(err.Error())
	}

	dir, _ := os.Getwd()
	sourceFolderPath:=path.Join(dir,fmt.Sprintf("%s/%s/source",config.OutPath,getVideoName(config)))
	mp3Path:=fmt.Sprintf("%s/%s_%s.mp3",sourceFolderPath,getVideoName(config),config.Mp3Bit)
	pngFolderPath:=fmt.Sprintf("%s/%s_w%d_h%d_f%d",sourceFolderPath,getVideoName(config),config.SourceWidth,config.SourceHeight,config.SourceFrame)
	binFolderPath:=strings.Replace(sourceFolderPath,"source",fmt.Sprintf("w%d_h%d_f%d_s%d/bin",config.OutWidth,config.OutHeight,config.OutFrame,config.ColorSize),1)
	gipFolderPath:=strings.Replace(binFolderPath,"bin",fmt.Sprintf("gip/br%d_bc%d",config.MaxBRowNum,config.MaxBColumnNum),1)
	//gppFolderPath:=strings.Replace(binFolderPath,"bin",fmt.Sprintf("gpp/bpo%d_bpg%d_gv%d",config.BPointNum,config.BPageNum,config.GvSeconds),1)
	gbpFolderPath:=strings.Replace(binFolderPath,"bin",fmt.Sprintf("gbp/br%d_bc%d_bg%d",config.MaxBRowNum,config.MaxBColumnNum,config.MaxBPageNum),1)
	//gvFolderPath:=strings.Replace(binFolderPath,"bin",fmt.Sprintf("gv/bpo%d_bpg%d_gv%d",config.BPointNum,config.BPageNum,config.GvSeconds),1)
	//zipFolderPath:=strings.Replace(binFolderPath,"bin",fmt.Sprintf("zip/bpo%d_bpg%d_gv%d_zip%d_mp3%s",config.BPointNum,config.BPageNum,config.GvSeconds,config.ZipSeconds,config.Mp3Bit),1)

	err=run.VideoToSource(sourceFolderPath,mp3Path,pngFolderPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.PngToBin(pngFolderPath,binFolderPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.BinToGip(binFolderPath,gipFolderPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err=run.BinToGbp(binFolderPath,gbpFolderPath,config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	//err=run.BinToGip(binFolderPath,gipFolderPath,config)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	//err=run.BinToGv(binFolderPath,gvFolderPath,config)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	//err=run.GvToZip(gvFolderPath,zipFolderPath,config)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	fmt.Println("run success!")
}


func getVideoName(config *lib.ConfigInfo) string {
	arr1 :=strings.Split(config.VideoPath, "/")
	arr2 :=strings.Split(arr1[len(arr1)-1],".")
	return arr2[0]
}


