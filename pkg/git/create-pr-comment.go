package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pull-request-formatter/pkg/config"
)

func CreatePRComment(versionsLog string) (err error) {
	payload, err := json.Marshal(struct {
		Body string `json:"body"`
	}{
		Body: versionsLog,
	})
	if err != nil {
		return
	}

	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d/comments", config.GitOwner, config.GitRepo, pr.Number)

	req, err := http.NewRequest("POST", apiUrl, bytes.NewReader(payload))
	if err != nil {
		return
	}

	_, err = send(req)

	return
}
