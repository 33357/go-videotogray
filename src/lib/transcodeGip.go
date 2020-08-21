package lib

func TranscodeGip(IPageArray [][] uint8,config *ConfigInfo) ([] uint8,[][] uint8) {
	var byteArray [] uint8
	reGrayArray:= make([][]uint8, config.OutWidth)
	for i:=0;i<config.OutWidth;i++ {
		reGrayArray[i] = make([]uint8, config.OutHeight)
	}
	maxRowSkip:=config.MaxBRowNum+1
	maxColumnSkip:=config.MaxBColumnNum+1
	for w:=0;;w+=maxRowSkip {
		if w>config.OutWidth-1 {
			maxRowSkip=config.OutWidth-1-(w-maxRowSkip)
			w=config.OutWidth-1
		}
		for h:=0;;h+=maxColumnSkip {
			transcodeIPageBasis(IPageArray,&byteArray,&reGrayArray,w,h)
			if h!=0 {
				transcodeIPageColumn(reGrayArray[w][h-maxColumnSkip], reGrayArray[w][h], IPageArray[w][h-maxColumnSkip+1:h], &byteArray, &reGrayArray, w, h-maxColumnSkip)
				if h+maxColumnSkip >= config.OutHeight-1{
					transcodeIPageBasis(IPageArray,&byteArray,&reGrayArray,w,config.OutHeight-1)
					transcodeIPageColumn(reGrayArray[w][h], reGrayArray[w][config.OutHeight-1], IPageArray[w][h+1:config.OutHeight-1], &byteArray, &reGrayArray, w, h)
					break
				}
			}
		}
		if w != 0 {
			transcodeIPageRow(reGrayArray[w-maxRowSkip],reGrayArray[w],IPageArray[w-maxRowSkip+1:w],&byteArray,&reGrayArray,w-maxRowSkip)
			if w==config.OutWidth-1 {
				break
			}
		}
	}
	return byteArray,reGrayArray
}

func transcodeIPageBasis(IPageArray [] [] uint8,byteArray *[]uint8,reGrayArray *[][]uint8,w int,h int) {
	*byteArray = append(*byteArray, IPageArray[w][h])
	(*reGrayArray)[w][h]= IPageArray[w][h]
}

func transcodeIPageColumn(beforeColumnPoint uint8,afterColumnPoint uint8,betweenColumnPoints [] uint8,byteArray *[] uint8,reGrayArray *[] [] uint8,w int ,ch int) {
	cd:=int(beforeColumnPoint)-int(afterColumnPoint)
	if cd<0{
		cd=-cd
	}
	columnSkip:=len(betweenColumnPoints)+1
	for cs:=1;cs<columnSkip;cs++{
		h:=ch+cs
		if cd>columnSkip{
			*byteArray = append(*byteArray,betweenColumnPoints[cs-1])
			(*reGrayArray)[w][h] = betweenColumnPoints[cs-1]
		}else if cd==0 {
			(*reGrayArray)[w][h] = beforeColumnPoint
		}else{
			if beforeColumnPoint>afterColumnPoint{
				if cs<cd{
					(*reGrayArray)[w][h]=beforeColumnPoint-uint8(cs)
				}else{
					(*reGrayArray)[w][h]=beforeColumnPoint-uint8(cd)
				}
			}else{
				if cs<cd{
					(*reGrayArray)[w][h]=beforeColumnPoint+uint8(cs)
				}else{
					(*reGrayArray)[w][h]=beforeColumnPoint+uint8(cd)
				}
			}
		}
	}
}

func transcodeIPageRow(beforeRowColumn []uint8,afterRowColumn [] uint8,betweenRowColumns [][] uint8,byteArray *[] uint8,reGrayArray *[] [] uint8,rw int) {
	rowSkip:=len(betweenRowColumns)+1
	length:=len(beforeRowColumn)
	for h:=0;h<length;h++{
		rd:=int(beforeRowColumn[h])- int(afterRowColumn[h])
		if rd<0{
			rd=-rd
		}
		for rs:=1;rs<rowSkip;rs++ {
			w:=rw+rs
			if rd>rowSkip{
				*byteArray = append(*byteArray,betweenRowColumns[rs-1][h])
				(*reGrayArray)[w][h] = betweenRowColumns[rs-1][h]
			}else if rd==0 {
				(*reGrayArray)[w][h] = beforeRowColumn[h]
			}else{
				if beforeRowColumn[h] > afterRowColumn[h] {
					if rs < rd {
						(*reGrayArray)[w][h] =beforeRowColumn[h] - uint8(rs)
					} else{
						(*reGrayArray)[w][h] =beforeRowColumn[h] - uint8(rd)
					}
				} else {
					if rs < rd {
						(*reGrayArray)[w][h] =beforeRowColumn[h] + uint8(rs)
					} else{
						(*reGrayArray)[w][h] =beforeRowColumn[h] + uint8(rd)
					}
				}
			}
		}
	}
}