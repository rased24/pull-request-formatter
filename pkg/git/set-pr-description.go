package git

import (
	"bytes"
	"encoding/json"
	"net/http"
	"pull-request-formatter/pkg/config"
)

func SetPRDescription(changelog string) (err error) {
	payload, err := json.Marshal(struct {
		Title string `json:"title"`
		Body  string `json:"body"`
		State string `json:"state"`
		Base  string `json:"base"`
	}{
		Title: pr.Title,
		Body:  changelog,
		State: "open",
		Base:  config.GitBranch,
	})
	if err != nil {
		return
	}

	req, err := http.NewRequest("PATCH", pr.URL, bytes.NewReader(payload))
	if err != nil {
		return
	}

	_, err = send(req)

	return
}
