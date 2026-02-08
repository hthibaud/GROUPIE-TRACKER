package handlers

import (
	"GROUPIE-TRACKER/internal/api"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"fmt"
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
	// Récupérer l'id depuis l'URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// calling the API
	artist, err := api.GetArtist(id)
	if err != nil {
		http.Error(w, "Artiste non trouvé", http.StatusInternalServerError)
		return
	}

	// send the data to the HTML page
	tmpl, err := template.ParseFiles("web/templates/artist.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, artist)
}

// init template artists.html
func Artists(w http.ResponseWriter, r *http.Request) {
	artists, err := api.GetArtists()
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	// --- FILTRAGE MINIMALISTE ---
	filtered := []api.Artist{}
	q := r.URL.Query()
	
	for _, a := range artists {
		keep := true

		// Nom
		if s := q.Get("search"); s != "" && !strings.Contains(strings.ToLower(a.Name), strings.ToLower(s)) { keep = false }
		
		// Création (Année)
		if min := q.Get("creation_min"); keep && min != "" {
			if v, _ := strconv.Atoi(min); a.CreationDate < v { keep = false }
		}
		if max := q.Get("creation_max"); keep && max != "" {
			if v, _ := strconv.Atoi(max); a.CreationDate > v { keep = false }
		}

		// 1er Album (Année - extraction simple depuis dd-mm-yyyy)
		if keep && (q.Get("album_min") != "" || q.Get("album_max") != "") {
			parts := strings.Split(a.FirstAlbum, "-")
			if len(parts) == 3 {
				y, _ := strconv.Atoi(parts[2])
				if min := q.Get("album_min"); min != "" {
					if v, _ := strconv.Atoi(min); y < v { keep = false }
				}
				if max := q.Get("album_max"); max != "" {
					if v, _ := strconv.Atoi(max); y > v { keep = false }
				}
			}
		}

		// Membres
		if members := q["members"]; keep && len(members) > 0 {
			found := false
			for _, m := range members {
				if strconv.Itoa(len(a.Members)) == m { found = true; break }
			}
			if !found { keep = false }
		}

		// Lieux (Recherche textuelle simple)
		if loc := q.Get("location"); keep && loc != "" {
			if !strings.Contains(strings.ToLower(fmt.Sprintf("%v", a.Locations)), strings.ToLower(loc)) { keep = false }
		}

		if keep { filtered = append(filtered, a) }
	}
	// ----------------------------

	tmpl, err := template.ParseFiles("web/templates/artists.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// On envoie 'filtered' au lieu de 'artists'
	// Comme c'est toujours une liste, pas besoin de changer le {{range .}} dans le HTML !
	tmpl.Execute(w, filtered)
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
