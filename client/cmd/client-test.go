package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getObject(filename string) {
	response, err := http.Get("http://localhost:8080/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileLocation := response.Header.Get("location")
	if fileLocation == "" {
		fmt.Println("File not exist.")
		return
	}
	response, err = http.Get(fileLocation)
	fileWriter, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileWriter.Close()
	io.Copy(fileWriter, response.Body)
}

func main() {
	filename := "test.file"
	getObject(filename)
}
