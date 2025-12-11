package main

import (
	"log"
	"net/http"
)

func main() {
	handlers.Handler()
	log.Println("Server Start on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
