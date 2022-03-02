package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

//struct to hold game data
type GameInfo struct {
	Name        string
	Price       string
	ReleaseDate string
}

func getEnvVariable(key string) string {
	//method to retrieve slackbot token from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Problem loading .env file.")
	}
	return os.Getenv(key)
}

func scrapeSteam() []GameInfo {
	//setup colly scraper
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	//setup games slice
	games := make([]GameInfo, 0)

	c.OnHTML("a.search_result_row", func(e *colly.HTMLElement) {
		e.ForEach("div.responsive_search_name_combined", func(i int, h *colly.HTMLElement) {
			if h.ChildText("div.discounted") != "" {
				newGame := GameInfo{}
				newGame.Name = h.ChildText("span.title")
				newGame.ReleaseDate = h.ChildText("div.search_released")
				newGame.Price = h.ChildText("div.discounted")
				games = append(games, newGame)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Received error:", e)
	})

	c.Visit("https://store.steampowered.com/search/?filter=topsellers")

	//create json for data persistence points on rubric
	createJson(games, "output.json")

	return games
}

func createJson(games []GameInfo, fileName string) error {
	jsonFile, _ := json.MarshalIndent(games, "", " ")
	err := ioutil.WriteFile(fileName, jsonFile, 0644)
	if err != nil {
		return err
	}
	return nil
}

func formatGames(games []GameInfo) string {
	formattedGames := ""
	for _, game := range games {
		formattedGames += "\nTitle: *" + game.Name + "* \nPrice:(original vs discounted) *" + game.Price + "* \nRelease: *" + game.ReleaseDate + "*\n"
	}
	return formattedGames
}

//func to handle events from the socket properly based on event type.
func handleEventMessage(event slackevents.EventsAPIEvent, api *slack.Client, chanID string) error {
	switch event.Type {
	//if its a callback event
	case slackevents.CallbackEvent:
		innerEvent := event.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			log.Print(ev)
			//scrape steam best sellers and return slice of games
			gameList := scrapeSteam()
			//format them to be added to a slack message
			formattedGames := formatGames(gameList)
			postSlackMessage(formattedGames, chanID, api)
		}
	default:
		return errors.New("We don't support this type of event yet")

	}
	return nil
}
func postSlackMessage(message string, chanID string, api *slack.Client) {
	//create slack attachment
	attachment := slack.Attachment{
		Pretext: "Freshly Scraped Steam Top Sellers",
		Text:    "Served up at your command",
		Color:   "#36a64f",
		Fields: []slack.AttachmentField{
			{
				Title: "Game",
				Value: message,
			},
		},
	}
	//send hello world to channel id
	_, timestamp, err := api.PostMessage(
		chanID,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("Message sent successfully to channel %s at %s", chanID, timestamp)
}

func main() {
	channelId := getEnvVariable("CHANNELID")
	apptoken := getEnvVariable("STEAMBOT-WEBSOCKET")
	bottoken := getEnvVariable("GOSECRET")

	//setup slack websocket for listening to events
	client := slack.New(bottoken, slack.OptionDebug(true), slack.OptionAppLevelToken(apptoken))
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		//logger
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	//context to cancel a goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//goroutine to handle events from the socket client
	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
		//for/while loop for context cancellation or incoming events.
		for {
			select {
			// if context cancel, exit goroutine
			case <-ctx.Done():
				log.Println("Shutting down socketmode listener")
				return
			case event := <-socketClient.Events:
				//New events are in here
				switch event.Type {
				//handle eventapi events
				case socketmode.EventTypeEventsAPI:
					eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)
					if !ok {
						log.Printf("Could not type cast event to the event: %v\n", event)
						continue
					}
					//send ack to slack server
					socketClient.Ack(*event.Request)
					//custom func to handle event types to avoid too much nested switching
					err := handleEventMessage(eventsAPIEvent, client, channelId)
					if err != nil {
						log.Panic("Unable to handle this type of event")
					}
				}
			}
		}
	}(ctx, client, socketClient)

	//run socket to block program from ending and listen to any incoming events
	socketClient.Run()
}
