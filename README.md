# これはなに？
みんな大好きjsonでリクエストを記述しておけば、実行&結果表示&ドキュメント作成をしてくれるコマンドラインツール

# 準備
configフォルダに共通設定のjsonを置く
```
config
└common.json
└auth.json
```

runsフォルダにリクエストを送りたいjsonを置く
```
runs
└get_me.json
└post_user_profile.json
```

docsフォルダを作るする
```
docs
```

# 各ファイルのサンプル
common.json
```json
{  
    "endpoints": {
        "local": "http://localhost:8080",
        "staging": "https://staging.appspot.com",
        "production": "https://appspot.com"
    },
    "headers": {
        "Content-Type": "application/json"
    }
}
```

auth.json
```json
{
    "local": "sample_auth_token",
    "staging": "sample_auth_token",
    "production": "sample_auth_token"
}
```

request.json
```json
{
    "description": "ほげほげをするAPI",
    "path": "/hogehoge/fuga",
    "method": "post",
    "headers": {
        "X-OS": "iOS",
    },
    "params": {
        "hoge": "aaaaa",
        "fuga": "xxxxx"
    }
}
```

# 使い方
リクエスト名のリストを表示する
```bash
./rundoc -l
```

任意のリクエストをローカル環境で実行する
```bash
./rundoc -n get_me
```

任意のリクエストを環境を指定して実行する
```bash
./rundoc -n get_me -e staging
```

任意のリクエストをローカル環境で実行してドキュメントを作成する
```bash
./rundoc -n get_me -d
```
