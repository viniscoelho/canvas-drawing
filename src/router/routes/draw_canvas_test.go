package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"canvas-drawing/src/types"
	"canvas-drawing/src/types/common"
	typesmocks "canvas-drawing/src/types/mocks"
)

func TestDrawCanvas_Success(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	canvasMock := typesmocks.NewMockCanvasDrawing(ctrl)

	rectDTO := drawOnCanvasAndReturnDTO(canvasMock)
	bodyBytes, err := json.Marshal(rectDTO)
	assert.NoError(err)

	body := bytes.NewReader(bodyBytes)
	req := httptest.NewRequest(http.MethodPost, "/canvas", body)
	rw := httptest.NewRecorder()

	router := newFakeRouter(canvasMock)
	router.ServeHTTP(rw, req)

	resp := rw.Result()
	assert.Equal(http.StatusCreated, resp.StatusCode)
}

func TestDrawCanvas_InvalidRoutes(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		path string
	}{
		{path: "/"},
		{path: "/canvas/"},
		{path: "/canvas/draw"},
		{path: "/canvas/draw/something"},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			body := bytes.NewReader([]byte{})
			req := httptest.NewRequest(http.MethodPost, tc.path, body)
			rw := httptest.NewRecorder()

			router := newFakeRouter(nil)
			router.ServeHTTP(rw, req)

			resp := rw.Result()
			assert.Equal(http.StatusNotFound, resp.StatusCode)
		})
	}
}

func drawOnCanvasAndReturnDTO(canvasMock *typesmocks.MockCanvasDrawing) common.RectangleDTO {
	fill := '.'
	rect := common.Rectangle{
		Location: common.Coordinates{
			X: 0,
			Y: 0,
		},
		Height: 1,
		Width:  1,
		Fill:   &fill,
	}
	canvasMock.EXPECT().
		FillCanvas(rect)

	return common.NewDTOFromRectangle(rect)
}

func newFakeRouter(cd types.CanvasDrawing) *mux.Router {
	r := mux.NewRouter()

	r.Path("/canvas").
		Methods(http.MethodPost).
		Name("FakeDrawCanvas").
		Handler(NewDrawCanvasHandler(cd))
	r.Path("/canvas").
		Methods(http.MethodGet).
		Name("FakeReadCanvas").
		Handler(NewReadCanvasHandler(cd))

	return r
}
