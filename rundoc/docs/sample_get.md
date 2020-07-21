# sample_get

サンプルAPIの詳細

## Request

|ENV|URL|
|---|---|
|Local|`GET` http://localhost:8080/sample|
|Staging|`GET` https://staging.appspot.com/sample|
|Production|`GET` https://appspot.com/sample|

```
Authorization: xxxxxxxxxx
Content-Type: application/json
X-OS: iOS
```
```json
{
    "fuga": 1,
    "hoge": "xxxxx"
}
```

## Response

```
Status 200
```
```json
{
    "fuga": 1,
    "hoge": "aaaaaa"
}
```

