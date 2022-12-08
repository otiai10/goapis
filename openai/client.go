package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client: You ain't gonna need it.
// Just use your *http.Client and set headers by yourself.
type Client struct {
	HTTPClient *http.Client
	APIKey     string
	BaseURL    string
}

type Persona string

const (
	Davinci Persona = TextDavinci003
)

// AskDavinci: You ain't gonna need it.
func (client Client) Ask(persona Persona, prompt []string) (response ResponseBody, err error) {
	body := RequestBody{
		Model:     string(persona),
		Prompt:    prompt,
		MaxTokens: 140,
	}
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return response, fmt.Errorf("failed to encode requst body: %v", err)
	}
	req, err := http.NewRequest("POST", client.BaseURL+"/completions", buf)
	if err != nil {
		return response, fmt.Errorf("failed to init request: %v", err)
	}
	if client.HTTPClient == nil {
		client.HTTPClient = http.DefaultClient
	}
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("failed to execute request: %v", err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	return response, err
}
