package lib

func TranscodeGV(array [] [] []uint8,config *ConfigInfo) [] uint8 {
	var IPageArrays [][][]uint8
	var BPageArrays [][][][]uint8
	length := len(array)
	pageSkip := config.MaxBPageNum + 1
	for i := 0; i < length; i += pageSkip {
		IPageArrays = append(IPageArrays, array[i])
		var _BPageArrays [][][]uint8
		for j := 1; j < pageSkip && i+j < length; j++ {
			if i+j == length-1 {
				IPageArrays = append(IPageArrays, array[i+j])
			} else {
				_BPageArrays = append(_BPageArrays, array[i+j])
			}
		}
		BPageArrays = append(BPageArrays, _BPageArrays)
	}
	var IPageByteArrays [][]uint8
	for _, arr := range IPageArrays {
		byteArray, _ := TranscodeGip(arr, config)
		IPageByteArrays = append(IPageByteArrays, byteArray)
	}
	var BPageByteArrays [][]uint8
	for i, arr := range BPageArrays {
		byteArray, _ := TranscodeGbp(IPageArrays[i], IPageArrays[i+1], arr, config)
		BPageByteArrays = append(BPageByteArrays, byteArray)
	}
	var byteArray = IPageByteArrays[0]
	for i, arr := range IPageByteArrays {
		if i > 0 {
			byteArray = append(byteArray, append(arr, BPageByteArrays[i-1]...)...)
		}
	}
	return byteArray
}