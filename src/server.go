package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/", index) // index 为向 url发送请求时，调用的函数
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	filePath:=path.Join(dir,"./src/html/gp.html")
	fmt.Println(filePath)
	content, _ := ioutil.ReadFile(filePath)
	w.Write(content)
}