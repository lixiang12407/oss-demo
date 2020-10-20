package handler

import (
	"fmt"
	"net/http"
	"oss-demo/meta-server/database"
	"strings"
)

// Handler method.
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	switch m {
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPut:
		putHandler(w, r)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.EscapedPath(), "/")[1]
	// 查询数据库
	mydb := database.GetDatabase()
	stmt, err := mydb.Prepare("SELECT location FROM file_meta WHERE filename = ?")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var location string
	err = stmt.QueryRow(filename).Scan(&location)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fileLocation := location + filename
	w.Header().Set("location", fileLocation)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	// todo
	filename := strings.Split(r.URL.EscapedPath(), "/")[1]
	fileLocation := "http://localhost:8081/" + filename
	w.Header().Set("location", fileLocation)
}

// MetaHandler method.
func MetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filename := r.Form.Get("filename")
	filehash := r.Form.Get("filehash")
	location := r.Form.Get("location")
	mydb := database.GetDatabase()

	stmt, err := mydb.Prepare("INSERT INTO file_meta VALUES(?,?,?)")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(filename, filehash, location)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
