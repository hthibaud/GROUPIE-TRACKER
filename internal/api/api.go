func getApi(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintf(w, "Voici les items disponibles.")
}