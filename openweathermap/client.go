package openweathermap

import "net/http"

// Client is an API client
type Client struct {
	APIKey     string
	HTTPClient *http.Client
	BaseURL    string
	// cache   interface{}
}

// New ...
func New(apikey string, client ...*http.Client) *Client {
	if len(client) == 0 {
		client = append(client, http.DefaultClient)
	}
	return &Client{
		APIKey:     apikey,
		HTTPClient: client[0],
		BaseURL:    "https://api.openweathermap.org",
	}
}
