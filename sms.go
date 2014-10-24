package hoiiogo

import (
	"encoding/json"
	"net/url"
)

const API_SMS_URL = "/sms/send"

type SMSTxn struct {
  TxnId string `json:"txn_ref"`
	Status string `json:"status"`
}

func SendSMS(client Client, dest, msg, senderName, tag, notifyURL string) (*SMSTxn, error) {

  var smsTxn *SMSTxn
  
  params := url.Values{}
  params.Set("app_id", client.AppId())
  params.Set("access_token", client.AccessToken())
  params.Set("dest", dest)
	params.Set("msg", msg)
  params.Set("sender_name", senderName)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  
  res, err := client.Do(params, API_SMS_URL)

	smsTxn = new(SMSTxn)
	err = json.Unmarshal(res, smsTxn)

	return smsTxn, err
}