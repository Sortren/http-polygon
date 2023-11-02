package main

import (
	"fmt"
	"http-polygon/internal/api"
	"log"
	"net/http"
)

const (
	Port = 8080
)

func main() {
	http.HandleFunc("/draw-polygon", api.DrawPolygonOnFile)

	log.Printf("starting server on port %d\n", Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil); err != nil {
		log.Fatal("can't listen and serve", err)
	}
}
