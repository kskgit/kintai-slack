# やりたい事
「開始」、「休憩」、「終了」のいずれかの文字が投稿された場合イベントを受け取りたい
https://qiita.com/kanaxx/items/c29267d88c3fb2cc381c
- 全てのイベントタイプ
  - https://api.slack.com/events

# Slackの設定
## 結論
- Slash Commandsでイベントを受け取る
- `/勤怠 開始`の様な形式

## 試したけど採用しなかったもの
- `message.im`イベントを用いて文字列を受け取って判定
→日本語を直接受け取れず（utf8形式でしか受け取れず）、分かりにくいと判断したため不採用
- https://api.slack.com/events/message.im

# 
