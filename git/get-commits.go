package git

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"pull-request-formatter/config"
)

var pr PullRequest

func GetCommits() (commits []CommitBody, err error) {
	Init()

	// get the pull request number for the branch.
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?base=%s", config.GitOwner, config.GitRepo, config.GitBranch)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	var pullRequests []PullRequest

	err = json.NewDecoder(res.Body).Decode(&pullRequests)
	if err != nil {
		return
	}

	if len(pullRequests) == 0 {
		return commits, errors.New(fmt.Sprintf("No pull requests found for branch %s\n", config.GitBranch))
	}

	pr = pullRequests[0]

	// get the commits associated with the pull request.
	req, err = http.NewRequest("GET", pr.Links.Commits.Href, nil)
	if err != nil {
		return
	}

	res, err = client.Do(req)
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
