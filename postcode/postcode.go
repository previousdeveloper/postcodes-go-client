package postcode

import (
	"encoding/json"
	"fmt"
	"github.com/previousdeveloper/postcodes-go-client/model"
	"net/http"
	"time"
)

type Client interface {
	LookupPostcode(postCodeParam string) *model.PostCodeModel
	BulkLookupPostcode(request *model.BulkPostCodeRequest) *model.BulkPostCodeResponse
	GetNearestPostPostcode(longitude string, latitude string) *model.PostCodeModel
	BulkReverseGeocoding(request *model.BulkReverseGeocodingRequest) *model.PostCodeModel
}

type postCodeClient struct {
	Config            *configuration
	httpClientWrapper HttpClientWrapper
}

func (p *postCodeClient) LookupPostcode(postCodeParam string) *model.PostCodeModel {
	var postCode model.PostCodeModel
	result := p.httpClientWrapper.Get(postCodeParam)
	if err := json.Unmarshal(result, &postCode); err != nil {
		panic(err)
	}
	return &postCode
}

func (p *postCodeClient) BulkLookupPostcode(request *model.BulkPostCodeRequest) *model.BulkPostCodeResponse {
	var postCode model.BulkPostCodeResponse
	result := p.httpClientWrapper.Post(request)
	if err := json.Unmarshal(result, &postCode); err != nil {
		panic(err)
	}
	return &postCode
}

func (p *postCodeClient) GetNearestPostPostcode(longitude string, latitude string) *model.PostCodeModel {
	var postCode model.PostCodeModel
	result := p.httpClientWrapper.Get(fmt.Sprintf("?lon=%s&lat=%s", longitude, latitude))
	if err := json.Unmarshal(result, &postCode); err != nil {
		panic(err)
	}
	return &postCode
}

func (p *postCodeClient) BulkReverseGeocoding(request *model.BulkReverseGeocodingRequest) *model.PostCodeModel {
	var postCode model.PostCodeModel
	result := p.httpClientWrapper.Post(request)
	if err := json.Unmarshal(result, &postCode); err != nil {
		panic(err)
	}
	return &postCode
}

func NewPostCode(config *configuration) Client {
	return &postCodeClient{Config: config, httpClientWrapper: config.HttpClientWrapper}
}

type configuration struct {
	BaseUrl           string
	Timeout           time.Duration
	Transport         *http.Transport
	HttpClientWrapper HttpClientWrapper
}

func DefaultConfig() *configuration {
	return defaultConfig(func() *http.Transport {
		return &http.Transport{}
	})
}

func defaultConfig(transportFn func() *http.Transport) *configuration {
	return &configuration{BaseUrl: "api.postcodes.io/", Timeout: 3 * time.Second, Transport: transportFn(), HttpClientWrapper: newHttpClientWrapper(DefaultConfig())}
}
