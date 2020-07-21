# これはなに？
jsonでリクエストを記述すると、APIの実行 & 表示 & ドキュメント作成をしてくれるコマンドラインツール

# 準備

## インストール

```bash
go get github.com/aikizoku/rundoc
```

## 初期化

```bash
rundoc init
```

## 上記コマンドを実行すると下記ディレクトリが作成される

```
rundoc 
  └ config
    └ common.json
    └ auth.json
  └ runs
    └ sample.json
  └ docs
```
common.json と auth.json に共通情報を記載する
runs/sample.jsonを参考にリクエストを記載したjsonを作る

# 使い方
rundocフォルダがあるディレクトリで下記コマンドを実行する

## リクエストのリストを表示する

```bash
rundoc list
```

## リクエストを実行する

```bash
rundoc
```

実行したいAPIを選ぶ
```bash
sample_delete
sample_get
sample_post
sample_put
```

実行する環境を選択する
```bash
production
staging
local
```

ドキュメントを作成するか？を選択する
```
false
true
```

## 実行結果

```
------------- Request
DELETE /sample
Authorization: sample_local_token
Content-Type: application/json
X-OS: iOS
{
    "fuga": "aaaaa",
    "hoge": "xxxxx"
}

------------- Response
time: 7ms
status: 200
{
    "fuga": 1,
    "hoge": "aaaaaa"
}

------------- Command
rundoc sample_delete -e staging -d
```