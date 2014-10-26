package hoiiogo

import (
  "net/url"
)

func SendBulkSMS(app *App, recipients, msg, senderName, tag, notifyURL string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("recipients", recipients)
  params.Set("msg", msg)
  params.Set("sender_name", senderName)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, SMS_BULK_SEND)
}

func SendSMS(app *App, dest, msg, senderName, tag, notifyURL string) (*Txn, error) {
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("dest", dest)
  params.Set("msg", msg)
  params.Set("sender_name", senderName)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, SMS_SEND)
}