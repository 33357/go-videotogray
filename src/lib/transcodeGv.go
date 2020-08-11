package lib


func TranscodeGV(array [] [] []uint8,config *ConfigInfo) [] uint8 {
	length:=len(array)
	var arr [] uint8
	var IPageArray [] [] uint8
	var BPageArray [] [] [] uint8
	var PPageArray [] [] [] uint8
	for i:=0;i<length;i+=config.BPageNum+1 {
		if i==0{
			IPageArray=array[i]
		}else{
			PPageArray=append(PPageArray,array[i])
		}
		for j:=1;j<config.BPageNum+1&&i+j<length;j++ {
			if i+j==length{
				PPageArray=append(PPageArray,array[i+j])
			}else{
				BPageArray=append(BPageArray,array[i+j])
			}
		}
	}
	arr=append(arr,transcodeIPage(IPageArray,config)...)

	return arr
}

func transcodeIPage(array [] []uint8,config *ConfigInfo) [] uint8 {
	return TranscodeGP(array,config)
}

func transcodeBPage()  {

}

func transcodePPage()  {

}