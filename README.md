# gakujo_Slack

成績を自動で教えてくれるSlackBotを作りました。

## 使い方
各自で.envを作成してください。
```console
$ echo -e 'J_USERNAME=学情のID\nJ_PASSWORD=学情のPSWD\nBOT_TOKEN=自分で作ったSlackAPIのBotToken(xoxbで始まるやつ)\nBOT_CHANNEL=送信したいチャンネル名' > ./.env
```

時間指定は、各自のリポジトリにコピーしてGithub Actionを使って指定してください。
このリポジトリの.github/workflowsを参考にしてください。
**注意**
時間設定はUTC時間になっているため、日本時間にしたい場合は、入力した時間＋9時間で考えてください

## 実行
```console
$ go run main.go
```
