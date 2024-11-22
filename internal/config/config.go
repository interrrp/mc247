package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Username           string
	Address            string
	RejoinIntervalMins uint
}

func LoadFromEnvironmentVariables() (*Config, error) {
	rejoinIntervalMinsStr := os.Getenv("MC247_REJOIN_INTERVAL_MINS")
	rejoinIntervalMins, err := strconv.ParseUint(rejoinIntervalMinsStr, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid uint MC247_REJOIN_INTERVAL_MINS: %s", rejoinIntervalMinsStr)
	}

	cfg := &Config{
		Username:           os.Getenv("MC247_USERNAME"),
		Address:            os.Getenv("MC247_ADDRESS"),
		RejoinIntervalMins: uint(rejoinIntervalMins),
	}

	if cfg.Username == "" {
		return nil, fmt.Errorf("missing environment variable MC247_USERNAME")
	}
	if cfg.Address == "" {
		return nil, fmt.Errorf("missing environment variable MC247_ADDRESS")
	}

	return cfg, nil
}
