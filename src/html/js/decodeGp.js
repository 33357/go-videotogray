function decodeGP(array,config){
    console.log("decodeGP")
    let [basisArrays,differenceArray]=getBasisArrays(array,config)
    let grayArray=getGrayArray(basisArrays,differenceArray,config)
    console.log(grayArray.length)
    return grayArray
}

function getBasisArrays(array,config){
    let basisArrays=[]
    let width=config['outWidth']/(config['bPointNum']+1)
    let height=config['outHeight']/(config['bPointNum']+1)
    for(let i=0;i<height;i++){
        basisArrays.push(array.slice(i*width,(i+1)*width))
    }
    return [basisArrays,array.slice(width*height)]
}

function getGrayArray(basisArrays,differenceArray,config) {
    let d_index=0
    let grayArray=[]
    let lineArrays=[]
    let lineArraysWidth=config['outWidth']
    let lineArraysHeight=config['outHeight']
    for (let i=0;i<lineArraysWidth;i++){
        lineArrays.push(new Array(lineArraysHeight))
    }
    for (let index in basisArrays) {
        let len=basisArrays[index].length-1
        for (let i=0;i<len;i++) {
            let d=basisArrays[index][i+1]-basisArrays[index][i]
            if(d<0){
                d=-d
            }
            lineArrays[index*(config['bPointNum']+1)][i*(config['bPointNum']+1)]=basisArrays[index][i*(config['bPointNum']+1)]
            if (d>config['bPointNum']+1) {
                for (let j=1;j<config['bPointNum']+1;j++){
                    lineArrays[index*(config['bPointNum']+1)+j][i*(config['bPointNum']+1)]=differenceArray[d_index++]
                }
            }else{
                for (let j=1;j<config['bPointNum']+1;j++){
                    if(basisArrays[index][i+1]>basisArrays[index][i]) {
                        if(j<d){
                            lineArrays[index*(config['bPointNum']+1)+j][i*(config['bPointNum']+1)]=basisArrays[index][i]+j
                        }else{
                            lineArrays[index*(config['bPointNum']+1)+j][i*(config['bPointNum']+1)]=basisArrays[index][i]+d
                        }
                    }else{
                        if(j<d){
                            lineArrays[index*(config['bPointNum']+1)+j][i*(config['bPointNum']+1)]=basisArrays[index][i]-j
                        }else{
                            lineArrays[index*(config['bPointNum']+1)+j][i*(config['bPointNum']+1)]=basisArrays[index][i]-d
                        }
                    }
                }
            }
        }
    }
    for (let index in lineArrays){
        let len=lineArrays[index].length-1
        for (let i=0;i<len;i++) {
            let d=lineArrays[index][i+1]-lineArrays[index][i]
            if(d<0){
                d=-d
            }
            if (d>config['bPointNum']+1){
                for (let j=1;j<config['bPointNum']+1;j++){
                    lineArrays[index][i+j]=differenceArray[d_index++]
                }
            }else{
                for (let j=1;j<config['bPointNum']+1;j++){
                    if(lineArrays[index][i+1]>lineArrays[index][i]) {
                        if(j<d){
                            lineArrays[index][i+j]=lineArrays[index][i]+j
                        }else{
                            lineArrays[index][i+j]=lineArrays[index][i]+d
                        }
                    }else{
                        if(j<d){
                            lineArrays[index][i+j]=lineArrays[index][i]-j
                        }else{
                            lineArrays[index][i+j]=lineArrays[index][i]-d
                        }
                    }
                }
            }
        }
    }
    for (let j=0;j<lineArraysHeight;j++){
        for (let i=0;i<lineArraysWidth;i++){
            grayArray.push(lineArrays[i][j])
        }
    }
    return grayArray
}