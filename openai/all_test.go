package openai

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestRequestBody(t *testing.T) {
	body := RequestBody{
		Model:     TextDavinci003,
		Prompt:    []string{"お元気ですか？"},
		MaxTokens: 140,
	}
	Expect(t, body).TypeOf("openai.RequestBody")
}
