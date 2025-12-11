package handlers

import (
	"html/template"
	"net/http"
)

// prepare handler function permite switch enter the all page of the website
func Handler() {
	http.HandleFunc("/", Home)
}

// init template index.html
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, filename string) {
	t, err := template.ParseFiles("web/templates/" + filename)
	if err != nil {
		http.Error(w, "Erreur de template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
