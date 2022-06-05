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

	"exercise/src/types"
	"exercise/src/types/common"
	typesmocks "exercise/src/types/mocks"
)

func TestDrawCanvas(t *testing.T) {
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

	fillString := "."
	return common.RectangleDTO{
		Location: &common.CoordinatesDTO{
			X: &rect.Location.X,
			Y: &rect.Location.Y,
		},
		Height: &rect.Height,
		Width:  &rect.Width,
		Fill:   &fillString,
	}
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
