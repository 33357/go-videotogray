package lib

import "fmt"

func TranscodeGV(array [] [] []uint8,config *ConfigInfo) [] uint8 {
	var arr [] uint8
	var IPageArrays [] [] uint8
	var BPageArrays [] [] [] uint8
	var PPageArrays [] [] [] uint8
	length:=len(array)
	pageSkip:=config.BPageNum+1
	for i:=0;i<length;i+=pageSkip {
		if i==0{
			IPageArrays=array[i]
		}else{
			PPageArrays=append(PPageArrays,array[i])
		}
		for j:=1;j<pageSkip&&i+j<length;j++ {
			if i+j==length{
				PPageArrays=append(PPageArrays,array[i+j])
			}else{
				BPageArrays=append(BPageArrays,array[i+j])
			}
		}
	}
	IPageArray:=TranscodeGip(IPageArrays,config)
	arr=append(arr,IPageArray...)
	fmt.Println(len(IPageArray))
	PPageArray:=TranscodeGpp(IPageArrays,PPageArrays[0],config)
	//arr=append(arr,PPageArray...)
	fmt.Println(len(PPageArray))
	//BPageArray:=TranscodeGbp(IPageArrays,PPageArrays,BPageArrays,config)
	//arr=append(arr,BPageArray...)
	//fmt.Println(len(BPageArray))
	//fmt.Println(len(arr))
	return arr
}
