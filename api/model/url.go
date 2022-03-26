package model

type ShortenRequest struct {
	Url string `json:"url" validate:"required"`
}

type ShortenResponse struct {
	Url        string `json:"url"`
	ShortenUrl string `json:"short_url"`
}
