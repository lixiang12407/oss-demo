package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"oss-demo/data-server-1/tools"
)

// Handler method.
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	switch m {
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPut:
		putHandler(w, r)
	case http.MethodDelete:
		delHandler(w, r)
	}
}

func delHandler(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.EscapedPath(), "/")[2]
	fmt.Println(filename)
	err := os.Remove("storage/" + filename)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = http.Get("http://localhost:8090/deletefile/" + filename)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 回复客户端
	w.WriteHeader(http.StatusOK)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.EscapedPath(), "/")[1]
	fileReader, err := os.Open("storage/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileReader.Close()

	io.Copy(w, fileReader)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	// 文件存储
	filename := strings.Split(r.URL.EscapedPath(), "/")[1]
	fmt.Println(filename)
	fileWriter, err := os.Create("storage/" + filename)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer fileWriter.Close()
	_, err = io.Copy(fileWriter, r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 向接口服务器发送元数据
	fileCrypto := tools.FileCrypto(filename, "sha1")
	urlSuffix := "?filename=" + filename + "&filehash=" + fileCrypto + "&location=http://localhost:8081"
	_, err = http.Get("http://localhost:8090/updatemeta/" + urlSuffix)
	// 与接口服务器通信失败
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 回复客户端
	w.WriteHeader(http.StatusOK)
}
