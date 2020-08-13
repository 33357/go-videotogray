function decodeGpp(binArray,byteArray, config){
    console.log("decodeGP")
    let [basisArrays,differenceArray]=getBasisArrays(byteArray,config)
    let grayArray=getGrayArray(binArray,basisArrays,differenceArray,config)
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

function getGrayArray(binArray,basisArrays,differenceArray,config) {
    let d_index=0
    let grayArrays=[]
    for (let i=0;i<config['outHeight'];i++){
        grayArrays.push(new Array(config['outWidth']))
    }
    let skip=config['bPointNum']+1
    for (let h=0;h<config['outHeight'];h+=skip) {
        for (let w=0;w<config['outWidth'];w+=skip){
            grayArrays[h][w]=basisArrays[h/skip][w/skip]
            if (w!=0) {
                let d=grayArrays[h][w-skip]- grayArrays[h][w]
                if (d<0){
                    d=-d
                }
                for (let ws=1;ws<skip;ws++){
                    if (d==0){
                        grayArrays[h][w -skip+ws] = grayArrays[h][w]
                    }else if (d>skip){
                        grayArrays[h][w -skip+ws]=differenceArray[d_index++]
                    }else {
                        if (grayArrays[h][w - skip] > grayArrays[h][w]) {
                            if (ws < d) {
                                grayArrays[h][w -skip + ws] = grayArrays[h][w -skip] - ws
                            } else {
                                grayArrays[h][w -skip + ws] = grayArrays[h][w -skip] - d
                            }
                        } else {
                            if (ws < d) {
                                grayArrays[h][w -skip + ws] = grayArrays[h][w -skip] + ws
                            } else {
                                grayArrays[h][w -skip + ws] = grayArrays[h][w -skip] + d
                            }
                        }
                    }
                }
            }
            if (h!=0&&w!=0) {
                for (let ws=0;ws<skip;ws++)
                {
                    let d = grayArrays[h-skip][w-skip+ws] - grayArrays[h][w-skip+ws]
                    if (d < 0) {
                        d = -d
                    }
                    for (let hs = 1; hs < skip; hs++) {
                        if (d == 0) {
                            grayArrays[h - skip + hs][w-skip + ws] = grayArrays[h - skip][w-skip + ws]
                        } else if (d > skip) {
                            grayArrays[h - skip + hs][w-skip + ws] = differenceArray[d_index++]
                        } else {
                            if (grayArrays[h - skip][w-skip + ws] > grayArrays[h][w-skip+ ws]) {
                                if (hs < d) {
                                    grayArrays[h - skip + hs][w-skip + ws] = grayArrays[h - skip][w-skip + ws] - hs
                                } else{
                                    grayArrays[h - skip + hs][w-skip + ws] = grayArrays[h - skip][w-skip + ws] - d
                                }
                            } else {
                                if (hs < d) {
                                    grayArrays[h - skip + hs][w-skip + ws] = grayArrays[h - skip][w-skip + ws] + hs
                                } else{
                                    grayArrays[h - skip + hs][w-skip + ws] = grayArrays[h - skip][w-skip + ws] + d
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    return grayArrays
}