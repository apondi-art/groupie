package main

import (
	"fmt"
	"groupie/handlers"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("views/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", handlers.Homehandler)

	fmt.Println("server starting @http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error creating server", err)
	}
}
