package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getObject(filename string) {
	response, err := http.Get("http://localhost:8090/" + filename)
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

func putObject(filename string) {
	request, err := http.NewRequest(http.MethodPut, "http://localhost:8090/"+filename, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, _ := http.DefaultClient.Do(request)
	fileLocation := response.Header.Get("location")

	fileReader, _ := os.Open(filename)
	defer fileReader.Close()

	request, err = http.NewRequest(http.MethodPut, fileLocation, fileReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, _ = http.DefaultClient.Do(request)
	fmt.Println(http.StatusText(response.StatusCode))
}

func delObject(filename string) {
	request, err := http.NewRequest(http.MethodDelete, "http://localhost:8090/"+filename, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, _ := http.DefaultClient.Do(request)
	fileLocation := response.Header.Get("location")
	fmt.Println(fileLocation)

	request, _ = http.NewRequest(http.MethodDelete, "http://localhost:8081/deletefile"+filename, nil)
	response, _ = http.DefaultClient.Do(request)
	fmt.Println(http.StatusText(response.StatusCode))
}

func main() {
	filename := "test1.file"
	// putObject(filename)
	delObject(filename)
}
