package hoiiogo

import (
  "net/http"
  "net/url"
  "strings"
)

const (
  ROOT = "https://secure.hoiio.com/open"
  FAX_API = "/fax/send"
  SMS_API = "/sms/send"
)

type Service interface {
  Do(url.Values, string) (*Txn, error)
  AppId() string
  AccessToken() string
}

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

func NewService(appId, accessToken string) *App {
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