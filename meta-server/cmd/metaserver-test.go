package main

import (
	"fmt"
	"httpdemo/oss-demo/meta-server/handler"
	"net/http"
)

func main() {
	// todo
	fmt.Println("Meta server started.")
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe("localhost:8080", nil)
}
