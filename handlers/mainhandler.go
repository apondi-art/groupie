package handlers

import (
	"encoding/json"
	"groupie/models"
	"net/http"
	"path/filepath"
	"text/template"
)

var app =  models.NewConfig()

func Homehandler(w http.ResponseWriter, r *http.Request) {
	 body,err := app.FetchData("artists")
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusNotFound)
		return
	}

	var artists []models.Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		http.Error(w, "Oops something went wrong: unmarshaling", http.StatusInternalServerError)
	}

	tmplDir := filepath.Join("views/templates", "index.html")
	

	tmpl, err := template.ParseFiles(tmplDir)
	if err != nil {
		http.Error(w, "OOPs something went wrong : partsing templates", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Artists": artists,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "OOPs something went wrong : data", http.StatusInternalServerError)

	}

}
