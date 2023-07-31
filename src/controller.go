package main

import (
	"encoding/json"
	"fmt"
	"github.com/Logotipiwe/dc_go_auth_lib/auth"
	utils "github.com/logotipiwe/dc_go_utils/src"
	"github.com/logotipiwe/dc_go_utils/src/config"
	"log"
	"net/http"
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

	http.HandleFunc("/api/get-history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		user, err := auth.FetchUserData(r)
		if err != nil {
			w.WriteHeader(403)
			return
		}
		log.Println("User " + user.Id + " is playing history...")
		history, err := getHistory()
		if err != nil {
			handleErrInController(w, err)
			return
		}
		historyDtos := utils.Map(history, func(h History) HistoryDto {
			return h.toDto()
		})
		err = json.NewEncoder(w).Encode(historyDtos)
		if err != nil {
			handleErrInController(w, err)
			return
		}
	})

	pool := NewPool()
	go pool.Start()
	http.HandleFunc("/api/socket/listen-changes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/listen-changes")
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

func handleErrInController(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	fmt.Fprintf(w, "{\"ok\": \"false\", \"err\":\"%s\"}", err.Error())
}
