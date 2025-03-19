package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go dockerizado!")
	})

	fmt.Println("Running on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
