package main

import (
	"fmt"
	"pull-request-formatter/pkg/config"
	"pull-request-formatter/pkg/git"
	"pull-request-formatter/pkg/log"
	"pull-request-formatter/pkg/openai"
	"regexp"
	"strings"
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
		message := c.Commit.Message

		// check if message starts with "Merge branch" or "Merge pull request"
		if strings.HasPrefix(message, "Merge branch") || strings.HasPrefix(message, "Merge pull request") {
			continue
		}

		//check if the commit is just a followup commit of the previous ones
		if strings.HasPrefix(message, "@ignore") {
			continue
		}

		// remove links from the message
		re := regexp.MustCompile(`\bhttps?://\S+`)
		message = re.ReplaceAllString(message, "")

		// remove all newlines from the message
		message = strings.ReplaceAll(message, "\n", "")

		prompt += "\n - " + message
	}

	prompt += "\n" + config.PromptAfterText

	log.SaveToFile(prompt, "prompt")

	return
}

func getVersionsText() (versionsLogText string, err error) {
	versions, err := git.GetVersions()
	if err != nil {
		return
	}

	versionsLogText = config.VersionsLogPreText

	for _, versionObj := range versions {
		versionsLogText += fmt.Sprintf("| %s | %s | %d |\n", versionObj.Name, versionObj.OldVersion, versionObj.OldIntVersion)
	}

	versionsLogText += config.VersionsLogPostText

	for _, versionObj := range versions {
		versionsLogText += fmt.Sprintf("| %s | %s | %d |\n", versionObj.Name, versionObj.NewVersion, versionObj.NewIntVersion)
	}

	return
}
