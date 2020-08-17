package lib

func TranscodeGV(array [] [] []uint8,config *ConfigInfo) [] uint8 {
	var IPageArrays [] [] [] uint8
	var BPageArrays [] [] [] [] uint8
	length:=len(array)
	pageSkip:=config.MaxBPageNum+1
	for i:=0;i<length;i+=pageSkip {
		IPageArrays=append(IPageArrays,array[i])
		var _BPageArrays [] [] [] uint8
		for j:=1;j<pageSkip&&i+j<length;j++ {
			if i+j==length-1 {
				IPageArrays=append(IPageArrays,array[i+j])
			}else{
				_BPageArrays=append(_BPageArrays,array[i+j])
			}
		}
		BPageArrays=append(BPageArrays,_BPageArrays)
	}
	var IPageArray [] uint8
	for _,arr :=range IPageArrays{
		IPageArray=append(IPageArray,TranscodeGip(arr,config)...)
	}
	var BPageArray [] uint8
	for i,arr :=range BPageArrays{
		_BPageArrays:=TranscodeGbp(IPageArrays[i],IPageArrays[i+1],arr,config)
		for _,arr :=range _BPageArrays{
			BPageArray=append(BPageArray,arr...)
		}
	}
	outArray:=append(IPageArray,BPageArray...)
	return outArray
}
