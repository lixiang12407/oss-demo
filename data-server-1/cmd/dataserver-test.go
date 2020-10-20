package main

import (
	"fmt"
	"net/http"
	"oss-demo/data-server-1/handler"
)

func main() {
	fmt.Println("Data server started.")
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe("localhost:8081", nil)
}
