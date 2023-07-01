package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		println("hi")
		fmt.Fprintf(w, "hi")
	})

	http.HandleFunc("/load-pixels", func(w http.ResponseWriter, r *http.Request) {
		pixels, err := LoadAllPixels()
		if err != nil {
			log.Fatalln(err)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(pixels)
		if err != nil {
			log.Fatalln(err)
		}
	})

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatalln(err)
	} else {
		println(fmt.Sprintf("Hello, we're up!"))
	}
}
