function decodeGP(array,config){
    console.log("decodeGP")
    let [basisArrays,differenceArray]=getBasisArrays(array,config)
    let grayArray=getGrayArray(basisArrays,differenceArray,config)
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
    let grayArrays=[]
    for (let i=0;i<config['outHeight'];i++){
        grayArrays.push(new Array(config['outWidth']))
    }
    let skip=config['bPointNum']+1
    for (let i=0;i<config['outHeight'];i+=skip) {
        for (let j=0;j<config['outWidth'];j+=skip){
            grayArrays[i][j]=basisArrays[i/skip][j/skip]
            if (j!=0) {
                let d=grayArrays[i][j-skip]- grayArrays[i][j]
                if (d<0){
                    d=-d
                }
                for (let k=1;k<skip;k++){
                    if (d>skip){
                        grayArrays[i][j+k]=differenceArray[d_index++]
                    }else {
                        if (grayArrays[i][j - skip] > grayArrays[i][j]) {
                            if (j < d) {
                                grayArrays[i][j + k] = grayArrays[i][j] + k
                            } else {
                                grayArrays[i][j + k] = grayArrays[i][j] + d
                            }
                        } else {
                            if (j < d) {
                                grayArrays[i][j + k] = grayArrays[i][j] - k
                            } else {
                                grayArrays[i][j + k] = grayArrays[i][j] - d
                            }
                        }
                    }
                }
            }
            if (i!=0) {
                for (let k=0;k<skip;k++)
                {
                    let d = grayArrays[i-skip][j+k] - grayArrays[i][j+k]
                    if (d < 0) {
                        d = -d
                    }
                    for (let l = 1; l < skip; l++) {
                        // if (d > skip) {
                        //     //grayArrays[i-skip+l][j+k] = differenceArray[d_index++]
                        //     d_index++
                        //
                        // } else {
                            if (grayArrays[i-skip][j+k] > grayArrays[i][j+k]) {
                                if (l < d) {
                                    grayArrays[i-skip+l][j+k] = grayArrays[i-skip][j+k] - l
                                } else {
                                    grayArrays[i-skip+l][j+k] = grayArrays[i-skip][j+k] - d
                                }
                            } else {
                                if (l < d) {
                                    grayArrays[i-skip+l][j+k] = grayArrays[i-skip][j+k] + l
                                } else {
                                    grayArrays[i-skip+l][j+k] = grayArrays[i-skip][j+k] + d
                                }
                            }
                        // }
                    }
                }
            }
        }
    }
    return grayArrays
}