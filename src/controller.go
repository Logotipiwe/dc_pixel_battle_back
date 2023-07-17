package main

import (
	"encoding/json"
	"fmt"
	"github.com/Logotipiwe/dc_go_auth_lib/auth"
	utils "github.com/logotipiwe/dc_go_utils/src"
	"github.com/logotipiwe/dc_go_utils/src/config"
	"log"
	"net/http"
	"strconv"
)

func main() {
	config.LoadDcConfig()
	err := InitDb()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api/load-pixels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") //TODO only on dev
		pixels, err := LoadAllPixels()
		if err != nil {
			w.WriteHeader(500)
			println(err.Error())
			return
		}
		dtos := utils.Map(pixels, func(p Pixel) PixelDto {
			return p.toDto()
		})
		err = json.NewEncoder(w).Encode(dtos)
		if err != nil {
			println(err.Error())
		}
	})

	http.HandleFunc("/api/load-colors", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		err := json.NewEncoder(w).Encode(getDefaultColors())
		if err != nil {
			println(fmt.Sprintf("error writing colors %s", err.Error()))
		}
	})

	http.HandleFunc("/api/set-pixel", func(w http.ResponseWriter, r *http.Request) {
		println("/set-pixel")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		user, err := auth.FetchUserData(r)
		if err != nil {
			w.WriteHeader(403)
			println(err.Error())
			return
		}
		color := r.URL.Query().Get("color")
		rowStr := r.URL.Query().Get("row")
		colStr := r.URL.Query().Get("col")
		if colStr == "" || rowStr == "" || !isColorExist(color) {
			w.WriteHeader(400)
			println("Wrong data sent")
			return
		}
		row, err1 := strconv.Atoi(rowStr)
		col, err2 := strconv.Atoi(colStr)
		if err1 != nil || err2 != nil {
			w.WriteHeader(500)
			println(fmt.Sprintf("%s; %s", err1, err2))
			return
		}
		pixel := Pixel{row, col, color, user.Id}
		err = pixel.savePixel()
		if err != nil {
			println("Error updating pixel %s", err.Error())
			w.WriteHeader(500)
			return
		}
	})

	pool := NewPool()
	go pool.Start()
	http.HandleFunc("/api/socket/listen-changes", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	println(fmt.Sprint("Hello, we're up!"))
	port := config.GetConfig("CONTAINER_PORT")
	fmt.Println("Port: " + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
