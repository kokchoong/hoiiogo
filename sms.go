package hoiiogo

import (
	"encoding/json"
	"net/url"
)

const API_SMS_URL = "/sms/send"

type HoiioResponse struct {
  TxnId string `json:"txn_ref"`
	Status string `json:"status"`
}

func SendMessage(client Client, senderName, dest, msg, tag, notifyURL string) (*HoiioResponse, error) {

  var hoiioResponse *HoiioResponse
  
  params := url.Values{}
  params.Set("app_id", client.appId)
  params.Set("access_token", client.accessToken)
	params.Set("dest", dest)
  params.Set("dest", dest)
	params.Set("sender_name", senderName)
	params.Set("msg", msg)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  
  res, err := client.do(params, API_SMS_URL)

	hoiioResponse = new(HoiioResponse)
	err = json.Unmarshal(res, hoiioResponse)

	return hoiioResponse, err
}