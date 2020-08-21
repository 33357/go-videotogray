console.log('load decodeGv')
function decodeGv(byteArray,config){
    console.log("decodeGv")
    this.byteArray=byteArray
    this.byteArrayIndex=0
    let grayArrays=[]
    let pageSkip=config['maxBPageNum']+1
    let length=config['gvSeconds']*config['outFrame']
    let BeforeIPageGrayArray
    let AfterIPageGrayArray
    let BPageGrayArrays
    [BeforeIPageGrayArray,this.byteArrayIndex]=decodeGip(this.byteArray,config,this.byteArrayIndex)
    grayArrays.push(BeforeIPageGrayArray)
    for (let i=1+pageSkip;;i+=pageSkip){
        if (i > length) {
            pageSkip=length-(i-pageSkip)
            i = length
        }
        console.log(i);
        [AfterIPageGrayArray,this.byteArrayIndex]=decodeGip(this.byteArray,config,this.byteArrayIndex);
        if(this.byteArrayIndex==this.byteArray.length){
            grayArrays.push(AfterIPageGrayArray)
            break
        }
        console.log(this.byteArrayIndex,pageSkip-1);
        [BPageGrayArrays,this.byteArrayIndex]=decodeGbp(BeforeIPageGrayArray,AfterIPageGrayArray,this.byteArray,config,pageSkip-1,this.byteArrayIndex)
        for (let j=0;j<BPageGrayArrays.length;j++){
            grayArrays.push(BPageGrayArrays[j])
        }
        grayArrays.push(AfterIPageGrayArray)
        BeforeIPageGrayArray=AfterIPageGrayArray
        if(i == length){
            break
        }
    }
    console.log(grayArrays.length)
    return grayArrays
}