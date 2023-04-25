package main

import (
	"pull-request-formatter/pkg/config"
	"pull-request-formatter/pkg/git"
	"pull-request-formatter/pkg/log"
	"pull-request-formatter/pkg/openai"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Error(err)
		return
	}

	prompt, err := getPrompt()
	if err != nil {
		log.Error(err)
		return
	}

	changelog, err := openai.Send(prompt)
	if err != nil {
		log.Error(err)
		return
	}

	log.SaveToFile(changelog, "changelog")

	err = git.SetPRDescription(changelog)
	if err != nil {
		log.Error(err)
		return
	}

	log.Success()
}

func getPrompt() (prompt string, err error) {
	commits, err := git.GetCommits()
	if err != nil {
		return
	}

	prompt = config.PromptPreText

	for _, c := range commits {
		prompt += "\n - " + c.Commit.Message
	}

	prompt += "\n" + config.PromptAfterText

	log.SaveToFile(prompt, "prompt")

	return
}
