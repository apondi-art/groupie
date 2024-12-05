package models

import "sync"

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
