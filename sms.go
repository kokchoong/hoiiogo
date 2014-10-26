package hoiiogo

import (
  "encoding/json"
  "net/url"
)

func SendSMS(client Client, dest, msg, senderName, tag, notifyURL string) (*HoiioTxn, error) {
  
  params := url.Values{}
  params.Set("app_id", client.AppId())
  params.Set("access_token", client.AccessToken())
  params.Set("dest", dest)
  params.Set("msg", msg)
  params.Set("sender_name", senderName)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  
  res, err := client.Do(params, SMS_API)
  
  hoiioTxn := new(HoiioTxn)
  err = json.Unmarshal(res, hoiioTxn)
  
  return hoiioTxn, err
}