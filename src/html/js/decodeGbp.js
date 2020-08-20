function decodeGbp(beforePageArray,afterPageArray,byteArray,config,bPageLength,byteArrayIndex){
    console.log("decodeGbp")
    this.grayArrays =[]
    this.byteArrayIndex = byteArrayIndex
    this.byteArray = byteArray
    for (let i=0;i<bPageLength;i++){
        this.grayArrays.push(new Array(config['outWidth']))
        for (let j=0;j<config['outWidth'];j++){
            this.grayArrays[i][j]=new Array(config['outHeight'])
        }
    }
    let maxRowSkip=config['maxBRowNum']+1
    let maxColumnSkip=config['maxBColumnNum']+1;
    for(let w=0;;w+=maxRowSkip){
        if(w>=config['outWidth']-1) {
            w=config['outWidth']-1
        }
        for(let h=0;;h+=maxColumnSkip){
            decodeBPageBasis(beforePageArray[w][h],afterPageArray[w][h],this,w,h,bPageLength)
            if(h!=0){
                decodeBPageColumn(this, w, h-maxColumnSkip,maxColumnSkip,bPageLength)
                if(h+maxColumnSkip>=config['outHeight']-1){
                    decodeBPageBasis(beforePageArray[w][config['outHeight']-1],afterPageArray[w][config['outHeight']-1],this,w,h,bPageLength)
                    decodeBPageColumn(this, w, h,config['outHeight']-1-h,bPageLength)
                    break
                }
            }
        }
        if(w!=0){
            decodeBPageRow(this,w-maxRowSkip,maxRowSkip,bPageLength)
        }
        if(w>=config['outWidth']-1){
            break
        }
    }
    return [this.grayArrays,this.byteArrayIndex]
}

function decodeBPageBasis(beforePageBasisPoint,afterPageBasisPoint,object,w,h,betweenPageLength){
    let pd = beforePageBasisPoint - afterPageBasisPoint
    if(pd < 0){
        pd = -pd
    }
    for(let p=0;p<betweenPageLength;p++) {
        let ps=p+1
        if (pd > betweenPageLength + 1) {
            object.grayArrays[p][w][h] = object.byteArray[object.byteArrayIndex++]
        } else if(pd==0){
            object.grayArrays[p][w][h] = beforePageBasisPoint
        }else{
            if(beforePageBasisPoint > afterPageBasisPoint){
                if(ps < pd){
                    object.grayArrays[p][w][h] = beforePageBasisPoint - ps
                } else {
                    object.grayArrays[p][w][h] = beforePageBasisPoint - pd
                }
            } else {
                if(ps < pd){
                    object.grayArrays[p][w][h] = beforePageBasisPoint + ps
                } else {
                    object.grayArrays[p][w][h] = beforePageBasisPoint + pd
                }
            }
        }
    }
}

function decodeBPageColumn(object,w,ch,columnSkip,betweenPageLength) {
    for(let p=0;p<betweenPageLength;p++){
        let beforeColumnPoint = object.grayArrays[p][w][ch]
        let afterColumnPoint = object.grayArrays[p][w][ch + columnSkip]
        let cd = beforeColumnPoint - afterColumnPoint
        if (cd < 0) {
            cd = -cd
        }
        for (let cs = 1; cs < columnSkip; cs++) {
            let h = ch + cs
            if (cd > columnSkip) {
                object.grayArrays[p][w][h] = object.byteArray[object.byteArrayIndex++]
            } else if (cd == 0) {
                object.grayArrays[p][w][h] = beforeColumnPoint
            } else {
                if (beforeColumnPoint > afterColumnPoint) {
                    if (cs < cd) {
                        object.grayArrays[p][w][h] = beforeColumnPoint - cs
                    } else {
                        object.grayArrays[p][w][h] = beforeColumnPoint - cd
                    }
                } else {
                    if (cs < cd) {
                        object.grayArrays[p][w][h] = beforeColumnPoint + cs
                    } else {
                        object.grayArrays[p][w][h] = beforeColumnPoint + cd
                    }
                }
            }
        }
    }
}

function decodeBPageRow(object,rw,rowSkip,betweenPageLength) {
    for(let p=0;p<betweenPageLength;p++) {
        let beforeRowColumn=object.grayArrays[p][rw]
        let afterRowColumn=object.grayArrays[p][rw+rowSkip]
        let len = beforeRowColumn.length
        for (let h = 0; h < len; h++) {
            let rd = beforeRowColumn[h] - afterRowColumn[h]
            if (rd < 0) {
                rd = -rd
            }
            for (let rs = 1; rs < rowSkip; rs++) {
                let w = rw + rs
                if (rd > rowSkip) {
                    object.grayArrays[p][w][h] = object.byteArray[object.byteArrayIndex++]
                } else if (rd == 0) {
                    object.grayArrays[p][w][h] = beforeRowColumn[h]
                } else {
                    if (beforeRowColumn[h] > afterRowColumn[h]) {
                        if (rs < rd) {
                            object.grayArrays[p][w][h] = beforeRowColumn[h] - rs
                        } else {
                            object.grayArrays[p][w][h] = beforeRowColumn[h] - rd
                        }
                    } else {
                        if (rs < rd) {
                            object.grayArrays[p][w][h] = beforeRowColumn[h] + rs
                        } else {
                            object.grayArrays[p][w][h] = beforeRowColumn[h] + rd
                        }
                    }
                }
            }
        }
    }
}