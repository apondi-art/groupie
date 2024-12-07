package main

import (
	"encoding/json"
	"fmt"
	"groupie/handlers"
	"groupie/models"
	"log"
	"net/http"
)

func main() {
	app := models.AppConfig()
	body, err := app.FetchData(app.Apis.Artist)
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

	fs := http.FileServer(http.Dir("views/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", handlers.Homehandler)

	fmt.Println("server starting @http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error creating server", err)
	}
}
