package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Handler method.
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	filename := strings.Split(r.URL.EscapedPath(), "/")[1]
	switch m {
	case http.MethodGet:
		getHandler(w, filename)
	case http.MethodPut:
		putHandler(w, filename)
	}
}

func getHandler(w http.ResponseWriter, filename string) {
	fileReader, err := os.Open("storage/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileReader.Close()

	io.Copy(w, fileReader)
}

func putHandler(w http.ResponseWriter, filename string) {
	//todo
}
