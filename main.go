package main

import (
	"log"
	"pull-request-formatter/config"
	"pull-request-formatter/git"
	"pull-request-formatter/gpt"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Println(err)
		return
	}

	commits, err := git.GetCommits()

	promptBody := "Write changelog based on these commit titles:"

	for _, commit := range commits {
		promptBody += "\n" + commit.Commit.Message
	}

	promptBody += "\nYou can take this changelog as an example: \nChange status link now displays static expiration time and time remaining until the appointment starts;\nFixed issue with limited booking days and timeslots not being displayed at the end of the month in the calendar service;\nFixed right/left alignment buttons in RTL mode in the calendar;\nFixed multiple RTL issues;\nFixed bug where staff names appeared twice in the day view of the calendar;\nAdded company name link in customer panel that redirects to the company address;\nAdded English (US) to dynamic translations;\nFixed issue with creating new custom status when there is a waiting list;\nFixed table width problem in invoices;\nImproved customer panel to prevent conflicts when multiple panels are open at once;\nFixed timezone problem in recurring appointments;\nFixed timeslot overlapping issue."

	changelog, err := gpt.Send(promptBody)
	if err != nil {
		log.Println(err)
		return
	}

	err = git.SetPRDescription(changelog)
	if err != nil {
		log.Println(err)
	}
}
