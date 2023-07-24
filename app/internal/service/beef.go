package service

//go:generate mockgen -source=./beef.go -destination=./mock_service/mock_beef.go -package=mock_service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Napat/pfd-api/app/internal/model"
	"github.com/Napat/pfd-api/app/internal/repository/beeflist_api"
)

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

var Now = time.Now

func (s *BeefService) Summary(ctx context.Context) (*model.BeefSummaryResponse, error) {

	const HOUR_0800AM = 7
	if Now().Hour() < HOUR_0800AM {
		return nil, errors.New("not working hour")
	}

	beefListResp, err := s.beefListAdaptor.GetList(ctx)
	if err != nil {
		log.Println("Summary", err)
		return nil, err
	}

	resp := &model.BeefSummaryResponse{
		BeefSummary: beefListResp.BeefSummary,
	}

	return resp, nil
}
