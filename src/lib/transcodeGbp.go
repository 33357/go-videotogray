package lib

func TranscodeGbp(IPageArrays [] []uint8,PPageArrays[] [] [] uint8,BPageArrays[] [] [] uint8,config *ConfigInfo) [] uint8 {
	var basisArray [] uint8
	var differenceArray [] uint8
	beforePageArrays:=IPageArrays
	pointSkip:=config.BPointNum+1
	pageSkip:=config.BPageNum+1
	for index, arr := range PPageArrays{
		afterPageArrays:=arr
		for h:=0;h<config.OutHeight;h+=pointSkip {
			for w:=0;w<config.OutWidth;w+=pointSkip {
				if index==len(PPageArrays)-1{
					pageSkip=len(BPageArrays)-config.BPointNum*len(PPageArrays)
				}
				bd:=int8(beforePageArrays[h][w])-int8(afterPageArrays[h][w])
				if bd<0 {
					bd=-bd
				}
				if bd>int8(pageSkip){
					bd=int8(beforePageArrays[h][w])-int8(afterPageArrays[h][w])
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
		beforePageArrays=afterPageArrays
	}
	outArray:=append(basisArray,differenceArray...)
	return outArray
}
