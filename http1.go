package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "helloWorld.html")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
