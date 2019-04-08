# post_api_rpc_entry

プッシュ通知のエントリーを行う

## Request

|ENV|URL|
|---|---|
|Local|`POST` http://localhost:8080/api/rpc|
|Staging|`POST` https://staging.appspot.com/api/rpc|
|Production|`POST` https://appspot.com/api/rpc|

```
Authorization: xxxxxxxxxx
Content-Type: application/json
```
```json
{
    "id": "0",
    "jsonrpc": "2.0",
    "method": "entry",
    "params": {
        "app_id": "sample_app",
        "device_id": "sample_device_id",
        "platform": "ios",
        "token": "sample_token",
        "user_id": "sample_user_id"
    }
}
```

## Response

```
Status 200
```
```json
{
    "jsonrpc": "2.0",
    "id": "0",
    "result": {
        "success": true
    }
}
```

