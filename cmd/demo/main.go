package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var name string
		if name = r.URL.Path[1:]; name == "" {
			name = "stranger"
		}
		fmt.Fprintf(w, "hello %s!", name)
	})
	log.Fatal(http.ListenAndServe(":9999", nil))
}
