package main

import (
	"fmt"
	"httpdemo/oss-demo/data-server-1/handler"
	"net/http"
)

func main() {
	fmt.Println("Data server started.")
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe("localhost:8081", nil)
}
