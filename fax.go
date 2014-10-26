package hoiiogo

import (
  "encoding/json"
  "net/url"
  "encoding/base64"
  "io/ioutil"
)

func SendFaxByFilePath(client Client, dest, filepath, filename, callerId, faxHeader, tag, notifyURL string) (*HoiioTxn, error) {
  
  file, err := ioutil.ReadFile(filepath)
  
  if err != nil {
    return nil, err
  }
  
  fileBase64 := EncodeFile(file)
  return SendFax(client , dest, fileBase64, filename, callerId, faxHeader, tag, notifyURL)
}

func SendFax(client Client, dest, fileBase64, filename, callerId, faxHeader, tag, notifyURL string) (*HoiioTxn, error) {
  
  params := url.Values{}
  params.Set("app_id", client.AppId())
  params.Set("access_token", client.AccessToken())
  params.Set("file", fileBase64)
  params.Set("filename", filename)
  params.Set("caller_id", callerId)
  params.Set("fax_header", faxHeader)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  
  res, err := client.Do(params, FAX_API)
  
  hoiioTxn := new(HoiioTxn)
  err = json.Unmarshal(res, hoiioTxn)
  
  return hoiioTxn, err
}

func EncodeFile(bytes []byte) (b64 string) {
  b64 = base64.URLEncoding.EncodeToString(bytes)
  return b64
}