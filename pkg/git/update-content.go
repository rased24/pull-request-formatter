package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pull-request-formatter/pkg/config"
)

func updateContent(commit updatedFileBody, path string) (response fileResponse) {
	payload, err := json.Marshal(commit)
	if err != nil {
		return
	}

	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", config.GitOwner, config.GitRepo, path)
	req, err := http.NewRequest("PUT", apiUrl, bytes.NewReader(payload))
	if err != nil {
		return
	}

	res, err := send(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return
	}

	return

}
