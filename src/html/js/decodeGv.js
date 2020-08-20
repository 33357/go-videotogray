function decodeGv(byteArray,config){
    console.log("decodeGV")
    this.byteArray=byteArray
    this.byteArrayIndex=0
    let grayArrays=[]
    let pageSkip=config['maxBPageNum']+1
    let length=config['gvSeconds']*config['outFrame']
    let IPageGrayArray
    let BPageGrayArrays
    for (let i=0;i<length;i+=pageSkip) {
        [IPageGrayArray,this.byteArrayIndex]=decodeGip(this.byteArray,config,this.byteArrayIndex)
        grayArrays.push(IPageGrayArray)
        for(let j=1;j<pageSkip&&i+j<length;j++) {
            if(i+j==length-1){
                [IPageGrayArray,this.byteArrayIndex]=decodeGip(this.byteArray,config,this.byteArrayIndex)
                grayArrays.push(IPageGrayArray)
            }else{
                [BPageGrayArrays,this.byteArrayIndex]=decodeGbp(byteArray,config,config['maxBPageNum'],this.byteArrayIndex)
                for (let k=0;k<BPageGrayArrays.length;k++){
                    grayArrays.push(BPageGrayArrays[i])
                }
            }
        }
    }
    return grayArrays
}