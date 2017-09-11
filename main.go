package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "REQUEST URI: %s\n", r.RequestURI)
		fmt.Fprintf(w, "REQUEST HOST HEADER: %s\n", r.Host)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
