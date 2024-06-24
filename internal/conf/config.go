package conf

import (
	"os"
	"strconv"
)

type Config struct {
	Port uint16
}

func newConfig() *Config {
	portStr, ok := os.LookupEnv("WOWPOW_PORT")
	if !ok {
		return nil
	}

	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return nil
	}

	return &Config{
		Port: uint16(port),
	}
}
