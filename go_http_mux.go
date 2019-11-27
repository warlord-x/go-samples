package main

import (
	"net/http"
)

const message = "some sample response"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		writer.Write([]byte(message))
	})
	http.ListenAndServe(":8080", mux)
}
