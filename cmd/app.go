package main

import (
	"fmt"
	"http-polygon/internal/api"
	"http-polygon/pkg/polygon"
	"log"
	"net/http"
)

const (
	Port = 8080
)

func main() {
	standardDrawService := polygon.NewStandardDrawService()
	concurrentDrawService := polygon.NewConcurrentDrawService(25)

	handlerWithStandardDraw := api.NewDrawPolygonHandler(standardDrawService)
	handlerWithConcurrentDraw := api.NewDrawPolygonHandler(concurrentDrawService)

	http.HandleFunc("/draw-polygon", handlerWithStandardDraw.Handle)
	http.HandleFunc("/v2/draw-polygon", handlerWithConcurrentDraw.Handle)

	log.Printf("starting server on port %d\n", Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil); err != nil {
		log.Fatal("can't listen and serve", err)
	}
}
