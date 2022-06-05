package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"exercise/src/router"
	"exercise/src/types/canvas"
)

var (
	height = flag.Int("height", 50, "height of the canvas")
	width  = flag.Int("width", 50, "width of the canvas")
	port   = flag.Int("port", 3000, "application address port")
)

func main() {
	flag.Parse()

	cd, err := canvas.NewCanvasDrawing(*height, *width)
	if err != nil {
		log.Fatal("could not initialize storage")
	}

	s := &http.Server{
		Handler:      router.CreateRoutes(cd),
		ReadTimeout:  0,
		WriteTimeout: 0,
		Addr:         fmt.Sprintf(":%d", *port),
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(s.ListenAndServe())
}
