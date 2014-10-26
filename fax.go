package hoiiogo

import (
  "net/url"
  "io/ioutil"
)

func SendFaxByFilePath(app Service, dest, filepath, filename, callerId, faxHeader, tag, notifyURL string) (*Txn, error) {
  
  file, err := ioutil.ReadFile(filepath)
  
  if err != nil {
    return nil, nil
  }
  
  fileBase64 := EncodeFile(file)
  return SendFax(app, dest, fileBase64, filename, callerId, faxHeader, tag, notifyURL)
}

func SendFax(app Service, dest, fileBase64, filename, callerId, faxHeader, tag, notifyURL string) (*Txn, error) {
  
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("file", fileBase64)
  params.Set("filename", filename)
  params.Set("caller_id", callerId)
  params.Set("fax_header", faxHeader)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  
  return app.Do(params, FAX_API)
}