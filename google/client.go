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
}

// CustomSearch ...
func (client *Client) CustomSearch(query url.Values) (*CustomSearchResponse, error) {

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
		return nil, fmt.Errorf("Google API error: `%s`", resp.Error.Message)
	}

	if len(resp.Items) == 0 {
		return nil, fmt.Errorf("no entries found for query: `%s`", query.Get("q"))
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
