package hoiiogo

import (
  "fmt"
  "time"
)

func CSendFax(app *App, fileBase64 string, recipients []string, batchSize int) {

  totalSize := len(recipients)
  ch := make(chan []string, totalSize)
  
  for i := 0; i < totalSize; {
    for j := 0; j < batchSize && i < len(recipients);  {
      fmt.Println("Sending fax to: " + recipients[i])
      go CSendFaxProcess(app, recipients[i], fileBase64, "", "+6594378817", "", "", "", ch)
      i++
      j++
    }
    time.Sleep(10 * time.Second)
  }
  
  for i := 0; i < totalSize; i++ {
    fmt.Println(<-ch)
  }
}

func CSendFaxProcess(app *App, dest, fileBase64, filename, callerId, faxHeader, tag, notifyURL string, ch chan []string) {
  txn, err := SendFax(app, dest, fileBase64, filename, callerId, faxHeader, tag, notifyURL)
  if (err == nil) {
    ch <- []string{txn.Id(), txn.Status()}
  } else {
    ch <- []string{"", err.Error()}    
  }
}
