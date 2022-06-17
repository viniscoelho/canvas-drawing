package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"canvas-drawing/src/types"
	"canvas-drawing/src/types/common"
)

type readCanvas struct {
	cd types.CanvasDrawing
}

func NewReadCanvasHandler(cd types.CanvasDrawing) *readCanvas {
	return &readCanvas{cd}
}

func (h *readCanvas) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	canvas := h.cd.GetCanvas()

	content, err := serializeCanvas(canvas)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(content)
}

func serializeCanvas(canvas common.CanvasString) ([]byte, error) {
	dto := common.CanvasDTO{
		Canvas: canvas,
	}
	return json.Marshal(dto)
}
