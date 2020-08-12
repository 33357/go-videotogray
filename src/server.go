package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:9000/bin.html")
	fmt.Println("http://localhost:9000/gip.html")
	fmt.Println("http://localhost:9000/gpp.html")
	fmt.Println("http://localhost:9000/gbp.html")
	fmt.Println("http://localhost:9000/gv.html")

	http.Handle("/", http.FileServer(http.Dir("./src/html")))
	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		fmt.Println(err)
	}
}