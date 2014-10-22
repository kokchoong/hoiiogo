package hoiiogo

import (
	"encoding/json"
	"net/url"
  "fmt"
)

const API_SMS_URL = "sms/send"

type HoiioResponse struct {
  TxnId string `json:"txn_ref"`
	Status string `json:"status"`
}

func SendMessage(client Client, dest, msg, senderName, tag, notifyURL string) (*HoiioResponse, error) {
	fmt.Printf("send message 1")
  var hoiioResponse *HoiioResponse
  
  params := url.Values{}
  params.Set("app_id", client.AppId())
  params.Set("access_token", client.AccessToken())
	params.Set("dest", dest)
  params.Set("dest", dest)
	params.Set("sender_name", senderName)
	params.Set("msg", msg)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  
  fmt.Printf("send message 2")
	
  res, err := client.post(params, API_SMS_URL)

  fmt.Printf("send message 3")
  
	if err != nil {
    fmt.Printf("send message 4")
		return hoiioResponse, err
	}

	hoiioResponse = new(HoiioResponse)
	err = json.Unmarshal(res, hoiioResponse)

	return hoiioResponse, err
}