package service_test

//go:generate mockgen -source=./beef.go -destination=./mock_service/mock_beef.go -package=mock_service

import (
	"context"
	"errors"
	"testing"

	"github.com/Napat/pfd-api/internal/model"
	"github.com/Napat/pfd-api/internal/repository/beeflist_api"
	mock_beeflist_api "github.com/Napat/pfd-api/internal/repository/beeflist_api/mock_service"
	"github.com/Napat/pfd-api/internal/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSummary(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	t.Run("happy path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockBeefBeefListAdaptor := mock_beeflist_api.NewMockIBeefListAdaptor(ctrl)
		mockBeefBeefListAdaptor.EXPECT().GetList(ctx).Return(&model.BeeflistAdaptorResponse{
			BeefSummary: map[string]map[string]int{
				"beef": {
					"tip":   1,
					"ipsum": 1,
					"magna": 1,
					"ut":    4,
					"chop":  1,
					"sunt":  1,
					"cupim": 2,
				},
			},
		}, nil)

		blasvc := service.NewBeefService(mockBeefBeefListAdaptor)
		resp, err := blasvc.Summary(ctx)

		assert.NoError(t, err)
		assert.Equal(t, &model.BeefSummaryResponse{
			BeefSummary: map[string]map[string]int{
				"beef": {
					"tip":   1,
					"ipsum": 1,
					"magna": 1,
					"ut":    4,
					"chop":  1,
					"sunt":  1,
					"cupim": 2,
				},
			},
		}, resp)
	})

	t.Run("sad path: got some error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockBeefBeefListAdaptor := mock_beeflist_api.NewMockIBeefListAdaptor(ctrl)
		mockBeefBeefListAdaptor.EXPECT().GetList(ctx).Return(nil, errors.New("some error"))

		blasvc := service.NewBeefService(mockBeefBeefListAdaptor)
		_, err := blasvc.Summary(ctx)

		assert.Error(t, err)
	})
}

type IBeefService interface {
	Summary(ctx context.Context) (*model.BeefSummaryResponse, error)
}

type BeefService struct {
	beefListAdaptor beeflist_api.IBeefListAdaptor
}

func NewBeefService(
	beefListAdaptor beeflist_api.IBeefListAdaptor,
) *BeefService {
	return &BeefService{
		beefListAdaptor: beefListAdaptor,
	}
}

func (s *BeefService) Summary(ctx context.Context) (*model.BeefSummaryResponse, error) {
	beefListResp, err := s.beefListAdaptor.GetList(ctx)
	if err != nil {
		return nil, err
	}

	resp := &model.BeefSummaryResponse{
		BeefSummary: beefListResp.BeefSummary,
	}

	return resp, nil
}
