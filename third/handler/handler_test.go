package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"piefiredire/service"
	"piefiredire/service/mocks"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_BeefSummary(t *testing.T) {
	var (
		mockService *mocks.PieFireDireService
		h           PieFireDireHandler
		e           *echo.Echo

		resBeefSummary service.BeefSummary
		errBeefSummary error
	)

	beforeEach := func() {
		mockService = &mocks.PieFireDireService{}
		h = NewHandler(mockService)
		e = echo.New()
		e.GET("/beef/summary", h.BeefSummary)

		resBeefSummary = service.BeefSummary{}
		errBeefSummary = nil

		mockService.On("BeefSummary", mock.Anything).Return(
			func(context.Context) service.BeefSummary {
				return resBeefSummary
			},
			func(context.Context) error {
				return errBeefSummary
			},
		)
	}

	t.Run("should call BeefSummary from service", func(t *testing.T) {
		beforeEach()
		req := httptest.NewRequest(http.MethodGet, "/beef/summary", nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req.WithContext(context.TODO()))

		mockService.AssertNumberOfCalls(t, "BeefSummary", 1)
	})

	t.Run("should return 200 OK if BeefSummary from service is successful", func(t *testing.T) {
		beforeEach()
		req := httptest.NewRequest(http.MethodGet, "/beef/summary", nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req.WithContext(context.TODO()))

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return 500 Internal Server Error if BeefSummary from service fails", func(t *testing.T) {
		beforeEach()
		errBeefSummary = errors.New("error")

		req := httptest.NewRequest(http.MethodGet, "/beef/summary", nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req.WithContext(context.TODO()))

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("should return BeefSummary from service", func(t *testing.T) {
		beforeEach()
		resBeefSummary = service.BeefSummary{
			Beef: map[string]int{
				"fatback":  1,
				"t-bone":   4,
				"pastrami": 1,
				"pork":     1,
				"meatloaf": 1,
				"jowl":     1,
				"enim":     1,
				"bresaola": 1,
			},
		}

		req := httptest.NewRequest(http.MethodGet, "/beef/summary", nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req.WithContext(context.TODO()))

		assert.Equal(t, http.StatusOK, rec.Code)
		actual := service.BeefSummary{}
		err := json.NewDecoder(rec.Body).Decode(&actual)
		assert.NoError(t, err)
		assert.Equal(t, resBeefSummary, actual)
	})

}
