package main

import (
	"fmt"
	"net/http"
)

func main() {
	//dir, _ := os.Getwd()
	//filePath:=path.Join(dir,"./src/html/")
	fmt.Println("http://localhost:9000/gp.html")
	http.Handle("/", http.FileServer(http.Dir("./src/html")))
	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("http://localhost:9000/gp.html")
}