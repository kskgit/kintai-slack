# やりたい事
勤務状態をfirestoreで管理したい


# 参考
https://firebase.google.com/docs/firestore/quickstart?hl=ja#add_data



## go getしたライブラリはどこに格納されてる？
- `$GOPATH/go/`
  - https://golang.org/doc/articles/go_command.html#tmp_3
  >Both of these projects are now downloaded and installed into $HOME/go（Getting started with the go command）

### $GOPATHとは
- Unixのデフォルトは`$HOME/go`（変更する事も可能）
- `go env GOPATH`で今のGOPATHを確認する事が出来る
https://golang.org/doc/gopath_code.html#GOPATH

## go getしたライブラリをどう呼び出すか
### go modulesについて
https://qiita.com/propella/items/e49bccc88f3cc2407745
### モジュール名とは
よく分からなかったがとりあえずgithubのURLを指定してみる
```
go mod init  github.com/kskgit/kintai-slack
```
- go.modファイルが作成された

```
go build
```
- 関連するmoduleが取得された
