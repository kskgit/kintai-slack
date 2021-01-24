# SocketModeでのSlackAPI連携
開発中は外部に公開する事なくSlackに接続したいため、SocketModeを使用してSlackに接続する。

# socketモードとは
- 外部にAPIを公開する事なくSlackと通信可能
- socketモードについての説明
- https://api.slack.com/apis/connections/socket

# 参考にしたコード
- slack-goを用いたsocket接続のサンプルコード 
- https://github.com/slack-go/slack/blob/master/examples/socketmode/socketmode.go

# ソースコード
```go
//TODO コード貼付
```
# 接続までの流れ

## 1 slack.New
```go
api := slack.New(
	botToken,
	slack.OptionDebug(true),
	slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	slack.OptionAppLevelToken(appToken),
)
```
- 第一引数トークンを受け取る事でSlackクライアントを返す
- https://pkg.go.dev/github.com/slack-go/slack#New

## 2 socketmode.New
```go
client := socketmode.New(
	api,
	socketmode.OptionDebug(true),
	socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
)
```
- 第一引数にSlackクライアントを受け取る事でSocketモードクライアントを返す
- https://pkg.go.dev/github.com/slack-go/slack/socketmode?utm_source=gopls#New

## 3 socketmode.Run()
```go
client.Run()
```
- RunメソッドにてソケットモードでSlackと接続する
- https://pkg.go.dev/github.com/slack-go/slack/socketmode?utm_source=gopls#Client.Run




