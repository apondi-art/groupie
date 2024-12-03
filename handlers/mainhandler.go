package handlers

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	tmplDir := filepath.Join("views/templates", "index.html")

	tmpl, err := template.ParseFiles(tmplDir)
	if err != nil {
		http.Error(w, "OOPs something went wrong", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "OOPs something went wrong", http.StatusInternalServerError)
		
	}

}
