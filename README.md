# これはなに？
みんな大好きjsonでリクエストを記述しておけば、実行&結果表示&ドキュメント作成をしてくれるコマンドラインツール

# 準備
初期化
```
./rundoc -i

上記コマンドを実行すると下記ディレクトリが作成される
config
└common.json
└auth.json
runs
└sample.json
docs

runs/sample.jsonを参考にリクエストを作る
```

# 使い方
リクエスト名のリストを表示する
```bash
./rundoc -l
```

任意のリクエストをローカル環境で実行する
```bash
./rundoc -n sample
```

任意のリクエストを環境を指定して実行する
```bash
./rundoc -n sample -e staging
```

任意のリクエストをローカル環境で実行してドキュメントを作成する
```bash
./rundoc -n sample -d
```
