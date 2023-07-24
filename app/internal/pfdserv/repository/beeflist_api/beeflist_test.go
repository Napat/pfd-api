package beeflist_api_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/Napat/pfd-api/app/internal/pfdserv/config"
	"github.com/Napat/pfd-api/app/internal/pfdserv/model"
	"github.com/Napat/pfd-api/app/internal/pfdserv/repository/beeflist_api"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"time"
)

func TestGetList(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		ctx := context.Background()
		cfg := &config.BeeflistAdaptor{
			Url:     "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text",
			Timeout: time.Duration(3 * time.Second),
			Retry: config.Retry{
				MaxRetries:       3,
				WaitTime:         time.Duration(1 * time.Second),
				BaseRetryBackoff: time.Duration(2 * time.Second),
				MaxWaitTime:      time.Duration(20 * time.Second),
			},
		}

		cassette, err := recorder.New("testdata/beef/success_getlist")
		if err != nil {
			t.Fatalf("Failed to create recorder: %v", err)
		}
		defer cassette.Stop()

		client := resty.New()
		client.SetTransport(cassette)
		blAdaptor := beeflist_api.NewBeefListAdaptor(cfg, client)

		gotResp, err := blAdaptor.GetList(ctx)
		if err != nil {
			t.Fatalf("GetList Failed to make request: %v", err)
		}

		expectResp := model.BeeflistAdaptorResponse{
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
		}

		if !compareMaps(expectResp.BeefSummary, gotResp.BeefSummary) {
			t.Errorf("Unexpected response body. Expected '%v', got '%v'", expectResp, gotResp)
		}
	})

	t.Run("Sad path: call wrong url scheme", func(t *testing.T) {
		ctx := context.Background()
		cfg := &config.BeeflistAdaptor{
			Url:     "xxx://localhost",
			Timeout: time.Duration(3 * time.Second),
			Retry: config.Retry{
				MaxRetries:       3,
				WaitTime:         time.Duration(1 * time.Second),
				BaseRetryBackoff: time.Duration(2 * time.Second),
				MaxWaitTime:      time.Duration(20 * time.Second),
			},
		}

		client := resty.New()
		blAdaptor := beeflist_api.NewBeefListAdaptor(cfg, client)
		_, err := blAdaptor.GetList(ctx)
		assert.Error(t, err)
	})

	t.Run("Sad path: error 500", func(t *testing.T) {
		ctx := context.Background()
		cfg := &config.BeeflistAdaptor{
			Url: "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text",
		}

		cassette, err := recorder.New("testdata/beef/error_getlist_500")
		if err != nil {
			t.Fatalf("Failed to create recorder: %v", err)
		}
		defer cassette.Stop()

		client := resty.New()
		client.SetTransport(cassette)
		blAdaptor := beeflist_api.NewBeefListAdaptor(cfg, client)

		_, err = blAdaptor.GetList(ctx)
		assert.Error(t, err)
	})
}

func TestProcessRawBeef(t *testing.T) {
	input := `tIp. ,ut ipsUm magna ut. cupim

Ut cupim chop ut, 
sunt.`

	beefCount := map[string]int{
		"tip":   1,
		"ipsum": 1,
		"magna": 1,
		"ut":    4,
		"chop":  1,
		"sunt":  1,
		"cupim": 2,
	}

	expect := map[string]map[string]int{
		"beef": beefCount,
	}

	got := beeflist_api.SpliterCounterResp(input)

	if !compareMaps(expect, got) {
		t.Errorf("Unexpected response body. Expected '%v', got '%v'", expect, got)
	}
}

func compareMaps(m1, m2 map[string]map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok {
			return false
		}
		if !reflect.DeepEqual(v1, v2) {
			return false
		}
	}
	return true
}
