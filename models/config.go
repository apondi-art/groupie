package models

import (
	"io"
	"net/http"
	"sync"
)

type Config struct {
	Apis *Apis
}

//idiomatic go when developing different applications  backend

// Declare global variables
var (
	instance *Config
	once     sync.Once
)

func AppConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Apis: NewApis(),
		}
	})
	return instance
}

func (c *Config) FetchData(apiType string) ([]byte, error) {
	res, err := http.Get(apiType)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
