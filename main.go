package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var hello = "Hello World!"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, hello)
}

func indexHtmlHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(content))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index.html", indexHtmlHandler)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
