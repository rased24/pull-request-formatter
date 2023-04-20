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

func GetCommits() (commits []CommitBody, err error) {
	// get the pull request number for the branch.
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?base=%s", config.GitOwner, config.GitRepo, config.GitBranch)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.GitAccessToken))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	var pullRequests []PullRequests

	err = json.NewDecoder(res.Body).Decode(&pullRequests)
	if err != nil {
		return
	}

	if len(pullRequests) == 0 {
		return commits, errors.New(fmt.Sprintf("No pull requests found for branch %s\n", config.GitBranch))
	}

	// get the commits associated with the pull request.
	req, err = http.NewRequest("GET", pullRequests[0].Links.Commits.Href, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.GitAccessToken))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, err = http.DefaultClient.Do(req)
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
