package lib

func TranscodeGip(IPageArray [][] uint8,config *ConfigInfo) [] uint8 {
	var byteArray [] uint8
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
			byteArray = append(byteArray, IPageArray[w][h])
			reGrayArray[w][h]= IPageArray[w][h]
			if h!=0 {
				transcodeIPageColumn(reGrayArray[w][h-maxColumnSkip], reGrayArray[w][h], IPageArray[w][h-maxColumnSkip+1:h], &byteArray, &reGrayArray, w, h-maxColumnSkip)
				if h+maxColumnSkip >= config.OutHeight-1{
					byteArray = append(byteArray, IPageArray[w][config.OutHeight-1])
					reGrayArray[w][config.OutHeight-1]= IPageArray[w][config.OutHeight-1]
					transcodeIPageColumn(reGrayArray[w][h], reGrayArray[w][config.OutHeight-1], IPageArray[w][h+1:config.OutHeight-1], &byteArray, &reGrayArray, w, h)
				}
			}
		}
		if w != 0 {
			transcodeIPageRow(reGrayArray[w-maxRowSkip],reGrayArray[w],IPageArray[w-maxRowSkip+1:w],&byteArray)
		}
		if w>=config.OutWidth-1 {
			break
		}
	}
	//fmt.Println(len(basisArray),len(differenceArray))
	return byteArray
}

func transcodeIPageColumn(beforeColumnPoint uint8,afterColumnPoint uint8,betweenColumnPoints [] uint8,byteArray *[] uint8,reGrayArrays *[] [] uint8,w int ,ch int) {
	cd:=int(beforeColumnPoint)-int(afterColumnPoint)
	if cd<0{
		cd=-cd
	}
	columnSkip:=len(betweenColumnPoints)+1
	for cs:=1;cs<columnSkip;cs++{
		h:=ch+cs
		if cd>columnSkip{
			*byteArray = append(*byteArray,betweenColumnPoints[cs-1])
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

func transcodeIPageRow(beforeRowColumn []uint8,afterRowColumn [] uint8,betweenRowColumns [][] uint8,byteArray *[] uint8) {
	rowSkip:=len(betweenRowColumns)+1
	length:=len(beforeRowColumn)
	for h:=0;h<length;h++{
		cd:=int(beforeRowColumn[h])- int(afterRowColumn[h])
		if cd<0{
			cd=-cd
		}
		if cd>rowSkip{
			for rs:=1;rs<rowSkip;rs++ {
				*byteArray = append(*byteArray,betweenRowColumns[rs-1][h])
			}
		}
	}
}