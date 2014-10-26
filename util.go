package hoiiogo

import (
  "encoding/base64"
  "fmt"
)

func EncodeFile(bytes []byte) (b64 string) {
  b64 = base64.URLEncoding.EncodeToString(bytes)
  return b64
}

func CatchPanic(err error) {
  if r := recover()
  r != nil {
    err = fmt.Errorf("%v", r)
  }
}