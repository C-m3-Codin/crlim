package crlim

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config structure for rate limits
type Config struct {
	RateLimits map[string]RateLimitPolicy `json:"rate_limits" yaml:"rate_limits"`
}

// LoadConfig loads rate limits from a JSON or YAML file
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	switch ext := getFileExtension(filename); ext {
	case "json":
		err = json.Unmarshal(data, config)
	case "yaml", "yml":
		err = yaml.Unmarshal(data, config)
	default:
		return nil, fmt.Errorf("unsupported config format: %s", ext)
	}

	if err != nil {
		return nil, err
	}
	return config, nil
}

// Helper to get file extension
func getFileExtension(filename string) string {
	if len(filename) > 4 && filename[len(filename)-5:] == ".json" {
		return "json"
	} else if len(filename) > 5 && (filename[len(filename)-5:] == ".yaml" || filename[len(filename)-4:] == ".yml") {
		return "yaml"
	}
	return ""
}
