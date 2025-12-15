package artists

func GetArtists(w http.Response.Writer, r *http.request) {
if r.Method != "GET" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintf(w, "Voici les items disponibles.")
}
