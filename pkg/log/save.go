package log

import (
	"fmt"
	"log"
	"os"
	"pull-request-formatter/pkg/config"
)

// SaveToFile - gives the ability to save user generated messages to the log file
func SaveToFile(message string) {
	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	log.SetOutput(file)

	log.Println(fmt.Sprintf("[%d] %s", config.SessionId, message))

}

func Error(err error) {
	SaveToFile(err.Error())

	//----------------INTEGRATIONS-------------------//(NOTE: The integrations will only work if the required params configured in the .env file)

	telegramError() //send error notification to telegram
}

func Success() {
	//----------------INTEGRATIONS-------------------//(NOTE: The integrations will only work if the required params configured in the .env file)

	telegramSuccess() //send success notification to telegram
}
