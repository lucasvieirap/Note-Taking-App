package controllers

import (
	"html/template"
	"net/http"

	"github.com/lucasvieirap/Note-Taking-App/internal/util"
)

// Welcome user and Display Usability
func HandleHomePageRouter(w http.ResponseWriter, r *http.Request) {
	absolutePath := util.GetPath()
	filePath := absolutePath + "/web/index.html"

	htmlTemplate, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	htmlTemplate.Execute(w, nil)
}
