package openai

// RequestBody: See https://beta.openai.com/docs/api-reference/making-requests
type RequestBody struct {

	// Model: ID of the model to use.
	// You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-model
	Model string `json:"model"`

	// Prompt: The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.
	// Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-prompt
	Prompt []string `json:"prompt"`

	// MaxTokens: The maximum number of tokens to generate in the completion.
	// The token count of your prompt plus max_tokens cannot exceed the model's context length. Most models have a context length of 2048 tokens (except for the newest models, which support 4096).
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-max_tokens
	MaxTokens int `json:"max_tokens,omitempty"`

	// Temperature: What sampling temperature to use. Higher values means the model will take more risks. Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.
	// We generally recommend altering this or top_p but not both.
	// See https://beta.openai.com/docs/api-reference/completions/create#completions/create-temperature
	Temperature int `json:"temperature,omitempty"`

	Suffix           string         `json:"suffix,omitempty"`
	TopP             int            `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	LogProbs         int            `json:"logprobs,omitempty"`
	Echo             bool           `json:"echo,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  int            `json:"presence_penalty,omitempty"`
	FrequencyPenalty int            `json:"frequency_penalty,omitempty"`
	BestOf           int            `json:"best_of,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
}
