package handlers

import (
	"html/template"
	"net/http"

	"GROUPIE-TRACKER/internal/api"
)

// prepare handler function permite switch enter the all page of the website
func Handler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			Home(w, r)
			return
		}
		NotFound(w, r)
	})
	http.HandleFunc("/artist", Artist)
	http.HandleFunc("/artists", Artists)
	http.HandleFunc("/locations", Locations)
	http.HandleFunc("/gimstroll", Gimstroll)
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

// init template index.html
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

// init template artist.html
func Artist(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "artist.html")
}

// init template artists.html
func Artists(w http.ResponseWriter, r *http.Request) {
	artists, err := api.GetArtists()
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(
		"web/templates/artists.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, artists)
}

// init template locations.html
func Locations(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "locations.html")
}

// init template gimstroll.html
func Gimstroll(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "gimstroll.html")
}

func renderTemplate(w http.ResponseWriter, filename string) {
	t, err := template.ParseFiles("web/templates/" + filename)
	if err != nil {
		http.Error(w, "Erreur de template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
