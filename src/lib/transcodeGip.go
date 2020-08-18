package lib

import "fmt"

func TranscodeGip(GrayArray [][] uint8,config *ConfigInfo) [] uint8 {
	return transcodeIPage(GrayArray,config)
}

func transcodeIPage(IPageArray [][] uint8,config *ConfigInfo) [] uint8  {
	var basisArray [] uint8
	var differenceArray [] uint8
	reGrayArray:= make([][]uint8, config.OutWidth)
	for i:=0;i<config.OutWidth;i++ {
		reGrayArray[i] = make([]uint8, config.OutHeight)
	}
	maxRowSkip:=config.MaxBRowNum+1
	maxColumnSkip:=config.MaxBColumnNum+1
	for w:=0;;w+=maxRowSkip {
		if w>=config.OutWidth-1 {
			w=config.OutWidth-1
		}
		for h:=0;h<config.OutHeight-1;h+=maxColumnSkip {
			basisArray = append(basisArray, IPageArray[w][h])
			reGrayArray[w][h]= IPageArray[w][h]
			if h!=0 {
				transcodeBColumn(reGrayArray[w][h-maxColumnSkip], reGrayArray[w][h], IPageArray[w][h-maxColumnSkip+1:h], &differenceArray, &reGrayArray, w, h-maxColumnSkip)
				if h+maxColumnSkip >= config.OutHeight-1{
					basisArray = append(basisArray, IPageArray[w][config.OutHeight-1])
					reGrayArray[w][config.OutHeight-1]= IPageArray[w][config.OutHeight-1]
					transcodeBColumn(reGrayArray[w][h], reGrayArray[w][config.OutHeight-1], IPageArray[w][h+1:config.OutHeight-1], &differenceArray, &reGrayArray, w, h)
				}
			}
		}
		if w != 0 {
			transcodeBRow(reGrayArray[w-maxRowSkip],reGrayArray[w],IPageArray[w-maxRowSkip+1:w],&differenceArray)
		}
		if w>=config.OutWidth-1 {
			break
		}
	}
	fmt.Println(len(basisArray),len(differenceArray))
	return append(basisArray,differenceArray...)
}

func transcodeBColumn(beforeColumnPoint uint8,afterColumnPoint uint8,betweenColumnPoints [] uint8,differenceArray *[] uint8,reGrayArrays *[] [] uint8,w int ,ch int) {
	cd:=int(beforeColumnPoint)-int(afterColumnPoint)
	if cd<0{
		cd=-cd
	}
	columnSkip:=len(betweenColumnPoints)+1
	for cs:=1;cs<columnSkip;cs++{
		h:=ch+cs
		if cd>columnSkip{
			*differenceArray = append(*differenceArray,betweenColumnPoints[cs-1])
			(*reGrayArrays)[w][h] = betweenColumnPoints[cs-1]
		}else if cd==0 {
			(*reGrayArrays)[w][h] = beforeColumnPoint
		}else{
			if beforeColumnPoint>afterColumnPoint{
				if cs<cd{
					(*reGrayArrays)[w][h]=beforeColumnPoint-uint8(cs)
				}else{
					(*reGrayArrays)[w][h]=beforeColumnPoint-uint8(cd)
				}
			}else{
				if cs<cd{
					(*reGrayArrays)[w][h]=beforeColumnPoint+uint8(cs)
				}else{
					(*reGrayArrays)[w][h]=beforeColumnPoint+uint8(cd)
				}
			}
		}
	}
}

func transcodeBRow(beforeRowColumn []uint8,afterRowColumn [] uint8,betweenRowColumns [][] uint8,differenceArray *[] uint8) {
	rowSkip:=len(betweenRowColumns)+1
	length:=len(beforeRowColumn)
	for h:=0;h<length;h++{
		cd:=int(beforeRowColumn[h])- int(afterRowColumn[h])
		if cd<0{
			cd=-cd
		}
		if cd>rowSkip{
			for rs:=1;rs<rowSkip;rs++ {
				*differenceArray = append(*differenceArray,betweenRowColumns[rs-1][h])
			}
		}
	}
}