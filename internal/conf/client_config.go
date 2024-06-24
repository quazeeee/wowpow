package conf

import (
	"os"
	"strconv"
)

type ClientConfig struct {
	Config
	Requests int
	Host     string
}

func NewClientConfig() *ClientConfig {
	c := newConfig()
	if c == nil {
		return nil
	}

	reqStr, ok := os.LookupEnv("WOWPOW_REQUESTS")
	if !ok {
		return nil
	}

	req, err := strconv.ParseInt(reqStr, 10, 32)
	if err != nil {
		return nil
	}

	host, ok := os.LookupEnv("WOWPOW_HOST")
	if !ok {
		return nil
	}

	return &ClientConfig{
		Config:   *c,
		Requests: int(req),
		Host:     host,
	}
}
