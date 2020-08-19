function decodeGip(byteArray, config){
    console.log('decodeGip')
    this.grayArray =[]
    this.byteArrayIndex = 0
    this.byteArray=byteArray
    for (let i=0;i<config['outWidth'];i++){
        this.grayArray.push(new Array(config['outHeight']))
    }
    let maxRowSkip=config['maxBRowNum']+1
    let maxColumnSkip=config['maxBColumnNum']+1;
    for(let w=0;;w+=maxRowSkip){
        if(w>config['outWidth']-1) {
            maxRowSkip=config['outWidth']-1-(w-maxRowSkip)
            w=config['outWidth']-1
        }
        for(let h=0;h<config['outHeight']-1;h+=maxColumnSkip){
            this.grayArray[w][h]=this.byteArray[this.byteArrayIndex++]
            if(h!=0){
                decodeIPageColumn(this.grayArray[w][h-maxColumnSkip],this.grayArray[w][h],this,maxColumnSkip,w,h-maxColumnSkip)
                if(h+maxColumnSkip>=config['outHeight']-1){
                    this.grayArray[w][config['outHeight']-1]=this.byteArray[this.byteArrayIndex++]
                    decodeIPageColumn(this.grayArray[w][h],this.grayArray[w][config['outHeight']-1],this,config['outHeight']-1-h,w,h)
                }
            }
        }
        if(w!=0){
            decodeIPageRow(this.grayArray[w-maxRowSkip],this.grayArray[w],this,maxRowSkip,w-maxRowSkip)
            if(w==config['outWidth']-1){
                break
            }
        }
    }
    return this.grayArray
}

function getIPageBasisArrays(binArray,config,maxRowSkip,maxColumnSkip){
    let basisArrayHeight
    let basisArrayWidth
    let otherHeight=config['outHeight']-1
    let otherWidth=config['outWidth']-1
    if(otherHeight%maxColumnSkip==0){
        basisArrayHeight=Math.floor(otherHeight/maxColumnSkip)+1
    }else{
        basisArrayHeight=Math.floor(otherHeight/maxColumnSkip)+2
    }
    if(otherWidth%maxRowSkip==0){
        basisArrayWidth=Math.floor(otherWidth/maxRowSkip)+1
    }else{
        basisArrayWidth=Math.floor(otherWidth/maxRowSkip)+2
    }
    return [binArray.slice(0,basisArrayWidth*basisArrayHeight),binArray.slice(basisArrayWidth*basisArrayHeight)]
}

function decodeIPageColumn(beforeColumnPoint,afterColumnPoint,object,columnSkip,w,ch) {
    let cd=beforeColumnPoint-afterColumnPoint
    if(cd<0){
        cd=-cd
    }
    for(let cs=1;cs<columnSkip;cs++){
        let h=ch+cs
        if(cd>columnSkip){
            object.grayArray[w][h]=object.byteArray[object.byteArrayIndex++]
        }else if(cd==0) {
            object.grayArray[w][h]=beforeColumnPoint
        }else{
            if(beforeColumnPoint>afterColumnPoint){
                if(cs<cd){
                    object.grayArray[w][h]=beforeColumnPoint-cs
                }else{
                    object.grayArray[w][h]=beforeColumnPoint-cd
                }
            }else{
                if(cs<cd){
                    object.grayArray[w][h]=beforeColumnPoint+cs
                }else{
                    object.grayArray[w][h]=beforeColumnPoint+cd
                }
            }
        }
    }
}

function decodeIPageRow(beforeRowColumn,afterRowColumn,object,rowSkip,rw) {
    let len=beforeRowColumn.length
    for(let h=0;h<len;h++){
        let rd=beforeRowColumn[h]-afterRowColumn[h]
        if(rd<0){
            rd=-rd
        }
        for(let rs=1;rs<rowSkip;rs++){
            let w=rw+rs
            if(rd>rowSkip){
                object.grayArray[w][h] = object.byteArray[object.byteArrayIndex++]
            }else if(rd==0){
                object.grayArray[w][h] = beforeRowColumn[h]
            }else{
                if (beforeRowColumn[h] > afterRowColumn[h]) {
                    if (rs < rd) {
                        object.grayArray[w][h] =beforeRowColumn[h] - rs
                    } else{
                        object.grayArray[w][h] =beforeRowColumn[h] - rd
                    }
                } else {
                    if (rs < rd) {
                        object.grayArray[w][h] =beforeRowColumn[h] + rs
                    } else{
                        object.grayArray[w][h] =beforeRowColumn[h] + rd
                    }
                }
            }
        }
    }
}