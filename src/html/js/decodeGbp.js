function decodeGbp(binArray1,binArray2,binArray,config,index){
    console.log("decodeGbP")
    let grayArray=getBPageGrayArray(binArray1,binArray2,binArray,config,index)
    return grayArray
}

function getBPageGrayArray(beforePageArray,afterPageArray,binArrays,config,index) {
    let d_index=0
    let grayArrays=[]
    for (let i=0;i<config['outHeight'];i++){
        grayArrays.push(new Array(config['outWidth']))
    }
    let pointSkip=config['bPointNum']+1
    let pageSkip=config['bPageNum']+1
    for (let h=0;h<config['outHeight'];h+=pointSkip) {
        for (let w=0;w<config['outWidth'];w+=pointSkip){
            let bd=beforePageArray[h][w]-afterPageArray[h][w]
            if (bd<0) {
                bd=-bd
            }
            if (bd>pageSkip){
                for (let ps=0;ps<config['bPageNum'];ps++){
                    grayArrays[ps][h][w] = binArrays[ps][d_index++]
                }
            }else{
                let d=beforePageArray[h][w]-afterPageArray[h][w]
                if (d<0){
                    d=-d
                }
                if (beforePageArray[h][w]>afterPageArray[h][w]) {
                    if(ps<d){
                        reGrayArrays[ps][h][w]=beforePageArray[h][w]-uint8(ps)
                    }else{
                        reGrayArrays[ps][h][w]=beforePageArray[h][w]-uint8(d)
                    }
                }else{
                    if(ps<d){
                        reGrayArrays[ps][h][w]=beforePageArray[h][w]+uint8(ps)
                    }else{
                        reGrayArrays[ps][h][w]=beforePageArray[h][w]+uint8(d)
                    }
                }
            }
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