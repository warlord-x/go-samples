package main

import (
	"net/http"
)

const (
	message         = "some sample response"
	contenTypeValue = "text/html; charset=utf-91"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", contenTypeValue) // Tip : this should be written before calling writer.WriteHeader
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte(message))
	})
	http.ListenAndServe(":8080", mux)
}
