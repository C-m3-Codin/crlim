package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type RateLimitConfig struct {
	RequestsPerSecond float64 `json:"requests_per_second" yaml:"requests_per_second"`
	BurstSize         int     `json:"burst_size" yaml:"burst_size"`
}

type Config struct {
	RateLimits map[string]RateLimitConfig `json:"rate_limits" yaml:"rate_limits"`
}

// LoadConfig loads a JSON or YAML config file
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if strings.HasSuffix(filename, ".json") {
		err = json.Unmarshal(data, config)
	} else if strings.HasSuffix(filename, ".yaml") || strings.HasSuffix(filename, ".yml") {
		err = yaml.Unmarshal(data, config)
	} else {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return config, nil
}
