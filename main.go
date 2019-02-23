package main

import "net/http"

// for local.
func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
