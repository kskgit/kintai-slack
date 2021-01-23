# yamlファイルの読込方法
- src/secrets.yaml
```yml
env_variables:
  SLACK_SIGNING_SECRET: XXX
  SLACK_BOT_TOKEN: XXX
```
- src/main.go
```go
import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// yamlファイルを読込
secrets, err := ioutil.ReadFile("secrets.yaml")
if err != nil {
  return
}

// yamlファイルから値を取り出し
m := make(map[interface{}]map[interface{}]interface{})
err = yaml.Unmarshal(secrets, &m)
if err != nil {
  panic(err)
}
fmt.Println(m["env_variables"])
fmt.Println(m["env_variables"]["SLACK_SIGNING_SECRET"])
```
### gopkg.in/yaml.v2
- https://godoc.org/gopkg.in/yaml.v2
- yaml.Unmarshal
  - 第一引数をデコードし、第二引数に値を割り当てる
    - >Unmarshal decodes the first document found within the in byte slice and assigns decoded values into the out value
- 使用例
  - https://ota42y.com/blog/2014/11/13/go-yaml/

### ネストしたMapの使い方
```go
// 事前にネストしたMap型を定義する
m := make(map[interface{}]map[interface{}]interface{})

// 以下の形式で取り出し可能
fmt.Println(m["env_variables"]["SLACK_SIGNING_SECRET"])
```
- 参考
https://stackoverflow.com/questions/44305617/nested-maps-in-golang
