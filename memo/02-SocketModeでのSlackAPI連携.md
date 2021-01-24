# SocketModeでのSlackAPI連携
開発中は外部に公開する事なくSlackに接続したいため、SocketModeを使用してSlackに接続する。

### socketモードとは
- 外部にAPIを公開する事なくSlackと通信可能
- 【公式】socketモードについての説明
https://api.slack.com/apis/connections/socket

### 参考にしたコード
- slack-goを用いたsocket接続のサンプルコード 
https://github.com/slack-go/slack/blob/master/examples/socketmode/socketmode.go

### ソースコード

###### slack.New
```go
api := slack.New(
	botToken,
	slack.OptionDebug(true),
	slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	slack.OptionAppLevelToken(appToken),
)
```
- 第一引数トークンを受け取る事でSlackクライアントを返す
- 【公式】https://pkg.go.dev/github.com/slack-go/slack#New

###### socketmode.New
```go
	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)
```
- 第一引数にSlackクライアントを受け取る事でSocketモードクライアントを返す
- 【公式】https://pkg.go.dev/github.com/slack-go/slack/socketmode?utm_source=gopls#New





###### interfaceの型変換
```go
api := slack.New(m["env_slack_keys"]["SLACK_SIGNING_SECRET"].(string))
```
- `m["env_slack_keys"]["SLACK_SIGNING_SECRET"]`はinterface型のため、`(string)`で型を変換する必要がある

https://qiita.com/lostfind/items/ad7bfc1a4860bb108b9c#interface%E3%81%AE%E5%A4%89%E6%8F%9B

### os.Getenv


