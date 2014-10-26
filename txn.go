package hoiiogo

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type Txn struct {
  Id string `json:"txn_ref"`
  Status string `json:"status"`
  statusCode int
}

func (h *Txn) GetId() string {
  return h.Id
}

func (h *Txn) GetStatus() string {
  return h.Status
}

func (h *Txn) GetStatusCode() int {
  return h.statusCode
}

func NewTxn(res *http.Response) (txn *Txn, err error) {
  defer CatchPanic(err)
  body, err := ioutil.ReadAll(res.Body)
  txn = new(Txn)
  err = json.Unmarshal(body, txn)
  txn.statusCode = res.StatusCode
  return txn, err
}

func (txn *Txn) PrintStatus() string {
  return txn.GetStatus()
}