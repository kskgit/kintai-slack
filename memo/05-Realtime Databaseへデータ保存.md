# やりたい事
勤務状態をRealtime Databaseで管理したい

# firestore VS Realtime Database
https://techblog.kayac.com/rtdb-vs-firestore
- 複雑なデータを扱わないし今後もその予定は無いため`Realtime Database`を使用する事に決定

# 参考
https://firebase.google.com/docs/firestore/quickstart?hl=ja#add_data

#インストールしたライブラリの使い方
- `go get`したら使えるようになると思い込んでて少し詰まった・・・

## そもそもgo getしたライブラリはどこに格納されてる？
- デフォルトは`$GOPATH/go/`
- `$GOPATH`のデフォルトは`$HOME/go`
  - https://golang.org/doc/articles/go_command.html#tmp_3
  >Both of these projects are now downloaded and installed into $HOME/go（Getting started with the go command）
- `go env GOPATH`で今のGOPATHを確認する事が出来る
  - https://golang.org/doc/gopath_code.html#GOPATH

## go getしたライブラリをどう呼び出すか
- `go.mod`ファイルを作成し、`go build`で依存関係にあるライブラリをインストールしてくれる
https://qiita.com/propella/items/e49bccc88f3cc2407745

### go.modファイルの作成
よく分からなかったがとりあえずgithubのURLを指定してみる
```
go mod init  github.com/kskgit/kintai-slack
```
- 対象フォルダで`go build`すると依存関係が解決される

# Realtime Databaseの使い方
## 使用ライブラリ
- とりあえず公式の通り`cloud.google.com/go/firestore v1.4.0`を使用
## 初期化
```go
		ctx := context.Background()
		conf := &firebase.Config{
			DatabaseURL: "https://databaseName.firebaseio.com",
		}
		// Fetch the service account key JSON file contents
		opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")

		app, err := firebase.NewApp(ctx, conf, opt)
		if err != nil {
			log.Fatalln("Error initializing app:", err)
		}

		rtb_client, err := app.Database(ctx)
		if err != nil {
			log.Fatalln("Error initializing database client:", err)
		}
```
- （参考）https://firebase.google.com/docs/database/admin/start?hl=ja

## 値取得
```go
ref := rtb_client.NewRef("time-log")
```
- `time-log`キーが無ければ新規作成されるしあればキー配下の値を取得する
