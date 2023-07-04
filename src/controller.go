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
		user, err := GetUserData(r)
		if err != nil {
			println(err.Error())
			w.WriteHeader(403)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") //TODO only on dev
		if user == nil {
			w.WriteHeader(403)
		} else {
			pixels, err := LoadAllPixels()
			if err != nil {
				log.Fatalln(err)
			}
			dtos := utils.Map(pixels, func(p Pixel) PixelDto {
				return p.toDto()
			})
			err = json.NewEncoder(w).Encode(dtos)
			if err != nil {
				log.Fatalln(err)
			}
		}
	})

	println(fmt.Sprint("Hello, we're up!"))
	err := http.ListenAndServe(":"+env.GetContainerPort(), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
