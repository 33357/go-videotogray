package lib

func TranscodeGbp(beforePageArray [] []uint8,afterPageArray [] [] uint8,BPageArrays[] [] [] uint8,config *ConfigInfo) [] uint8 {
	bPageLength:=len(BPageArrays)
	var byteArray [] uint8
	reGrayArrays := make([][][]uint8, bPageLength)
	for i:=0;i< bPageLength;i++ {
		reGrayArrays[i] = make([][]uint8, config.OutWidth)
		for j:=0;j<config.OutWidth;j++ {
			reGrayArrays[i][j] = make([]uint8, config.OutHeight)
		}
	}
	maxRowSkip:=config.MaxBRowNum+1
	maxColumnSkip:=config.MaxBColumnNum+1
	for w:=0;;w+=maxRowSkip {
		if w > config.OutWidth-1 {
			maxRowSkip=config.OutWidth-1-(w-maxRowSkip)
			w = config.OutWidth - 1
		}
		for h := 0; h < config.OutHeight-1; h += maxColumnSkip {
			transcodeBPageBasis(beforePageArray[w][h],afterPageArray[w][h],BPageArrays,&reGrayArrays,&byteArray,w,h)
			if h!=0 {
				transcodeBPageColumn(BPageArrays, &reGrayArrays, &byteArray, w, h-maxColumnSkip,maxColumnSkip)
				if h+maxColumnSkip >= config.OutHeight-1{
					transcodeBPageBasis(beforePageArray[w][config.OutHeight-1],afterPageArray[w][config.OutHeight-1],BPageArrays,&reGrayArrays,&byteArray,w,config.OutHeight-1)
					transcodeBPageColumn(BPageArrays, &reGrayArrays, &byteArray, w, h,config.OutHeight-1-h)
				}
			}
		}
		if w != 0 {
			transcodeBPageRow(BPageArrays,&reGrayArrays,&byteArray,w-maxRowSkip,maxRowSkip)
			if w == config.OutWidth-1 {
				break
			}
		}
	}
	return byteArray
}

func transcodeBPageBasis(beforePagePoint uint8,afterPagePoint uint8 ,betweenPageArrays [][][] uint8,reGrayArrays *[][][]uint8,byteArray *[]uint8,w int,h int) {
	pd := int(beforePagePoint) - int(afterPagePoint)
	if pd < 0 {
		pd = -pd
	}
	betweenPageLength := len(betweenPageArrays)
	for p := 0; p < betweenPageLength; p++ {
		if pd > betweenPageLength + 1 {
			*byteArray = append(*byteArray, betweenPageArrays[p][w][h])
			(*reGrayArrays)[p][w][h] = betweenPageArrays[p][w][h]
		}else if pd==0{
			(*reGrayArrays)[p][w][h] =beforePagePoint
		}else{
			if beforePagePoint > afterPagePoint {
				if p+1 < pd {
					(*reGrayArrays)[p][w][h] = beforePagePoint - uint8(p+1)
				} else {
					(*reGrayArrays)[p][w][h] = beforePagePoint - uint8(pd)
				}
			} else {
				if p+1 < pd {
					(*reGrayArrays)[p][w][h] = beforePagePoint + uint8(p+1)
				} else {
					(*reGrayArrays)[p][w][h] = beforePagePoint + uint8(pd)
				}
			}
		}
	}
}

func transcodeBPageColumn(betweenPageArrays [][][] uint8,reGrayArrays *[][][] uint8,byteArray *[] uint8,w int ,ch int,columnSkip int) {
	betweenPageLength := len(betweenPageArrays)
	for p := 0; p < betweenPageLength; p++ {
		beforeColumnPoint:=(*reGrayArrays)[p][w][ch]
		afterColumnPoint:=(*reGrayArrays)[p][w][ch+columnSkip]
		cd:=int(beforeColumnPoint)-int(afterColumnPoint)
		if cd<0{
			cd=-cd
		}
		for cs := 1; cs < columnSkip; cs++ {
			h := ch + cs
			if cd > columnSkip {
				*byteArray = append(*byteArray, betweenPageArrays[p][w][cs-1])
				(*reGrayArrays)[p][w][h] = betweenPageArrays[p][w][cs-1]
			} else if cd == 0 {
				(*reGrayArrays)[p][w][h] = beforeColumnPoint
			} else {
				if beforeColumnPoint > afterColumnPoint {
					if cs < cd {
						(*reGrayArrays)[p][w][h] = beforeColumnPoint - uint8(cs)
					} else {
						(*reGrayArrays)[p][w][h] = beforeColumnPoint - uint8(cd)
					}
				} else {
					if cs < cd {
						(*reGrayArrays)[p][w][h] = beforeColumnPoint + uint8(cs)
					} else {
						(*reGrayArrays)[p][w][h] = beforeColumnPoint + uint8(cd)
					}
				}
			}
		}
	}
}

func transcodeBPageRow(betweenPageArrays [][][] uint8,reGrayArrays *[][][] uint8,byteArray *[] uint8,rw int,rowSkip int) {
	betweenPageLength := len(betweenPageArrays)
	for p := 0; p < betweenPageLength; p++ {
		beforeRowColumn:=(*reGrayArrays)[p][rw]
		afterRowColumn:=(*reGrayArrays)[p][rw+rowSkip]
		length := len(beforeRowColumn)
		for h := 0; h < length; h++ {
			cd := int(beforeRowColumn[h]) - int(afterRowColumn[h])
			if cd < 0 {
				cd = -cd
			}
			if cd > rowSkip {
				for rs := 1; rs < rowSkip; rs++ {
					*byteArray = append(*byteArray, betweenPageArrays[p][rw+rs][h])
				}
			}
		}
	}
}
