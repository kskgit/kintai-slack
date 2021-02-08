package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"google.golang.org/api/option"
)

var env_values = make(map[interface{}]map[interface{}]interface{})

type TimeLog map[string]map[string]struct {
	Start string `json:"start"`
}

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
		//
		// Realtime Database
		//
		ctx := context.Background()
		conf := &firebase.Config{
			DatabaseURL: "https://kintai-slack-default-rtdb.firebaseio.com/",
		}
		opt := option.WithCredentialsFile("./kintai-slack-firebase-adminsdk-1pbri-90fe41bf4c.json")
		app, err := firebase.NewApp(ctx, conf, opt)
		if err != nil {
			log.Fatalln("Error initializing app:", err)
		}
		rtb_client, err := app.Database(ctx)
		if err != nil {
			log.Fatalln("Error initializing database client:", err)
		}

		// イベント受け取り
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
				// 開始
				if cmd.Text == "開始" {
					base_time := time.Now()
					const date_layout = "2006-01-02"
					date := base_time.Format(date_layout)
					const time_layout = "15:04:05"
					time := base_time.Format(time_layout)
					// 事前に同日の開始時間有無を確認
					// 既に開始時間がある場合は開始時間を更新するか確認
					ex_ref := rtb_client.NewRef("time-log/" + cmd.UserID + "/" + date + "/start")
					var startTimeLog string
					if err := ex_ref.Get(ctx, &startTimeLog); err != nil {
						fmt.Println("===error")
						log.Fatalln("Error reading value:", err)
					}
					if len(startTimeLog) > 0 {
						// 開始時間を更新するか確認する処理
						fmt.Println("===開始時間を更新するか確認する")
					}
					// 開始時間更新
					ref := rtb_client.NewRef("time-log")
					err := ref.Set(ctx, TimeLog{
						cmd.UserID: {
							date: {
								Start: time,
							},
						},
					})
					if err != nil {
						log.Fatalln("Error setting value:", err)
					}
				}

			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
			}
		}
	}()

	// サーバー起動
	client.Run()

}
