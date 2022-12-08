package openai

type ResponseBody struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage
	/*
		{
		  "id": "cmpl-uqkvlQyYK7bGYrRHQ0eXlWi7",
		  "object": "text_completion",
		  "created": 1589478378,
		  "model": "text-davinci-003",
		  "choices": [
		    {
		      "text": "\n\nThis is indeed a test",
		      "index": 0,
		      "logprobs": null,
		      "finish_reason": "length"
		    }
		  ],
		  "usage": {
		    "prompt_tokens": 5,
		    "completion_tokens": 7,
		    "total_tokens": 12
		  }
		}

	*/
}

type Choice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	LogProbs     int    `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
