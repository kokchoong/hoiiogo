package hoiiogo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const ROOT = "https://secure.hoiio.com/open"

type Client interface {
	Do(url.Values, string) ([]byte, error)
  AppId() string
  AccessToken() string
}

type HoiioClient struct {
	appId string
	accessToken string
}

function (h *HoiioClient) AppId() string {
  return h.appId;
}

function (h *HoiioClient) AccessToken() string {
  return h.accessToken;
}

func NewClient(appId, accessToken string) *HoiioClient {
	return &HoiioClient{appId, accessToken}
}

func (client *HoiioClient) Do(values url.Values, uri string) ([]byte, error) {
	req, err := http.NewRequest("POST", ROOT+uri, strings.NewReader(values.Encode()))
    
	if err != nil {
		return nil, err
	}
  
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpClient := &http.Client{}
  
	res, err := httpClient.Do(req)
    
	if err != nil {
		return nil, err
	}
  
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return body, err
	}
  
	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 500 {
			return body, Error{"Server Error"}
		} else {
			hoiioError := new(HoiioError)
			json.Unmarshal(body, hoiioError)
			return body, hoiioError
		}
	}

	return body, err
}