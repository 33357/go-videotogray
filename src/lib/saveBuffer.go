package lib

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func ArraySaveAsBufferFile(array [] uint8,path string) error {
	buf,err:=arraytoBuffer(array)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = file.Write(buf.Bytes())
	if err != nil {
		return err
	}
	file.Close()
	fmt.Println("save:"+path)
	return nil
}

func arraytoBuffer(arr [] uint8)(*bytes.Buffer,error){
	buf := new(bytes.Buffer)
	for _,value := range arr{
		err := binary.Write(buf, binary.LittleEndian,value)
		if err != nil {
			return buf,err
		}
	}
	return buf,nil
}
