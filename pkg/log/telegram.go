package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pull-request-formatter/pkg/config"
)

type tgResponse struct {
	Ok     bool `json:"ok"`
	Result *struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
	ErrorCode   *int    `json:"error_code"`
	Description *string `json:"description"`
}

func telegramSend(message string) (err error) {
	if config.TgAccessToken == "" || config.TgChatId == "" {
		return
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?text=%s&chat_id=%s", config.TgAccessToken, message, config.TgChatId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	var response tgResponse

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("[TELEGRAM] %s (Error Code: %d )", *response.Description, response.ErrorCode))
	}

	return
}

func telegramError() {
	err := telegramSend("Changelog generation failed. Please, check the log file for details")
	if err != nil {
		SaveToFile(err.Error())
	}
}

func telegramSuccess() {
	err := telegramSend("Changelog generation succeeded")
	if err != nil {
		SaveToFile(err.Error())
	}
}
