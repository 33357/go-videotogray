package lib

func TranscodeGV(array [] [] []uint8,config *ConfigInfo) [] uint8 {
	var arr [] uint8
	var IPageArray [] [] uint8
	var BPageArray [] [] [] uint8
	var PPageArray [] [] [] uint8
	length:=len(array)
	pageSkip:=config.BPageNum+1
	for i:=0;i<length;i+=pageSkip {
		if i==0{
			IPageArray=array[i]
		}else{
			PPageArray=append(PPageArray,array[i])
		}
		for j:=1;j<pageSkip&&i+j<length;j++ {
			if i+j==length{
				PPageArray=append(PPageArray,array[i+j])
			}else{
				BPageArray=append(BPageArray,array[i+j])
			}
		}
	}
	arr=append(arr,transcodeIPage(IPageArray,config)...)
	arr=append(arr,transcodePPage(IPageArray,PPageArray,config)...)
	arr=append(arr,transcodeBPage(IPageArray,PPageArray,BPageArray,config)...)
	return arr
}

func transcodeIPage(array [] []uint8,config *ConfigInfo) [] uint8 {
	return TranscodeGP(array,config)
}

func transcodePPage(IPageArray [] []uint8,PPageArray[] [] [] uint8,config *ConfigInfo) [] uint8 {
	var basisArray [] uint8
	var differenceArray [] uint8
	beforePageArray:=IPageArray
	pointSkip:=config.BPointNum+1
	for _, arr := range PPageArray{
		for h:=0;h<config.OutHeight;h+=pointSkip {
			for w:=0;w<config.OutWidth;w+=pointSkip {
				//fmt.Println(h,w)
				bd:=int8(beforePageArray[h][w])-int8(arr[h][w])
				if bd<0 {
					bd+=int8(config.ColorSize)
				}
				basisArray=append(basisArray,uint8(bd))
				if w!=0 {
					dd:=int8(beforePageArray[h][w])-int8(arr[h][w])
					if dd<0 {
						dd=-dd
					}
					if dd>int8(pointSkip){
						for ws:=0;ws<pointSkip;ws++ {
							pd:=int8(beforePageArray[h][w-pointSkip+ws])-int8(arr[h][w-pointSkip+ws])
							if pd<0 {
								pd=-pd
							}
							differenceArray=append(differenceArray,uint8(pd))
						}
					}
				}
				if h!=0&&w!=0 {
					for ws:=0;ws<pointSkip;ws++{
						dd:=int8(arr[h-pointSkip][w-pointSkip+ws])-int8(arr[h][w-pointSkip+ws])
						if dd<0{
							dd=-dd
						}
						if dd>int8(pointSkip){
							for hs:=1;hs<pointSkip;hs++ {
								pd:=int8(beforePageArray[h-pointSkip+hs][w-pointSkip+ws])-int8(arr[h-pointSkip+hs][w-pointSkip+ws])
								if pd<0 {
									pd=-pd
								}
								differenceArray = append(differenceArray,uint8(pd))
							}
						}
					}
				}
			}
		}
		beforePageArray=arr
	}
	outArray:=append(basisArray,differenceArray...)
	return outArray
}

func transcodeBPage(IPageArray [] []uint8,PPageArray[] [] [] uint8,BPageArray[] [] [] uint8,config *ConfigInfo) [] uint8 {
	var basisArray [] uint8
	var differenceArray [] uint8
	beforePageArray:=IPageArray
	pointSkip:=config.BPointNum+1
	pageSkip:=config.BPageNum+1
	for index, arr := range PPageArray{
		afterPageArray:=arr
		for h:=0;h<config.OutHeight;h+=pointSkip {
			for w:=0;w<config.OutWidth;w+=pointSkip {
				if index==len(PPageArray)-1{
					pageSkip=len(BPageArray)-config.BPointNum*len(PPageArray)
				}
				bd:=int8(beforePageArray[h][w])-int8(afterPageArray[h][w])
				if bd<0 {
					bd=-bd
				}
				if bd>int8(pageSkip){
					bd=int8(beforePageArray[h][w])-int8(afterPageArray[h][w])
					if bd<0 {
						bd+=int8(config.ColorSize)
					}
					basisArray=append(basisArray,uint8(bd))
				}
				//if w!=0 {
				//	dd:=int8(beforePageArray[h][w])-int8(arr[h][w])
				//	if dd<0 {
				//		dd=-dd
				//	}
				//	if dd>int8(pointSkip){
				//		for ws:=0;ws<pointSkip;ws++ {
				//			pd:=int8(beforePageArray[h][w-pointSkip+ws])-int8(arr[h][w-pointSkip+ws])
				//			if pd<0 {
				//				pd=-pd
				//			}
				//			differenceArray=append(differenceArray,uint8(pd))
				//		}
				//	}
				//}
				//if h!=0&&w!=0 {
				//	for ws:=0;ws<pointSkip;ws++{
				//		dd:=int8(arr[h-pointSkip][w-pointSkip+ws])-int8(arr[h][w-pointSkip+ws])
				//		if dd<0{
				//			dd=-dd
				//		}
				//		if dd>int8(pointSkip){
				//			for hs:=1;hs<pointSkip;hs++ {
				//				pd:=int8(beforePageArray[h-pointSkip+hs][w-pointSkip+ws])-int8(arr[h-pointSkip+hs][w-pointSkip+ws])
				//				if pd<0 {
				//					pd=-pd
				//				}
				//				differenceArray = append(differenceArray,uint8(pd))
				//			}
				//		}
				//	}
				//}
			}
		}
		beforePageArray=afterPageArray
	}
	outArray:=append(basisArray,differenceArray...)
	return outArray
}
