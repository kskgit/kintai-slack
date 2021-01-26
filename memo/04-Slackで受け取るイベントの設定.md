# やりたい事
「開始」、「休憩」、「終了」のいずれかの文字が投稿された場合イベントを受け取りたい

# Slackの設定
## 結論
- Slash Commandsでイベントを受け取る
- `/勤怠 開始`の様な形式
- 以下の形式でイベントを受取る事が可能※一部抜粋
```
{Token:XXX・・・ Command:/勤怠 Text:開始 ・・・}
```

## 試したけど採用しなかったもの
- `message.im`イベントを用いて文字列を受け取って判定
→日本語を直接受け取れず（utf8形式でしか受け取れず）、分かりにくいと判断したため不採用
- https://api.slack.com/events/message.im

# コード抜粋
```go
	go func() {
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
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
			}
		}
	}()
```
## イベントの受取
- `client.Events`をrangeでループ
- `socketmode.EventTypeSlashCommand`でスラッシュコマンドのイベントを受取

## 参考
https://github.com/slack-go/slack/blob/master/examples/socketmode/socketmode.go
https://pkg.go.dev/github.com/slack-go/slack/socketmode?utm_source=gopls#EventType
https://api.slack.com/events
