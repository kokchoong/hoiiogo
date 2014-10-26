package hoiiogo

import (
  "encoding/base64"
  "io/ioutil"
  "fmt"
)

func ReadFile(filepath string) (file []byte, err error) {

  file, err = ioutil.ReadFile(filepath)
  defer CatchPanic(err)
  
  return file, err
}

func EncodeFile(bytes []byte) (string) {
  return base64.URLEncoding.EncodeToString(bytes)
}

func CatchPanic(err error) {
  if r := recover()
  r != nil {
    err = fmt.Errorf("%v", r)
  }
}