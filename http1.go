package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "helloWorld.html")
}

func handleMyName(w http.ResponseWriter, r *http.Request) {
	myName := r.FormValue("myName")
	w.Write([]byte("My name is: "+ myName))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/myName", handleMyName)
	http.ListenAndServe(":8000", nil)
}
