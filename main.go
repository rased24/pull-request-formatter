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

	commits, err := git.GetCommits()
	if err != nil {
		log.Error(err)
		return
	}

	prompt := config.PromptPreText

	for _, commit := range commits {
		prompt += "\n" + commit.Commit.Message
	}

	prompt += "\n" + config.PromptAfterText

	changelog, err := openai.Send(prompt)
	if err != nil {
		log.Error(err)
		return
	}

	err = git.SetPRDescription(changelog)
	if err != nil {
		log.Error(err)
		return
	}

	log.Success()
}
