# これはなに？
jsonでリクエストを記述すると、APIの実行 & 表示 & ドキュメント作成をしてくれるコマンドラインツール

# 準備
インストール
```bash
go get github.com/aikizoku/rundoc
```

初期化
```bash
rundoc init
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
リクエストのリストを表示する
```bash
rundoc list
```

リクエストを実行する
```bash
rundoc run sample
```

環境を指定して実行する
```bash
rundoc run sample -e staging
```

実行してドキュメントを作成する
```bash
rundoc run sample -d
```
