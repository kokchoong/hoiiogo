package hoiiogo

import (
  "net/http"
  "net/url"
  "strings"
)

const (
  ROOT = "https://secure.hoiio.com/open"
  
  SMS_SEND = "/sms/send"
  SMS_BULK_SEND = "/sms/bulk_send"
  
  FAX_SEND = "/fax/send"
  
  IVR_DIAL = "/ivr/start/dial"
  IVR_PLAY = "/ivr/middle/play"
  IVR_GATHER = "/ivr/middle/gather"
  IVR_RECORD = "/ivr/middle/record"
  IVR_MONITOR = "/ivr/middle/monitor"
  IVR_HOLD = "/ivr/middle/hold"
  IVR_UNHOLD = "/ivr/middle/unhold"
  IVR_TRANSFER = "/ivr/end/transfer"
  IVR_HANGUP = "/ivr/end/hangup"
)

type App struct {
  appId string
  accessToken string
}

func (h *App) AppId() string {
  return h.appId
}

func (h *App) AccessToken() string {
  return h.accessToken
}

func NewApp(appId, accessToken string) *App {
  return &App{appId, accessToken}
}

func (h *App) Do(values url.Values, uri string) (txn *Txn, err error) {
  
  defer CatchPanic(err)
  
  req, err := http.NewRequest("POST", ROOT+uri, strings.NewReader(values.Encode()))
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  
  httpClient := &http.Client{}
  res, err := httpClient.Do(req)
  defer res.Body.Close()
  
  return NewTxn(res)
}