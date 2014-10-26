package hoiiogo

import (
  "encoding/json"
  "net/url"
)

type SMSTxn struct {
  TxnId string `json:"txn_ref"`
  Status string `json:"status"`
}

func SendSMS(client Client, dest, msg, senderName, tag, notifyURL string) (*SMSTxn, error) {
  
  params := url.Values{}
  params.Set("app_id", client.AppId())
  params.Set("access_token", client.AccessToken())
  params.Set("dest", dest)
  params.Set("msg", msg)
  params.Set("sender_name", senderName)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  
  res, err := client.Do(params, SMS_API)
  
  var smsTxn *SMSTxn
  smsTxn = new(SMSTxn)
  err = json.Unmarshal(res, smsTxn)
  
  return smsTxn, err
}