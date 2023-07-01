package main

import (
	"encoding/json"
	"fmt"
	env "github.com/logotipiwe/dc_go_env_lib"
	utils "github.com/logotipiwe/dc_go_utils/src"
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
		dtos := utils.Map(pixels, func(p Pixel) PixelDto {
			return p.toDto()
		})
		err = json.NewEncoder(w).Encode(dtos)
		if err != nil {
			log.Fatalln(err)
		}
	})

	err := http.ListenAndServe(":"+env.GetContainerPort(), nil)
	if err != nil {
		log.Fatalln(err)
	} else {
		println(fmt.Sprintf("Hello, we're up!"))
	}
}
