package hoiiogo

import (
  "net/url"
)

func SendFaxByFilePath(app *App, dest, filepath, filename, callerId, faxHeader, tag, notifyURL string) (*Txn, error) {
  file, err := ReadFile(filepath)
  if err != nil {
    return nil, err
  }
  fileBase64 := EncodeFile(file)
  return SendFax(app, dest, fileBase64, filename, callerId, faxHeader, tag, notifyURL)
}

func SendFax(app *App, dest, fileBase64, filename, callerId, faxHeader, tag, notifyURL string) (*Txn, error) {
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("dest", dest)
  params.Set("file", fileBase64)
  params.Set("filename", filename)
  params.Set("caller_id", callerId)
  params.Set("fax_header", faxHeader)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, FAX_SEND)
}

func ReadFaxStatus(app *App, txn *Txn) (*Txn, error) {
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("txn_ref", txn.Id())
  return app.Do(params, FAX_SEND)
}