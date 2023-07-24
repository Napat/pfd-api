package config

import (
	"time"
)

type (
	Config struct {
		ServerAddress   string          `mapstructure:"server-address" validate:"required"`
		BeeflistAdaptor BeeflistAdaptor `mapstructure:"beeflist-svc" validate:"required"`
	}
	BeeflistAdaptor struct {
		Url     string        `mapstructure:"url" validate:"required"`
		Timeout time.Duration `mapstructure:"timeout" validate:"required"`
		Retry   Retry         `mapstructure:"retry" validate:"required"`
	}
	Retry struct {
		MaxRetries       int           `mapstructure:"max-retries" validate:"required"`
		WaitTime         time.Duration `mapstructure:"wait-time" validate:"required"`
		BaseRetryBackoff time.Duration `mapstructure:"base-retry-backoff" validate:"required"`
		MaxWaitTime      time.Duration `mapstructure:"max-wait-time" validate:"required"`
	}
)
