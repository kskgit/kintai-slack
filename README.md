# kintai-slack
# 仕様
### 「開始」と入力
- 勤務時間の打刻開始
  - 「開始」と入力すると「開始」を受け付けなくなる
### 「休憩」と入力
- 勤務時間の打刻ストップ
- 休憩時間の打刻開始
  - 「休憩」と入力すると「再開」しか受け付けなくなる
### 「再開」と入力
- 勤務時間の打刻再開
- 休憩時間の打刻ストップ
  - 「休憩」と入力してない状態で「再開」と入力するとエラー
### 「終了」と入力
- 勤務時間の打刻ストップ
- 一日の勤務時間・休憩時間が一覧で表示される
  - 「開始」と入力してない状態で「終了」と入力するとエラー
