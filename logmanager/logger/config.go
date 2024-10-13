package logger

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

// Config struct for the logger.
type Config struct {
	LogLevel LogLevel `json:"logLevel"`
}

var (
	config     *Config
	configLock = new(sync.RWMutex)
)

// LoadConfig reads and parses the configuration from a JSON file.
func LoadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	configLock.Lock()
	defer configLock.Unlock()

	return json.Unmarshal(data, &config)
}

// GetConfig safely retrieves the current logging configuration.
func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}
