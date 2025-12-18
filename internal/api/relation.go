package api

type Relations struct {
	ID        int   `json:"id"`
	Relations []int `json:"relations"`
}
