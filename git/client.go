package git

import (
	"net/http"
	"pull-request-formatter/config"
)

var client *http.Client

func Init() {
	client = &http.Client{
		Transport: &BearerTokenTransport{
			token: config.GitAccessToken,
			base:  &http.Transport{},
		},
	}
}

type BearerTokenTransport struct {
	token string
	base  http.RoundTripper
}

func (t *BearerTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	return t.base.RoundTrip(req)
}
