package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"exercise/src/types/common"
	typesmocks "exercise/src/types/mocks"
)

func TestReadCanvas(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	canvasMock := typesmocks.NewMockCanvasDrawing(ctrl)
	canvasMock.EXPECT().
		GetCanvas().
		Return(common.CanvasString{})

	req := httptest.NewRequest(http.MethodGet, "/canvas", nil)
	rw := httptest.NewRecorder()

	router := newFakeRouter(canvasMock)
	router.ServeHTTP(rw, req)

	resp := rw.Result()
	assert.Equal(http.StatusOK, resp.StatusCode)
}
