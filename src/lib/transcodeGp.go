package lib

import "fmt"

func TranscodeGP(grayArrays [][] uint8,config *ConfigInfo) [] uint8 {
	if config.BPointNum >5 {
		fmt.Printf("BPointNum over 5 :%d\n",config.BPointNum)
	}
	return getOutArray(grayArrays,config)
}

func getOutArray(grayArrays [][] uint8,config *ConfigInfo) [] uint8 {
	var outArray [] uint8
	var basisArray [] uint8
	var differenceArray [] uint8
	skip:=config.BPointNum+1
	reGrayArrays := make([][]uint8, config.OutHeight)
	for i:=0;i<config.OutHeight;i++ {
		reGrayArrays[i] = make([]uint8, config.OutWidth)
	}
	for h:=0;h<config.OutHeight;h+=skip {
		for w:=0;w<config.OutWidth;w+=skip{
			reGrayArrays[h][w]=grayArrays[h][w]
			basisArray=append(basisArray,grayArrays[h][w])
			if w!=0 {
				d:=int8(grayArrays[h][w-skip])- int8(grayArrays[h][w])
				if d<0{
					d=-d
				}
				for ws:=1;ws<skip;ws++{
					if d==0{
						reGrayArrays[h][w -skip+ws]=grayArrays[h][w-skip]
					}else if d>int8(skip){
						reGrayArrays[h][w -skip+ws]=grayArrays[h][w-skip+ws]
						differenceArray=append(differenceArray,grayArrays[h][w-skip+ws])
					}else{
						if reGrayArrays[h][w-skip]>reGrayArrays[h][w] {
							if int8(ws)<d{
								reGrayArrays[h][w -skip+ws]=reGrayArrays[h][w-skip]-uint8(ws)
							}else{
								reGrayArrays[h][w -skip+ws]=reGrayArrays[h][w-skip]-uint8(d)
							}
						}else{
							if int8(ws)<d{
								reGrayArrays[h][w -skip+ws]=reGrayArrays[h][w-skip]+uint8(ws)
							}else{
								reGrayArrays[h][w -skip+ws]=reGrayArrays[h][w-skip]+uint8(d)
							}
						}
					}
				}
			}
			if h!=0&&w!=0 {
				for ws:=0;ws<skip;ws++{
					d:=int8(reGrayArrays[h-skip][w-skip+ws])- int8(reGrayArrays[h][w-skip+ws])
					if d<0{
						d=-d
					}
					if d>int8(skip){
						for hs:=1;hs<skip;hs++ {
							differenceArray = append(differenceArray, grayArrays[h-skip+hs][w-skip+ws])
						}
					}
				}
			}
		}
	}
	outArray=append(basisArray,differenceArray...)
	return outArray
}

//func getBlockDifferenceArray(basisArrays[] [] uint8,config *ConfigInfo,array [] uint8) [] uint8 {
//	var blockDifferenceArray [] uint8
//
//
//	return blockDifferenceArray
//}
