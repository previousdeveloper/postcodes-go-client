package postcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type HttpClientWrapper interface {
	Get(path string) []byte
	Post(param interface{}) []byte
}
type httpClientWrapper struct {
	BaseUrl string
	Http    *http.Client
}

func (wrapper *httpClientWrapper) Get(path string) []byte {

	response, err := wrapper.Http.Get(wrapper.BaseUrl + "/postcodes/" + path)
	if err != nil {
		panic(err)
	}

	defer func() {
		if response.Body != nil {
			closeResult := response.Body.Close()
			if closeResult != nil {
				fmt.Println("Error")
			}
		}
	}()
	body, err := ioutil.ReadAll(response.Body)

	return body
}

func (wrapper *httpClientWrapper) Post(param interface{}) []byte {

	jsonValue, _ := json.Marshal(param)
	response, err := wrapper.Http.Post(wrapper.BaseUrl+"/postcodes/", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	defer func() {
		if response.Body != nil {
			closeResult := response.Body.Close()
			if closeResult != nil {
				fmt.Println("Error")
			}
		}
	}()
	body, err := ioutil.ReadAll(response.Body)

	return body
}

func newHttpClientWrapper(config *configuration) HttpClientWrapper {
	return &httpClientWrapper{BaseUrl: "http://api.postcodes.io/", Http: &http.Client{Timeout: config.Timeout, Transport: config.Transport}}
}
