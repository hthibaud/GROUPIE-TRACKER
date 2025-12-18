package main

import (
	"fmt"
	"net/http"

	handlers "GROUPIE-TRACKER/internal/handlers"
)

func main() {

	handlers.Handler()
	fmt.Println("Server Start on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
