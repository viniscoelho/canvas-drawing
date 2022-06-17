package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"canvas-drawing/src/types"
	"canvas-drawing/src/types/common"
)

type drawCanvas struct {
	cd types.CanvasDrawing
}

func NewDrawCanvasHandler(cd types.CanvasDrawing) *drawCanvas {
	return &drawCanvas{cd}
}

func (h *drawCanvas) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	rectDTO := common.RectangleDTO{}
	err = json.Unmarshal(body, &rectDTO)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	rect, err := common.NewRectangleFromDTO(rectDTO)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
		message := fmt.Sprintf("one or more fields do not match the requirements: %s", err.Error())
		rw.Write([]byte(message))
		return
	}

	err = h.cd.FillCanvas(rect)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))

		return
	}

	rw.WriteHeader(http.StatusCreated)
}
