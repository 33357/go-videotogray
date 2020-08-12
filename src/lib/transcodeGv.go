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
	IPageArray:=transcodeIPage(IPageArrays,config)
	arr=append(arr,IPageArray...)
	fmt.Println(len(IPageArray))
	PPageArray:=transcodePPage(IPageArrays,PPageArrays,config)
	arr=append(arr,PPageArray...)
	fmt.Println(len(PPageArray))
	BPageArray:=transcodeBPage(IPageArrays,PPageArrays,BPageArrays,config)
	arr=append(arr,BPageArray...)
	fmt.Println(len(BPageArray))
	fmt.Println(len(arr))
	return arr
}

func transcodeIPage(IPageArrays [] []uint8,config *ConfigInfo) [] uint8 {
	return TranscodeGP(IPageArrays,config)
}

func transcodePPage(IPageArrays [] []uint8,PPageArrays[] [] [] uint8,config *ConfigInfo) [] uint8 {
	var basisArray [] uint8
	var differenceArray [] uint8
	beforePageArrays:=IPageArrays
	pointSkip:=config.BPointNum+1
	for _, arr := range PPageArrays{
		rePPageArrays := make([][]int8, config.OutHeight)
		for i:=0;i<config.OutHeight;i++ {
			rePPageArrays[i] = make([]int8, config.OutWidth)
		}
		for h:=0;h<config.OutHeight;h+=pointSkip {
			for w:=0;w<config.OutWidth;w+=pointSkip {
				bd:=int8(beforePageArrays[h][w])-int8(arr[h][w])
				if bd<0 {
					bd+=int8(config.ColorSize)
				}
				basisArray=append(basisArray,uint8(bd))
				rePPageArrays[h][w]=bd
				if w!=0 {
					dd:=rePPageArrays[h][w-pointSkip]-rePPageArrays[h][w]
					for ws:=0;ws<pointSkip;ws++ {
						if dd==0{
							rePPageArrays[h][w-pointSkip+ws]=rePPageArrays[h][w-pointSkip]
						}else if dd>int8(pointSkip){
							pd:=int8(beforePageArrays[h][w-pointSkip+ws])-int8(arr[h][w-pointSkip+ws])
							if pd<0 {
								pd+=int8(config.ColorSize)
							}
							differenceArray=append(differenceArray,uint8(pd))
						}else{
							if rePPageArrays[h][w-pointSkip]>rePPageArrays[h][w] {
								if int8(ws)<dd{
									rePPageArrays[h][w -pointSkip+ws]=rePPageArrays[h][w-pointSkip]-int8(ws)
								}else{
									rePPageArrays[h][w -pointSkip+ws]=rePPageArrays[h][w-pointSkip]-dd
								}
							}else{
								if int8(ws)<dd{
									rePPageArrays[h][w -pointSkip+ws]=rePPageArrays[h][w-pointSkip]+int8(ws)
								}else{
									rePPageArrays[h][w -pointSkip+ws]=rePPageArrays[h][w-pointSkip]+dd
								}
							}
						}
					}
				}
				if h!=0&&w!=0 {
					for ws:=0;ws<pointSkip;ws++{
						dd:=rePPageArrays[h-pointSkip][w-pointSkip+ws]-rePPageArrays[h][w-pointSkip+ws]
						if dd<0{
							dd=-dd
						}
						if dd>int8(pointSkip){
							for hs:=1;hs<pointSkip;hs++ {
								pd:=int8(beforePageArrays[h-pointSkip+hs][w-pointSkip+ws])-int8(arr[h-pointSkip+hs][w-pointSkip+ws])
								if pd<0 {
									pd+=int8(config.ColorSize)
								}
								differenceArray = append(differenceArray,uint8(pd))
							}
						}
					}
				}
			}
		}
		beforePageArrays=arr
	}
	fmt.Println(len(basisArray),len(differenceArray))
	outArray:=append(basisArray,differenceArray...)
	return outArray
}

func transcodeBPage(IPageArrays [] []uint8,PPageArrays[] [] [] uint8,BPageArrays[] [] [] uint8,config *ConfigInfo) [] uint8 {
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
