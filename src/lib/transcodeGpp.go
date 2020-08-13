package lib

import "fmt"

func TranscodeGpp(beforePageArray [] []uint8,PPageArray[] [] uint8,config *ConfigInfo) [] uint8 {
	var basisArray [] uint8
	var differenceArray [] uint8
	pointSkip:=config.BPointNum+1
	rePPageArrays := make([][]int8, config.OutHeight)
	for i:=0;i<config.OutHeight;i++ {
		rePPageArrays[i] = make([]int8, config.OutWidth)
	}
	for h:=0;h<config.OutHeight;h+=pointSkip {
		for w:=0;w<config.OutWidth;w+=pointSkip {
			bd:=int8(beforePageArray[h][w])-int8(PPageArray[h][w])
			//if bd<0 {
			//	bd+=int8(config.ColorSize)
			//}
			basisArray=append(basisArray,uint8(bd))
			//rePPageArrays[h][w]=bd
			//if w!=0 {
			//	dd:=rePPageArrays[h][w-pointSkip]-rePPageArrays[h][w]
			//	for ws:=0;ws<pointSkip;ws++ {
			//		if dd==0{
			//			rePPageArrays[h][w-pointSkip+ws]=rePPageArrays[h][w-pointSkip]
			//		}else if dd>int8(pointSkip){
			//			pd:=int8(beforePageArray[h][w-pointSkip+ws])-int8(PPageArray[h][w-pointSkip+ws])
			//			if pd<0 {
			//				pd+=int8(config.ColorSize)
			//			}
			//			differenceArray=append(differenceArray,uint8(pd))
			//		}else{
			//			if rePPageArrays[h][w-pointSkip]>rePPageArrays[h][w] {
			//				if int8(ws)<dd{
			//					rePPageArrays[h][w-pointSkip+ws]=rePPageArrays[h][w-pointSkip]-int8(ws)
			//				}else{
			//					rePPageArrays[h][w-pointSkip+ws]=rePPageArrays[h][w-pointSkip]-dd
			//				}
			//			}else{
			//				if int8(ws)<dd{
			//					rePPageArrays[h][w-pointSkip+ws]=rePPageArrays[h][w-pointSkip]+int8(ws)
			//				}else{
			//					rePPageArrays[h][w-pointSkip+ws]=rePPageArrays[h][w-pointSkip]+dd
			//				}
			//			}
			//		}
			//	}
			//}
			//if h!=0&&w!=0 {
			//	for ws:=0;ws<pointSkip;ws++{
			//		dd:=rePPageArrays[h-pointSkip][w-pointSkip+ws]-rePPageArrays[h][w-pointSkip+ws]
			//		if dd<0{
			//			dd=-dd
			//		}
			//		if dd>int8(pointSkip){
			//			for hs:=1;hs<pointSkip;hs++ {
			//				pd:=int8(beforePageArray[h-pointSkip+hs][w-pointSkip+ws])-int8(PPageArray[h-pointSkip+hs][w-pointSkip+ws])
			//				if pd<0 {
			//					pd+=int8(config.ColorSize)
			//				}
			//				differenceArray = append(differenceArray,uint8(pd))
			//			}
			//		}
			//	}
			//}
		}
	}
	fmt.Println(len(basisArray),len(differenceArray))
	outArray:=append(basisArray,differenceArray...)
	return outArray
}
