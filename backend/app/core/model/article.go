package model

type Article struct {
	ID          string   `json:"id"`
	URL         string   `json:"url"`
	Title       string   `json:"title"`
	Thumbnail   string   `json:"thumbnail"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
