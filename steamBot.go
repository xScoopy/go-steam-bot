package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func getEnvVariable(key string) string {
	//method to retrieve slackbot token from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Problem loading .env file.")
	}
	return os.Getenv(key)
}

func main() {
	//initialize slack api connection 
	api := slack.New(getEnvVariable("GOSECRET"))

	channelID, timestamp, err := api.PostMessage(
		"C034MB8U5M5",
		slack.MsgOptionText("Hello World", false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("Message sent successfully to channel %s at %s", channelID, timestamp)

}