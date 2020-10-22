package main

import (
	"fmt"
	"net/http"
	"oss-demo/meta-server/handler"
)

func main() {
	// todo
	fmt.Println("Meta server started.")
	http.HandleFunc("/", handler.Handler)
	http.HandleFunc("/updatemeta/", handler.MetaHandler)
	http.HandleFunc("/deletefile/", handler.DelHandler)
	http.ListenAndServe("localhost:8090", nil)
}
