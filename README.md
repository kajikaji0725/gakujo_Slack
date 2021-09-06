# gakujo_Slack

成績を自動で教えてくれるSlackBotを作りました。  
herokuを使おうかと思いましたが、できなかったのでAWSを使いました。

## 使い方
各自で.envを作成してください。
```console
$ echo -e 'J_USERNAME=学情のID\nJ_PASSWORD=学情のPSWD\nBOT_TOKEN=自分で作ったSlackAPIのBotToken(xoxbで始まるやつ)\nBOT_CHANNEL=送信したいチャンネル名' > ./.env
```
以下のコマンドを用いて実行ファイルを作成してください  
```console
$ go build -o ./app
```  
## 実行
```console
$ nohup ./app &
```  
## 削除
もし、何かプログラムに変更点があった場合、適宜削除してからもう一度buildして実行させる必要があります。  
```console
$ ps aux | grep app  
$ kill number  
```  
最初のコマンドでappのプロセス番号を確認します。  
左から2番目の数字がプロセス番号になります。  
killでは、numberのところに確認したプロセス番号を打ち込みます。  
  
***これらのコマンドはWindowsでは使えません***


