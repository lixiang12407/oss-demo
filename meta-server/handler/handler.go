package handler

import (
	"net/http"
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
	fileLocation := "http://localhost:8081/" + filename
	w.Header().Set("location", fileLocation)
}

func putHandler(w http.ResponseWriter, filename string) {
	// todo
}
