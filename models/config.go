package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

// Config is a structure that holds configuration settings
// Apis is a pointer to an Apis structure (defined elsewhere)
//initializes all the information needed for a program to run
type Config struct {
	Apis *Apis
}

// Declare global variables for singleton pattern
var (
	// instance will hold the single instance of Config
	instance *Config
	// once ensures the initialization code runs only once
	once sync.Once
)

// NewConfig creates or returns the existing Config instance (Singleton pattern)
func NewConfig() *Config {
	// once.Do ensures this code block executes only once
	once.Do(func() {
		// Create new Config instance with Apis configuration
		instance = &Config{
			Apis: NewApis(),
		}
		// Start a new goroutine (concurrent thread) to create initial data
		go instance.CreateInitialData()
	})
	return instance
}

// CreateInitialData fetches and processes initial data from an API
// Note: There's a syntax error in the receiver (*c **Config), should be (c *Config)
func (c *Config) CreateInitialData() {
	// Fetch data from API endpoint named "artists"
	body, err := c.FetchData("artists")
	if err != nil {
		// If error occurs, log it and terminate program
		log.Fatal(err)
	}

	// Declare slice to hold the parsed JSON data
	// Note: There's a syntax error in string declaration (*string*)
	var data []map[string]interface{}

	// Convert JSON data into Go data structure
	err = json.Unmarshal(body, &data)
	if err != nil {
		// Print error if JSON parsing fails
		fmt.Println(err)
		return
	}

	// If data exists, print the first item to show its format
	// if len(data) > 0 {
	// 	fmt.Printf("\n\nThe format of our data\n%v\n\n", data[0])
	// }
fmt.Printf("%+v",data[0]["members"])
}

// FetchData makes an HTTP GET request to fetch data from an API
// Note: There are syntax errors in the parameter types (*apiType* *string*)
func (c *Config) FetchData(apiType string) ([]byte, error) {
	// Make HTTP GET request using URL from Apis configuration
	res, err := http.Get(c.Apis.ApisUrls[apiType])
	if err != nil {
		return nil, err
	}
	// Ensure response body is closed after function returns
	defer res.Body.Close()

	// Read the entire response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	

	return body, nil
}
