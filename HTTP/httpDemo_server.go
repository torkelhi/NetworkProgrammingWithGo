package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// GET - forespørsler
func getIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to my server\n")
}

func getGreetings(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Greetings stranger\n")
}

func main() {

	http.HandleFunc("/", getIndex)
	http.HandleFunc("/hello", getGreetings)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
