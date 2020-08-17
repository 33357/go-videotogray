function decodeGip(binArray, config){
    console.log("decodeGP")
    return decodeIPage(binArray,config)
}

function decodeIPage(binArray,config){
    this.grayArray =[]
    this.difference_index = 0
    for (let i=0;i<config['outHeight'];i++){
        this.grayArray.push(new Array(config['outWidth']))
    }
    let maxRowSkip=config['MaxBRowNum']+1
    let maxColumnSkip=config['MaxBColumnNum']+1
    let array=getBasisArrays(binArray,config,maxRowSkip,maxColumnSkip)
    this.basisArray=array[0]
    this.differenceArray=array[1]
    for(let h=0;h<config['outHeight'];h+=maxColumnSkip){
        for(let w=maxRowSkip;w<config['outWidth'];w+=maxRowSkip){
            this.grayArray[h][w-maxRowSkip]=this.basisArray[h/maxColumnSkip][w/maxRowSkip]
            decodeBRow(this.grayArray[h][w-maxRowSkip],this,maxRowSkip,h,w)
            if(w+maxRowSkip>=config['outWidth']){
                if(w==config['outWidth']-1){
                    this.grayArray[h][w]=this.basisArray[h/maxColumnSkip][w/maxRowSkip]
                }else{
                    this.grayArray[h][w]=this.basisArray[h/maxColumnSkip][w/maxRowSkip]
                    decodeBRow(this.grayArray[h][w],this.grayArray[h][config['outWidth']-1],this,h,w)
                }
            }
        }
        if(h != 0){
            decodeBColumn(this.grayArray[h-maxColumnSkip],this.grayArray[h],this)
            if(h+maxColumnSkip>=config['outHeight']&&h!=config['outHeight']-1){
                decodeBColumn(this.grayArray[h],this.grayArray[config['outHeight']-1],this)
            }
        }
    }
    return this.grayArray
}


function getBasisArrays(binArray,config,maxRowSkip,maxColumnSkip){
    let basisArrayHeight
    let basisArrayWidth
    let basisArray=[]
    let otherHeight=config['outHeight']-1
    let otherWidth=config['outWidth']-1
    if(otherHeight%maxColumnSkip==0){
        basisArrayHeight=Math.floor(otherHeight/maxColumnSkip)+1
    }else{
        basisArrayHeight=Math.floor(config['outHeight']/maxColumnSkip)+2
    }
    if(otherWidth%maxRowSkip==0){
        basisArrayWidth=Math.floor(otherWidth/maxRowSkip)+1
    }else{
        basisArrayWidth=Math.floor(config['outWidth']/maxRowSkip)+2
    }
    for(let i=0;i<basisArrayHeight;i++){
        basisArray.push(binArray.slice(i*basisArrayWidth,(i+1)*basisArrayWidth))
    }
    return [basisArray,binArray.slice(basisArray.length)]
}

function decodeBRow(beforePoint,afterPoint,object,rowSkip,h,w) {
    let pd=beforePoint-afterPoint
    if(pd<0){
        pd=-pd
    }
    let rw=w-rowSkip
    for(let rs=0;rs<rowSkip;rs++){
        if(pd>rowSkip){
            object.grayArray[h][rw+rs] = object.differenceArray[object.difference_index++]
        }else if(pd==0) {
            object.grayArray[h][rw+rs]=beforePoint
        }else{
            if(beforePoint>afterPoint){
                if(rs<pd){
                    object.grayArray[h][rw+rs]=object.grayArray[h][rw]-rs
                }else{
                    object.grayArray[h][rw+rs]=object.grayArray[h][rw]-pd
                }
            }else{
                if(rs<pd){
                    object.grayArray[h][rw+rs]=object.grayArray[h][rw]+rs
                }else{
                    object.grayArray[h][rw+rs]=object.grayArray[h][rw]+pd
                }
            }
        }
    }
}

function decodeBColumn(beforeRow,afterRow,object,columnSkip,h,w) {
    let len=beforeRow.length
    for (let cs=0;cs<len;cs++){
        let cd=beforeRow[cs]- afterRow[cs]
        if(cd<0){
            cd=-cd
        }
        for (let cs = 1; cs < columnSkip; cs++) {
            if (cd == 0) {
                object.grayArray[h - columnSkip + cs][w-skip + ws] = object.grayArray[h - skip][w-skip + ws]
            } else if (cd > columnSkip) {
                object.grayArray[h - columnSkip + cs][w-skip + ws] = object.differenceArray[object.difference_index++]
            } else {
                if (object.grayArray[h - skip][w-skip + ws] > object.grayArray[h][w-skip+ ws]) {
                    if (cs < cd) {
                        object.grayArray[h - skip + hs][w-skip + ws] = object.grayArray[h - skip][w-skip + ws] - cs
                    } else{
                        object.grayArray[h - skip + hs][w-skip + ws] = object.grayArray[h - skip][w-skip + ws] - cd
                    }
                } else {
                    if (cs < cd) {
                        object.grayArray[h - skip + hs][w-skip + ws] = object.grayArray[h - skip][w-skip + ws] + cs
                    } else{
                        object.grayArray[h - skip + hs][w-skip + ws] = object.grayArray[h - skip][w-skip + ws] + cd
                    }
                }
            }
        }
    }
}