package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pull-request-formatter/pkg/config"
)

func Send(prompt string) (result string, err error) {
	buff, err := json.Marshal(Request{
		Model:       config.OpenAIModel,
		Prompt:      prompt,
		MaxTokens:   config.OpenAIMaxTokens,
		Temperature: config.OpenAITemperature,
	})
	if err != nil {
		return
	}

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

	return handleResponse(res)
}

func handleResponse(res *http.Response) (result string, err error) {
	if res.StatusCode == http.StatusOK {
		var successBody successResponse

		err = json.NewDecoder(res.Body).Decode(&successBody)
		if err != nil {
			return
		}

		result = successBody.Choices[0].Text

		return
	}

	var errorBody errorResponse

	err = json.NewDecoder(res.Body).Decode(&errorBody)
	if err != nil {
		return
	}

	err = errors.New(fmt.Sprintf("[OPENAI] %s (Error Type: %s)", errorBody.Error.Message, errorBody.Error.Type))

	return
}
