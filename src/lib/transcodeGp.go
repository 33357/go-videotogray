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
	for i:=0;i<config.OutHeight;i+=skip {
		for j:=0;j<config.OutWidth;j+=skip{
			reGrayArrays[i][j]=grayArrays[i][j]
			basisArray=append(basisArray,grayArrays[i][j])
			if j!=0 {
				d:=int8(grayArrays[i][j-skip])- int8(grayArrays[i][j])
				if d<0{
					d=-d
				}
				if d>int8(skip){
					for k:=1;k<skip;k++{
						reGrayArrays[i][j+k]=grayArrays[i][j+k]
					}
					differenceArray=append(differenceArray,grayArrays[i][j])
				}else{
					for k:=1;k<skip;k++{
						if reGrayArrays[i][j-skip]>reGrayArrays[i][j] {
							if int8(k)<d{
								reGrayArrays[i][j+k]=reGrayArrays[i][j]+uint8(k)
							}else{
								reGrayArrays[i][j+k]=reGrayArrays[i][j]+uint8(d)
							}
						}else{
							if int8(k)<d{
								reGrayArrays[i][j+k]=reGrayArrays[i][j]-uint8(k)
							}else{
								reGrayArrays[i][j+k]=reGrayArrays[i][j]-uint8(d)
							}
						}
					}
				}
			}
			//if i!=0 {
			//	for k:=0;k<skip;k++{
			//		d:=int8(reGrayArrays[i-skip][j+k])- int8(reGrayArrays[i][j+k])
			//		if d<0{
			//			d=-d
			//		}
			//		if d>int8(skip){
			//			for l:=1;l<skip;l++ {
			//				differenceArray = append(differenceArray, grayArrays[i-skip+l][j+k])
			//			}
			//		}
			//	}
			//}
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
