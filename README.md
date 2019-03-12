# rundoc


rundoc

みんな大好きjsonでリクエストを記述しておけば、実行&結果表示&Doc作成.mdをしてくれる

rundoc name -e local staging production -d

rundoc -l


common.json
{
    "endpoints": [
        {
            "env": "local",
            "url": "http://localhost:8080"
        },
        {
            "env": "staging",
            "url": "https://push.staging.salontia.rabee.jp"
        },
        {
            "env": "production",
            "url": "https://push.salontia.rabee.jp"
        }
    ],
    "headers": [
        {
            "field": "Content-Type",
            "value": "application/json"
        },
    ]
}

authorization.example.json
authorization.json
{
    "local": "token",
    "staging": "token",
    "production": ""
}


entry.json
{
    "title": "post_api_rpc_entry",
    "description": "",
    "path": "/api/rpc",
    "method": "post",
    "headers": {

    },
    "params": {

    }
}

+ rundoc            
  - rundoc
  + config
    - common.json
    - authorization.example.json
    - authorization.json
    - doc.tmpl
  + runs
    - get_entry.json
  + docs
    - get_entry.mk