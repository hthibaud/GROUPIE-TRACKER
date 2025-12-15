package main

import (
	"fmt"
	"net/http"
    "io"

	handlers "GROUPIE-TRACKER/internal/handlers"
)

func main() {

    resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Statut HTTP :", resp.Status)

    body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lecture :", err)
		return
	}

	fmt.Println(string(body))

	handlers.Handler()
	fmt.Println("Server Start on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
