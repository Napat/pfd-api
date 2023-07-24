package beeflist_api

//go:generate mockgen -source=./beeflist.go -destination=./mock_service/mock_beeflist.go -package=mock_beeflist_api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/Napat/pfd-api/app/internal/pfdserv/config"
	"github.com/Napat/pfd-api/app/internal/pfdserv/model"

	"github.com/go-resty/resty/v2"
)

type IBeefListAdaptor interface {
	GetList(ctx context.Context) (*model.BeeflistAdaptorResponse, error)
}

type BeefListAdaptor struct {
	config *config.BeeflistAdaptor
	client *resty.Client
}

func NewBeefListAdaptor(config *config.BeeflistAdaptor, client *resty.Client) *BeefListAdaptor {

	client.SetRetryCount(config.Retry.MaxRetries)
	client.SetTimeout(config.Timeout)
	client.SetRetryWaitTime(config.Retry.WaitTime)
	client.SetRetryMaxWaitTime(config.Retry.MaxWaitTime)
	client.SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
		requestCount := resp.Request.Attempt
		waitTime := time.Duration(1<<uint(requestCount)) * config.Retry.BaseRetryBackoff
		// fmt.Printf("_______%s_______ SetRetryAfter()........................Waiting for %v before retrying...\n", time.Now(), waitTime)

		return waitTime, nil
	})

	return &BeefListAdaptor{
		config: config,
		client: client,
	}
}

func (r *BeefListAdaptor) GetList(ctx context.Context) (*model.BeeflistAdaptorResponse, error) {
	request := r.client.NewRequest()
	request.SetHeader("Content-Type", "application/json")

	resp, err := request.Get(r.config.Url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		err := fmt.Errorf("status code(%d) != 200", resp.StatusCode())
		log.Println("Summary", err)
		return nil, err
	}

	respString := string(resp.Body())

	return &model.BeeflistAdaptorResponse{
		BeefSummary: spliterCounterResp(respString),
	}, err
}

func spliterCounterResp(rawbeef string) map[string]map[string]int {
	words := strings.FieldsFunc(rawbeef, func(r rune) bool {
		return unicode.IsSpace(r) || r == '.' || r == ',' || r == '\n'
	})

	wordCounts := make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word)
		wordCounts[word]++
	}

	responseMap := map[string]map[string]int{
		"beef": wordCounts,
	}

	return responseMap
}
