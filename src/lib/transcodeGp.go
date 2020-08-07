package lib

import "fmt"

func TranscodeGP(grayArrays [][] uint8,config *ConfigInfo) [] uint8 {
	if config.BPointNum >5 {
		fmt.Printf("BPointNum over 5 :%d\n",config.BPointNum)
	}
	var gpArray [] uint8
	basisArrays:=getBasisArrays(grayArrays,config)
	for  _,arr := range basisArrays  {
		gpArray=append(gpArray,arr...)
	}
	lineDifferenceArray:=getLineDifferenceArray(basisArrays,config,grayArrays)
	gpArray=append(gpArray,lineDifferenceArray...)
	//blockDifferenceArray:=getBlockDifferenceArray(basisArrays,config,array)
	//gpArray=append(gpArray,blockDifferenceArray...)
	return gpArray
}

func getBasisArrays(grayArrays [][] uint8,config *ConfigInfo) [] [] uint8 {
	basisArraysWidth:=config.OutWidth/(config.BPointNum+1)
	basisArraysHeight:=config.OutHeight/(config.BPointNum+1)
	basisArrays := make([][]uint8, basisArraysWidth)
	for i:=0;i<basisArraysWidth;i++ {
		basisArrays[i] = make([]uint8, basisArraysHeight)
	}
	for i:=0;i<basisArraysWidth;i++ {
		for j:=0;j<basisArraysHeight;j++{
			basisArrays[i][j]=grayArrays[i*(config.BPointNum+1)][j*(config.BPointNum+1)]
		}
	}
	return basisArrays
}

func getLineDifferenceArray(basisArrays[] [] uint8,config *ConfigInfo,array [] [] uint8) [] uint8 {
	var lineDifferenceArray [] uint8
	lineArraysWidth:=config.OutWidth
	lineArraysHeight:=config.OutHeight/(config.BPointNum+1)
	lineArrays := make([][]uint8, lineArraysWidth)
	for i:=0;i<lineArraysWidth;i++ {
		lineArrays[i] = make([]uint8,lineArraysHeight)
	}
	for  index,arr := range basisArrays  {
		length:=len(arr)-1
		for  i:=0;i<length;i++  {
			d:=int8(arr[i+1])-int8(arr[i])
			if d<0{
				d=-d
			}
			lineArrays[index*(config.BPointNum+1)][i]=array[index*(config.BPointNum+1)][i*(config.BPointNum+1)]
			if d>int8(config.BPointNum+1) {
				for j:=1;j<config.BPointNum+1;j++{
					lineDifferenceArray=append(lineDifferenceArray,array[index][i*(config.BPointNum+1)+j])
					lineArrays[index*(config.BPointNum+1)][i+j]=array[index*(config.BPointNum+1)][i*(config.BPointNum+1)+j]
				}
			}else{
				for j:=1;j<config.BPointNum+1;j++{
					if arr[i+1]>arr[i] {
						if int8(j)<d{
							lineArrays[index*(config.BPointNum+1)+j][i]=array[index][i*(config.BPointNum+1)]+uint8(j)
						}else{
							lineArrays[index*(config.BPointNum+1)+j][i]=array[index][i*(config.BPointNum+1)]+uint8(d)
						}
					}else{
						if int8(j)<d{
							lineArrays[index*(config.BPointNum+1)+j][i]=array[index][i*(config.BPointNum+1)]-uint8(j)
						}else{
							lineArrays[index*(config.BPointNum+1)+j][i]=array[index][i*(config.BPointNum+1)]-uint8(d)
						}
					}
				}
			}
		}
	}
	fmt.Println(len(lineArrays))
	for  index,arr := range lineArrays  {
		var length=len(arr)
		for  i:=0;i<length-1;i++  {
			d:=int8(arr[i+1])-int8(arr[i])
			if d<0{
				d=-d
			}
			if d>int8(config.BPointNum+1) {
				for j:=1;j<config.BPointNum+1;j++{
					lineDifferenceArray=append(lineDifferenceArray,array[index][i*(config.BPointNum+1)+j])
				}
			}
		}
	}
	fmt.Println(len(lineDifferenceArray))
	return lineDifferenceArray
}

//func getBlockDifferenceArray(basisArrays[] [] uint8,config *ConfigInfo,array [] uint8) [] uint8 {
//	var blockDifferenceArray [] uint8
//
//
//	return blockDifferenceArray
//}
