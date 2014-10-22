package hoiiogo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
  "fmt"
)

const ROOT = "https://secure.hoiio.com/open"

type Client interface {
	AppId() string
	AccessToken() string
	get(url.Values, string) ([]byte, error)
	post(url.Values, string) ([]byte, error)
}

type HoiioClient struct {
	appId string
	accessToken string
}

func NewClient(appId, accessToken string) *HoiioClient {
  fmt.Printf("creating new client")
	return &HoiioClient{appId, accessToken}
}

func (client *HoiioClient) post(values url.Values, uri string) ([]byte, error) {
	req, err := http.NewRequest("POST", ROOT+uri, strings.NewReader(values.Encode()))
  
  fmt.Printf("here in client post")
  
	if err != nil {
		return nil, err
	}
  
  fmt.Printf("Posting 1 ")
  
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpClient := &http.Client{}
fmt.Printf("Posting 2 ")
  
	res, err := httpClient.Do(req)
fmt.Printf("Posting 3 ")
  
  
	if err != nil {
		return nil, err
	}
  
  fmt.Printf("Posting 4 ")

	defer res.Body.Close()

  fmt.Printf("Posting 5 ")

	body, err := ioutil.ReadAll(res.Body)

  fmt.Printf("\n\n%s\n\n", string(body))
  fmt.Printf("Posting 6 ")

	if err != nil {
		return body, err
	}
  
    fmt.Printf("Posting 7 ")
  
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

func (client *HoiioClient) get(queryParams url.Values, uri string) ([]byte, error) {
	var params *strings.Reader

	if queryParams == nil {
		queryParams = url.Values{}
	}

	params = strings.NewReader(queryParams.Encode())
	req, err := http.NewRequest("GET", ROOT+uri, params)

	if err != nil {
		return nil, err
	}

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

func (client *HoiioClient) AppId() string {
	return client.appId
}

func (client *HoiioClient) AccessToken() string {
	return client.accessToken
}