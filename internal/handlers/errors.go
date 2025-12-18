package handlers

import (
	"net/http"
)

// init template index.html
func NotFound(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "404.html")
}
