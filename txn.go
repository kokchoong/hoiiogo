package hoiiogo

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "fmt"
)

type Txn struct {
  TxnId string `json:"txn_ref"`
  TxnStatus string `json:"status"`
  HttpStatusCode int
}

func NewTxn(res *http.Response) (txn *Txn, err error) {
  defer CatchPanic(err)
  body, err := ioutil.ReadAll(res.Body)
  txn = new(Txn)
  err = json.Unmarshal(body, txn)
  txn.HttpStatusCode = res.StatusCode
  fmt.Println(txn.Id() + " - " + txn.Status())
  return txn, err
}

func (txn *Txn) Id() string {
  return txn.TxnId
}

func (txn *Txn) Status() string {
  return txn.TxnStatus
}

func (txn *Txn) StatusCode() int {
  return txn.HttpStatusCode
}