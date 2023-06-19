package server

import (
	"example/hello"
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "%s %s\n", hello.Hello("from CI"), hostname)
}

func Server() {
	http.HandleFunc("/", index)
	fmt.Println("Server started.")
	http.ListenAndServe(":8000", nil)
}
