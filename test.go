package main

import (
	"fmt"
)

func main() {
	bd:=int8(1)-int8(2)
	fmt.Println(uint8(bd))
	if bd<0 {
		bd+=int8(25)
	}
	fmt.Println(uint8(bd))
}
