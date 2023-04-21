package git

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"pull-request-formatter/pkg/config"
)

var pr pullRequest

func GetCommits() (commits []commitBody, err error) {
	Init()

	err = getPr()
	if err != nil {
		return
	}

	// get the commits associated with the pull request.
	req, err := http.NewRequest("GET", pr.Links.Commits.Href, nil)
	if err != nil {
		return
	}

	res, err := send(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	err = json.NewDecoder(res.Body).Decode(&commits)

	return
}

func getPr() (err error) {
	// get the pull request number for the branch.
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?base=%s", config.GitOwner, config.GitRepo, config.GitBranch)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}

	res, err := send(req)
	if err != nil {
		return
	}

	var pullRequests []pullRequest

	err = json.NewDecoder(res.Body).Decode(&pullRequests)
	if err != nil {
		return
	}

	if len(pullRequests) == 0 {
		return errors.New(fmt.Sprintf("[GIT] No pull requests found for branch %s\n", config.GitBranch))
	}

	pr = pullRequests[0]

	return
}
