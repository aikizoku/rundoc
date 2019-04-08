# これはなに？
みんな大好きjsonでリクエストを記述すると、APIの実行 & 表示 & ドキュメント作成をしてくれるコマンドラインツール

# 準備
初期化する
```bash
./rundoc -i
```
上記コマンドを実行すると下記ディレクトリが作成される
```
config
  └ common.json
  └ auth.json
runs
  └ sample.json
docs
```
common.json と auth.json に共通情報を記載する
runs/sample.jsonを参考にリクエストを記載したjsonを作る

# 使い方
リクエスト名のリストを表示する
```bash
./rundoc -l
```

リクエストを実行する
```bash
./rundoc -r sample
```

リクエストを環境を指定して実行する
```bash
./rundoc -e staging -r sample
```

リクエストを実行してドキュメントを作成する
```bash
./rundoc -d -r sample
```
