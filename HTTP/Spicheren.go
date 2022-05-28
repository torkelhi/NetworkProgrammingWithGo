package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Head("https://spicheren.no/besokstall/")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	_ = resp.Body.Close()

	date := resp.Header.Get("Date")
	if date == "" {
		fmt.Println(date)
		log.Fatalf("Received %s", date)
	}
	fmt.Println(date)
}
