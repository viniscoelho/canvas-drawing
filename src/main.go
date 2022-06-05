package main

import (
	"log"
	"net/http"
	"time"

	"exercise/src/router"
	"exercise/src/types/canvas"
)

const (
	height = 100
	width  = 100
)

func main() {
	cd, err := canvas.NewCanvasDrawing(height, width)
	if err != nil {
		log.Fatal("could not initialize storage")
	}

	s := &http.Server{
		Handler:      router.CreateRoutes(cd),
		ReadTimeout:  0,
		WriteTimeout: 0,
		Addr:         ":3000",
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(s.ListenAndServe())
}
