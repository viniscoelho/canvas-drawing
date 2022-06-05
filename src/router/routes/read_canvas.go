package routes

import (
	"log"
	"net/http"

	"exercise/src/types"
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
