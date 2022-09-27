package model

type Article struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	ImageURL    string `json:"imageUrl"`
	Description string `json:"description"`
}
