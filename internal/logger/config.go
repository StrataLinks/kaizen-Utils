package logger

import "encoding/json"
import "os"

// Config defines the structure for logging configuration.
type Config struct {
	LogLevel string `json:"logLevel"`
}

// LoadConfig reads and parses the configuration from a JSON file.
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
