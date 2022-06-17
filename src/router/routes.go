package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"canvas-drawing/src/router/routes"
	"canvas-drawing/src/types"
)

func CreateRoutes(cd types.CanvasDrawing) *mux.Router {
	r := mux.NewRouter()

	r.Path("/canvas").
		Methods(http.MethodPost).
		Name("DrawCanvas").
		Handler(routes.NewDrawCanvasHandler(cd))
	r.Path("/canvas").
		Methods(http.MethodGet).
		Name("ReadCanvas").
		Handler(routes.NewReadCanvasHandler(cd))

	return r
}
