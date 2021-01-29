package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"google.golang.org/api/option"
)

var env_values = make(map[interface{}]map[interface{}]interface{})

func main() {
	err := init_env(&env_values)
	if err != nil {
		panic(err)
	}

	appToken := env_values["slack_keys"]["SLACK_APP_TOKEN"].(string)
	// todo エラー処理
	botToken := env_values["slack_keys"]["SLACK_BOT_TOKEN"].(string)
	// todo エラー処理

	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	go func() {
		ctx := context.Background()
		sa := option.WithCredentialsFile("src/kintai-slack-firebase-adminsdk-1pbri-90fe41bf4c.json")
		app, err := firebase.NewApp(ctx, nil, sa)
		if err != nil {
			log.Fatalln(err)
		}
		firebase_client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		defer firebase_client.Close()
		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeConnecting:
				fmt.Println("Connecting to Slack with Socket Mode...")
			case socketmode.EventTypeConnected:
				fmt.Println("Connected to Slack with Socket Mode.")
			case socketmode.EventTypeSlashCommand:
				cmd, ok := evt.Data.(slack.SlashCommand)
				if !ok {
					fmt.Printf("Ignored %+v\n", evt)

					continue
				}
				fmt.Printf("cmd received==== %+v\n", cmd)
				// 開始
				// firestoreに開始時間を保存
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
			}
		}
	}()

	client.Run()

}
