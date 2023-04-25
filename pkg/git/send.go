package git

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func send(req *http.Request) (res *http.Response, err error) {
	res, err = client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode == http.StatusOK {
		return
	}

	var response errorResponse

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return
	}

	return res, errors.New(fmt.Sprintf("[GIT] %s", response.Message))
}

type errorResponse struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}