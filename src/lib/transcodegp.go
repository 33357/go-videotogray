package lib

import "fmt"

func TranscodeGP(array [] uint8,config *ConfigInfo) [] uint8 {
	if config.BPointNum >5 {
		fmt.Printf("BPointNum over 5 :%d\n",config.BPointNum)
	}
	var basisArray [] [] uint8
	for i:=0;i<config.OutWidth;i+=config.BPointNum+1{
		var arr [] uint8
		for j:=0;j<config.OutHeight;i+=config.BPointNum+1{
			arr=append(arr,array[i+j*config.OutWidth])
		}
		basisArray=append(basisArray,arr)
	}
	var arr [] uint8
	var horizontalArray [] uint8
	for  _,_arr := range basisArray  {
		var length=len(_arr)
		for  i:=0;i<length-1;i++  {
			d:=_arr[i]-_arr[i+1]
			if d<0{
				d=-d
			}
			if d>uint8(config.BPointNum+1) {
				arr=append(arr,array[i*(config.BPointNum+1)+1],array[i*(config.BPointNum+1)+2])
				horizontalArray=append(horizontalArray,array[i*(config.BPointNum+1)],array[i*(config.BPointNum+1)+1],array[i*(config.BPointNum+1)+2])
			}else{
				
			}
		}
		horizontalArray=append(horizontalArray,array[length-1])
	}

	return
}
