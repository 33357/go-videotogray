package lib

func TranscodeGbp(beforePageArray [] []uint8,afterPageArray [] [] uint8,bPageArrays[] [] [] uint8,config *ConfigInfo) [][] uint8 {
	basisArrays := make([][]uint8, len(bPageArrays))
	differenceArrays := make([][]uint8, len(bPageArrays))
	reGrayArrays := make([][][]uint8, len(bPageArrays))
	for i:=0;i< len(bPageArrays);i++ {
		reGrayArrays[i] = make([][]uint8, config.OutHeight)
		for j:=0;j<config.OutHeight;j++ {
			reGrayArrays[i][j] = make([]uint8, config.OutWidth)
		}
	}
	pointSkip:=config.BPointNum+1
	pageSkip:=len(bPageArrays)+1
	for h:=0;h<config.OutHeight;h+=pointSkip{
		for w:=0;w<config.OutWidth;w+=pointSkip{
			bd:=int8(beforePageArray[h][w])-int8(afterPageArray[h][w])
			if bd<0 {
				bd=-bd
			}
			if bd>int8(pageSkip){
				for ps:=0;ps<len(bPageArrays);ps++ {
					basisArrays[ps]=append(basisArrays[ps],bPageArrays[ps][h][w])
					reGrayArrays[ps][h][w]=bPageArrays[ps][h][w]
				}
			}else{
				for ps:=0;ps<len(bPageArrays);ps++ {
					d:=int8(beforePageArray[h][w])-int8(afterPageArray[h][w])
					if d<0{
						d=-d
					}
					if beforePageArray[h][w]>afterPageArray[h][w] {
						if int8(ps)<d{
							reGrayArrays[ps][h][w]=beforePageArray[h][w]-uint8(ps)
						}else{
							reGrayArrays[ps][h][w]=beforePageArray[h][w]-uint8(d)
						}
					}else{
						if int8(ps)<d{
							reGrayArrays[ps][h][w]=beforePageArray[h][w]+uint8(ps)
						}else{
							reGrayArrays[ps][h][w]=beforePageArray[h][w]+uint8(d)
						}
					}
				}
			}
			if w!=0 {
				for ps:=0;ps<len(bPageArrays);ps++ {
					d:=int8(reGrayArrays[ps][h][w-pointSkip])- int8(reGrayArrays[ps][h][w])
					if d<0{
						d=-d
					}
					for ws:=1;ws<pointSkip;ws++{
						if d==0{
							reGrayArrays[ps][h][w -pointSkip+ws]=reGrayArrays[ps][h][w-pointSkip]
						}else if d>int8(pointSkip){
							reGrayArrays[ps][h][w -pointSkip+ws]=bPageArrays[ps][h][w-pointSkip+ws]
							differenceArrays[ps]=append(differenceArrays[ps],bPageArrays[ps][h][w-pointSkip+ws])
						}else{
							if reGrayArrays[ps][h][w-pointSkip]>reGrayArrays[ps][h][w] {
								if int8(ws)<d{
									reGrayArrays[ps][h][w -pointSkip+ws]=reGrayArrays[ps][h][w-pointSkip]-uint8(ws)
								}else{
									reGrayArrays[ps][h][w -pointSkip+ws]=reGrayArrays[ps][h][w-pointSkip]-uint8(d)
								}
							}else{
								if int8(ws)<d{
									reGrayArrays[ps][h][w -pointSkip+ws]=reGrayArrays[ps][h][w-pointSkip]+uint8(ws)
								}else{
									reGrayArrays[ps][h][w -pointSkip+ws]=reGrayArrays[ps][h][w-pointSkip]+uint8(d)
								}
							}
						}
					}
				}
			}
			if h!=0&&w!=0 {
				for ps:=0;ps<len(bPageArrays);ps++ {
					if h != 0 && w != 0 {
						for ws := 0; ws < pointSkip; ws++ {
							d := int8(reGrayArrays[ps][h-pointSkip][w-pointSkip+ws]) - int8(reGrayArrays[ps][h][w-pointSkip+ws])
							if d < 0 {
								d = -d
							}
							if d > int8(pointSkip) {
								for hs := 1; hs < pointSkip; hs++ {
									differenceArrays[ps] = append(differenceArrays[ps], bPageArrays[ps][h-pointSkip+hs][w-pointSkip+ws])
								}
							}
						}
					}
				}
			}
		}
	}
	outArrays := make([][]uint8, len(bPageArrays))
	for ps:=0;ps<len(bPageArrays);ps++ {
		outArrays[ps]=append(basisArrays[ps],differenceArrays[ps]...)
	}
	return outArrays
}
