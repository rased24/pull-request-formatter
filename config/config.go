package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var GitOwner string
var GitRepo string
var GitBranch string
var GitAccessToken string

var OpenAISecret string
var OpenAIModel string
var OpenAIMaxTokens int

func Init() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}

	GitOwner = os.Getenv("OWNER")
	GitRepo = os.Getenv("REPO")
	GitBranch = os.Getenv("BRANCH")
	GitAccessToken = os.Getenv("ACCESS_TOKEN")

	OpenAISecret = os.Getenv("OPEN_AI_SECRET")
	OpenAIModel = os.Getenv("OPEN_AI_MODEL")

	OpenAIMaxTokens, err = strconv.Atoi(os.Getenv("OPEN_AI_MAX_TOKENS"))

	return
}
