package config

import (
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var GitOwner string
var GitRepo string
var GitBranch string
var GitAccessToken string

var PromptPreText string
var PromptAfterText string

var OpenAISecret string
var OpenAIModel string
var OpenAIMaxTokens int
var OpenAITemperature float64

var TgAccessToken string
var TgChatId string

var SessionId int

func Init() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}

	GitOwner = os.Getenv("OWNER")
	GitRepo = os.Getenv("REPO")
	GitBranch = os.Getenv("BRANCH")
	GitAccessToken = os.Getenv("GIT_ACCESS_TOKEN")

	PromptPreText = os.Getenv("PROMPT_PRETEXT")
	PromptAfterText = os.Getenv("PROMPT_POSTTEXT")

	TgAccessToken = os.Getenv("TELEGRAM_ACCESS_TOKEN")
	TgChatId = os.Getenv("TELEGRAM_CHAT_ID")

	OpenAISecret = os.Getenv("OPEN_AI_SECRET")
	OpenAIModel = os.Getenv("OPEN_AI_MODEL")

	rand.Seed(time.Now().UnixNano())

	SessionId = rand.Intn(1000)

	OpenAIMaxTokens, err = strconv.Atoi(os.Getenv("OPEN_AI_MAX_TOKENS"))
	if err != nil {
		return
	}

	OpenAITemperature, err = strconv.ParseFloat(os.Getenv("OPEN_AI_TEMPERATURE"), 64)

	return
}
