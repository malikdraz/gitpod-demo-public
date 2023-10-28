package main

import (
	"fmt"
	"net/http"
)

func main() {
	webPort := ":8080"
	http.HandleFunc("/", homeHandler)
	if err := http.ListenAndServe(webPort, nil); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		http.Error(w, "Unable to load homepage", http.StatusInternalServerError)
		return
	}
}
