package lib

func TranscodeGip(GrayArray [][] uint8,config *ConfigInfo) [] uint8 {
	return TranscodeIPage(GrayArray,config)
}

func TranscodeIPage(IPageArray [][] uint8,config *ConfigInfo) [] uint8  {
	var basisArray [] uint8
	var differenceArray [] uint8
	reGrayArray:= make([][]uint8, config.OutHeight)
	for i:=0;i<config.OutHeight;i++ {
		reGrayArray[i] = make([]uint8, config.OutWidth)
	}
	maxRowSkip:=config.MaxBRowNum+1
	maxColumnSkip:=config.MaxBColumnNum+1
	for h:=0;h<config.OutHeight;h+=maxColumnSkip {
		for w:=maxRowSkip;w<config.OutWidth;w+=maxRowSkip {
			basisArray=append(basisArray,IPageArray[h][w-maxRowSkip])
			TranscodeBRow(IPageArray[h][w-maxRowSkip],IPageArray[h][w],IPageArray[h][w-maxRowSkip:w],&differenceArray,&reGrayArray,h,w)
			if w+maxRowSkip>=config.OutWidth{
				if w==config.OutWidth-1{
					basisArray=append(basisArray,IPageArray[h][w])
				}else{
					basisArray=append(basisArray,IPageArray[h][w],IPageArray[h][config.OutWidth-1])
					TranscodeBRow(IPageArray[h][w],IPageArray[h][config.OutWidth-1],IPageArray[h][w:config.OutWidth-1],&differenceArray,&reGrayArray,h,w)
				}
			}
		}
		if h != 0 {
			TranscodeBColumn(reGrayArray[h-maxColumnSkip],reGrayArray[h],IPageArray[h-maxColumnSkip:h],&differenceArray)
			if h+maxColumnSkip>=config.OutHeight&&h!=config.OutHeight-1 {
				TranscodeBColumn(reGrayArray[h],reGrayArray[config.OutHeight-1],IPageArray[h:config.OutHeight-1],&differenceArray)
			}
		}
	}

	return append(basisArray,differenceArray...)
}

func TranscodeBRow(beforePoint uint8,afterPoint uint8,betweenPoints [] uint8,differenceArray *[] uint8,reGrayArrays *[] [] uint8,h int ,w int) {
	pd:=int(beforePoint)-int(afterPoint)
	if pd<0{
		pd=-pd
	}
	rowSkip:=len(betweenPoints)
	rw:=w-rowSkip
	for rs:=0;rs<rowSkip;rs++{
		if pd>rowSkip{
			*differenceArray = append(*differenceArray, betweenPoints[rs])
			(*reGrayArrays)[h][rw+rs] = betweenPoints[rs]
		}else if pd==0 {
			(*reGrayArrays)[h][rw+rs]=beforePoint
		}else{
			if beforePoint>afterPoint{
				if rs<pd{
					(*reGrayArrays)[h][rw+rs]=(*reGrayArrays)[h][rw]-uint8(rs)
				}else{
					(*reGrayArrays)[h][rw+rs]=(*reGrayArrays)[h][rw]-uint8(pd)
				}
			}else{
				if rs<pd{
					(*reGrayArrays)[h][rw+rs]=(*reGrayArrays)[h][rw]+uint8(rs)
				}else{
					(*reGrayArrays)[h][rw+rs]=(*reGrayArrays)[h][rw]+uint8(pd)
				}
			}
		}
	}
}

func TranscodeBColumn(beforeRow []uint8,afterRow [] uint8,betweenRows [][] uint8,differenceArray *[] uint8) {
	columnSkip:=len(betweenRows)
	length:=len(beforeRow)
	for cs:=0;cs<length;cs++{
		cd:=int(beforeRow[cs])- int(afterRow[cs])
		if cd<0{
			cd=-cd
		}
		if cd>columnSkip{
			for hs:=0;hs<columnSkip;hs++ {
				*differenceArray = append(*differenceArray,betweenRows[hs][cs])
			}
		}
	}
}