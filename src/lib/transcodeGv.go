package lib


func TranscodeGV(array [] [] uint8,config *ConfigInfo,dvPath string)  {
	length:=len(array)
	var IPageArray [] uint8
	var BPageArray [] [] uint8
	var PPageArray [] [] uint8
	for i:=0;i<length;i+=config.BPageNum+1 {
		if i==0{
			IPageArray=array[i]
		}else{
			PPageArray=append(PPageArray,array[i])
		}
		for j:=1;j<config.BPageNum+1&&i+j<length;j++ {
			if i+j==length{
				PPageArray=append(PPageArray,array[i+j])
			}else{
				BPageArray=append(BPageArray,array[i+j])
			}
		}
	}
	transcodeIPage(IPageArray,config)
}

func transcodeIPage(array [] uint8,config *ConfigInfo)  {
	//let IPageStr='';
	//let beforeChar='';
	//for (let i=0;i<width;i++){
	//	let thisChar=str[i];
	//	if(i==0){
	//		IPageStr+=difference[value[thisChar]];
	//	}else{
	//		if (beforeChar==thisChar){
	//			IPageStr+=difference[0]
	//		}else{
	//			let difNum=value[thisChar]-value[beforeChar];
	//			if(difNum>=0){
	//				IPageStr+=difference[difNum]
	//			}else{
	//				IPageStr+=difference[length+difNum]
	//			}
	//		}
	//	}
	//	beforeChar=thisChar;
	//}
	//let beforeLine='';
	//for (let i=0;i<height;i++) {
	//	let thisLine=str.substring(i*width,(i+1)*width);
	//	if(i!=0){
	//		for(let j=0;j<beforeLine.length;j++){
	//			if (beforeLine[j]==thisLine[j]){
	//				IPageStr+=difference[0]
	//			}else{
	//				let difNum=value[thisLine[j]]-value[beforeLine[j]];
	//				if(difNum>=0){
	//					IPageStr+=difference[difNum]
	//				}else{
	//					IPageStr+=difference[length+difNum]
	//				}
	//			}
	//		}
	//	}
	//	beforeLine=thisLine;
	//}
	//return IPageStr;
}

func transcodeBPage()  {

}

func transcodePPage()  {

}