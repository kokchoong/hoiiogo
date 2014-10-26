package hoiiogo

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type Txn struct {
  TxnId string `json:"txn_ref"`
  TxnStatus string `json:"status"`
  StatusCode int
}

func NewTxn(res *http.Response) (txn *Txn, err error) {
  defer CatchPanic(err)
  body, err := ioutil.ReadAll(res.Body)
  txn = new(Txn)
  err = json.Unmarshal(body, txn)
  txn.StatusCode = res.StatusCode
  return txn, err
}

func (txn *Txn) Id() string {
  return txn.TxnId
}

func (txn *Txn) Status() string {
  return txn.TxnStatus
}