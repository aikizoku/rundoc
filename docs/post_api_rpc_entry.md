# post_api_rpc_entry

プッシュ通知のエントリーを行う

|ENV|URL|
|---|---|
|Local|http://localhost:8080|
|Staging|https://staging.appspot.com|
|Production|https://appspot.com|

## Request

```
POST
```
```
/api/rpc
```
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
        "device_id": "sample_device_id",
        "platform": "ios",
        "token": "sample_token",
        "user_id": "sample_user_id"
    }
}
```

## Response

```
200
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

