package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func NewConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Apis: ConfigApis(),
		}

		go instance.CreateInitialData()

	})
	return instance
}

func (c *Config) CreateInitialData() {
	body, err := c.FetchData("artists")
	if err != nil {
		log.Fatal(err)

	}

	// parse json response
	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(data) > 0 {

		fmt.Printf("\n\nThe format of our data\n%v\n\n", data[0])
	}

}

func (c *Config) FetchData(apiType string) ([]byte, error) {
	res, err := http.Get(c.Apis.ApisUrls[apiType])
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
