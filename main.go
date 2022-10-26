package main

import (
	"fmt"
	"net/http"
)

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func main() {
	http.HandleFunc("/", handlerfunc)
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
