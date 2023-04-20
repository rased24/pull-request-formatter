package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pull-request-formatter/config"
)

func Send(prompt string) (result string, err error) {
	buff, err := json.Marshal(Request{
		Model:       config.OpenAIModel,
		Prompt:      prompt,
		MaxTokens:   config.OpenAIMaxTokens,
		Temperature: config.OpenAITemperature,
	})

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(buff))
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.OpenAISecret))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	var response Response

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return
	}

	result = response.Choices[0].Text

	return
}
