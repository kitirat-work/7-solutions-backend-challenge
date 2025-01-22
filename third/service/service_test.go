package service_test

import (
	"context"
	"errors"
	"piefiredire/baconipsum/mocks"
	"piefiredire/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_BeefSummary(t *testing.T) {
	var (
		mockBaconipsumService mocks.PieFireDireBaconipsum
		srv                   service.PieFireDireService

		ctx context.Context
	)

	beforeEach := func() {
		mockBaconipsumService = mocks.PieFireDireBaconipsum{}
		srv = service.NewService(&mockBaconipsumService)

		ctx = context.Background()

	}

	t.Run("should get info from baconipsum service", func(t *testing.T) {
		beforeEach()
		mockBaconipsumService.On("Get").
			Return("Fatback t-bone t-bone, pastrami .. t-bone. pork, meatloaf jowl enim. Bresaola t-bone.", nil)

		_, err := srv.BeefSummary(ctx)

		assert.NoError(t, err)
		mockBaconipsumService.AssertNumberOfCalls(t, "Get", 1)
	})

	t.Run("should return error if baconipsum service fails", func(t *testing.T) {
		beforeEach()
		mockBaconipsumService.On("Get").
			Return("", errors.New("error"))

		_, err := srv.BeefSummary(ctx)

		assert.Error(t, err)
		mockBaconipsumService.AssertNumberOfCalls(t, "Get", 1)
	})

	t.Run("should count beef types", func(t *testing.T) {
		beforeEach()
		mockBaconipsumService.On("Get").
			Return("Fatback t-bone t-bone, pastrami .. t-bone. pork, meatloaf jowl enim. Bresaola t-bone.", nil)

		summary, err := srv.BeefSummary(ctx)

		assert.NoError(t, err)
		assert.Equal(t, service.BeefSummary{
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
		}, summary)
	})

}
