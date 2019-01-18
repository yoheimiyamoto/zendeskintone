## 概要
ZendeskとKintoneのwebhook連携

## 機能
* zendeskのチケット情報をkintoneにPostする
* フィールドコードがzendesk, Kintone間で一致している項目のみコピーされる

### 対応フィールド
Kintoneのフィールドが以下以外の場合、Post時にエラーが発生する可能性があります。

|フィールドタイプ|データ型|形式|
|:--|:--|:--|
|文字列（1行）|テキスト|てすと|
|日付|テキスト|YYYY-MM-DD|

## WebhookURL
```
https://zendesk-kintone-webhook-dot-rls-airpay.appspot.com
```

## リクエストボディ（例）
kintone_app_id は必須項目です。  
```
{
  "id":"{{ticket.id}}",
  "description": "{{ticket.description}}",
  "kintone_app_id":"100"
}
```

## 実行例
```
curl -X POST -H "Content-Type:application/json" -d '{"ticket_id":"100","zendesk_created_at":"2018-01-01"}' {URL}
```