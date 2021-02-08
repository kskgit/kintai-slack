# やりたい事
勤務状態をRealtime Databaseで管理したい

# firestore VS Realtime Database
https://techblog.kayac.com/rtdb-vs-firestore
- 複雑なデータを扱わないし今後もその予定は無いため`Realtime Database`を使用する事に決定

# 参考
https://firebase.google.com/docs/firestore/quickstart?hl=ja#add_data
https://qiita.com/yukpiz/items/9d81da697c9c9faab83d#%E3%83%87%E3%83%BC%E3%82%BF%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8Bget

# インストールしたライブラリの使い方
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
- とりあえず公式の通り`firebase.google.com/go`を使用
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
// start配下の値を取得する
ex_ref := rtb_client.NewRef("time-log/" + cmd.UserID + "/" + date + "/start")

// 受取る値を定義
// 型はstart配下の値の型を定義
var startTimeLog string

// 値を取得
if err := ex_ref.Get(ctx, &startTimeLog); err != nil {
	fmt.Println("===error")
	log.Fatalln("Error reading value:", err)
}

// 値の有無で条件分岐
// stringの空文字判定はlengthで行う
if len(startTimeLog) > 0 {
	// 開始時間を更新するか確認する処理
	fmt.Println("===開始時間を更新するか確認する")
}
```

