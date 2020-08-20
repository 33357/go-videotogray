package lib

func TranscodeGbp(beforePageArray [] []uint8,afterPageArray [] [] uint8,BPageArrays[] [] [] uint8,config *ConfigInfo) ([] uint8,[][][] uint8) {
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
		for h := 0;;h += maxColumnSkip {
			transcodeBPageBasis(beforePageArray[w][h],afterPageArray[w][h],BPageArrays,&byteArray,&reGrayArrays,w,h)
			if h!=0 {
				transcodeBPageColumn(BPageArrays, &byteArray, &reGrayArrays, w, h-maxColumnSkip,maxColumnSkip)
				if h+maxColumnSkip >= config.OutHeight-1{
					transcodeBPageBasis(beforePageArray[w][config.OutHeight-1],afterPageArray[w][config.OutHeight-1],BPageArrays,&byteArray,&reGrayArrays,w,config.OutHeight-1)
					transcodeBPageColumn(BPageArrays, &byteArray, &reGrayArrays, w, h,config.OutHeight-1-h)
					break
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
	return byteArray,reGrayArrays
}

func transcodeBPageBasis(beforePagePoint uint8,afterPagePoint uint8 ,betweenPageArrays [][][] uint8,byteArray *[]uint8,reGrayArrays *[][][]uint8,w int,h int) {
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
			(*reGrayArrays)[p][w][h] = beforePagePoint
		}else{
			ps:=p+1
			if beforePagePoint > afterPagePoint {
				if ps < pd {
					(*reGrayArrays)[p][w][h] = beforePagePoint - uint8(ps)
				} else {
					(*reGrayArrays)[p][w][h] = beforePagePoint - uint8(pd)
				}
			} else {
				if ps < pd {
					(*reGrayArrays)[p][w][h] = beforePagePoint + uint8(ps)
				} else {
					(*reGrayArrays)[p][w][h] = beforePagePoint + uint8(pd)
				}
			}
		}
	}
}

func transcodeBPageColumn(betweenPageArrays [][][] uint8,byteArray *[] uint8,reGrayArrays *[][][] uint8,w int ,ch int,columnSkip int) {
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
				*byteArray = append(*byteArray, betweenPageArrays[p][w][h])
				(*reGrayArrays)[p][w][h] = betweenPageArrays[p][w][h]
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
			rd := int(beforeRowColumn[h]) - int(afterRowColumn[h])
			if rd < 0 {
				rd = -rd
			}
			for rs := 1; rs < rowSkip; rs++ {
				w := rw + rs
				if rd > rowSkip {
					*byteArray = append(*byteArray, betweenPageArrays[p][w][h])
					(*reGrayArrays)[p][w][h] = betweenPageArrays[p][w][h]
				}else if rd == 0 {
					(*reGrayArrays)[p][w][h] = beforeRowColumn[h]
				} else {
					if beforeRowColumn[h] > afterRowColumn[h] {
						if rs < rd {
							(*reGrayArrays)[p][w][h] = beforeRowColumn[h] - uint8(rs)
						} else {
							(*reGrayArrays)[p][w][h] = beforeRowColumn[h] - uint8(rd)
						}
					} else {
						if rs < rd {
							(*reGrayArrays)[p][w][h] = beforeRowColumn[h] + uint8(rs)
						} else {
							(*reGrayArrays)[p][w][h] = beforeRowColumn[h] + uint8(rd)
						}
					}
				}
			}
		}
	}
}
