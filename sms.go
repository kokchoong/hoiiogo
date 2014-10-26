package hoiiogo

import (
  "net/url"
)

func SendSMS(app *App, dest, msg, senderName, tag, notifyURL string) (*Txn, error) {
  
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("dest", dest)
  params.Set("msg", msg)
  params.Set("sender_name", senderName)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  
  return app.Do(params, SMS_API)
}