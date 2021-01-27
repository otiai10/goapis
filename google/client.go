package google

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client represents API Client for Google Search API
type Client struct {
	APIKey               string
	CustomSearchEngineID string
	Referer              string
	HTTPClient           *http.Client

	// Eager search
	Eager bool
}

// CustomSearch ...
// https://developers.google.com/custom-search/v1/reference/rest/v1/cse/list
func (client *Client) CustomSearch(query url.Values) (*CustomSearchResponse, error) {

	if client.HTTPClient == nil {
		client.HTTPClient = http.DefaultClient
	}

	baseURL := "https://www.googleapis.com/customsearch/v1"
	req, err := http.NewRequest("GET", baseURL, nil)
	if client.Referer != "" {
		req.Header.Set("Referer", client.Referer)
	}

	if client.CustomSearchEngineID != "" {
		query.Add("cx", client.CustomSearchEngineID)
	}
	if client.APIKey != "" {
		query.Add("key", client.APIKey)
	}
	req.URL.RawQuery = query.Encode()

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resp := new(CustomSearchResponse)

	if err := json.NewDecoder(res.Body).Decode(resp); err != nil {
		return nil, err
	}

	if resp.Error.Code != 0 {
		return nil, fmt.Errorf("Google API error: %d %s", resp.Error.Code, resp.Error.Message)
	}

	if len(resp.Items) == 0 && client.Eager && query.Get("start") != "0" {
		return client.CustomSearch(client.compromise(query))
	}

	return resp, nil
}

// SearchImage ...
// TODO: add more controllability.
func (client *Client) SearchImage(query string, start int) (*CustomSearchResponse, error) {
	num := 5
	q := url.Values{}
	q.Add("q", query)
	q.Add("searchType", "image")
	q.Add("num", fmt.Sprintf("%d", num))
	q.Add("start", fmt.Sprintf("%d", start))
	return client.CustomSearch(q)
}

// SearchGIF ...
// TODO: add more controllability.
func (client *Client) SearchGIF(keyword string) (*CustomSearchResponse, error) {
	num := 5
	start := 1 // 6, 11, 26, 31, ...
	q := url.Values{}
	q.Add("q", keyword)
	q.Add("searchType", "image")
	q.Add("fileType", "gif")
	q.Add("hq", "animated")
	q.Add("num", fmt.Sprintf("%d", num))
	q.Add("start", fmt.Sprintf("%d", start))
	return client.CustomSearch(q)
}

func (client *Client) compromise(query url.Values) url.Values {
	query.Set("num", "10")
	query.Set("start", "1")
	return query
}
