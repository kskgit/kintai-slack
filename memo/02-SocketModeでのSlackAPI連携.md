# SocketModeでのSlackAPI連携
```go
import (
	"io/ioutil"

	"github.com/slack-go/slack"
	"gopkg.in/yaml.v2"
)
```

```go


```

### socketモードとは
- 【公式】socketモードについての説明
https://api.slack.com/apis/connections/socket

- slack-goを用いたsocket接続のサンプルコード 
https://github.com/slack-go/slack/blob/master/examples/socketmode/socketmode.go

### interfaceの型変換
```go
api := slack.New(m["env_slack_keys"]["SLACK_SIGNING_SECRET"].(string))
```
- `m["env_slack_keys"]["SLACK_SIGNING_SECRET"]`はinterface型のため、`(string)`で型を変換する必要がある

https://qiita.com/lostfind/items/ad7bfc1a4860bb108b9c#interface%E3%81%AE%E5%A4%89%E6%8F%9B

### os.Getenv


