# やりたい事
「開始」、「休憩」、「終了」のいずれかの文字が投稿された場合イベントを受け取りたい

https://qiita.com/kanaxx/items/c29267d88c3fb2cc381c

- 全てのイベントタイプ
  - https://api.slack.com/events

# Slackの設定
- 今回はボットにダイレクトメッセージを送った際にテキストを受け取るため、`message.im`イベントを用いる
- https://api.slack.com/events/message.im

- 日本語を直接受け取れないのか
- https://qiita.com/uchiko/items/1810ddacd23fd4d3c934
