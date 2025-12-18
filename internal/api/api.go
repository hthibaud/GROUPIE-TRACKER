package api

import (
	"fmt"
	"io"
	"net/http"
)

func getApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {

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

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Voici les items disponibles.")
}
