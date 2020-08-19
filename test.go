package main

import "fmt"

func main() {
	a:=[3][3] int{{1,2,3},{2,3,4},{4,5,6}}
	fmt.Println(a[0:2][1][1])
}
