package postcode

import (
	"github.com/previousdeveloper/postcodes-go-client/model"
	"testing"
)
import "github.com/golang/mock/gomock"
import mock "github.com/previousdeveloper/postcodes-go-client/mocks"

//go:generate mockgen -destination=../mocks/mock-http-client-wrapper.go -source ../postcode/httpclient.go HttpClientWrapper

func TestPostCodeClient_LookupPostcode(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)

	mockHttpClientWrapper := mock.NewMockHttpClientWrapper(mockCtrl)

	text := "{\"status\": 200}"
	bytes := []byte(text)
	mockHttpClientWrapper.EXPECT().Get("postCode").Return(bytes)
	configuration := &configuration{HttpClientWrapper: mockHttpClientWrapper}
	newPostCode := NewPostCode(configuration)
	postcodeResult := newPostCode.LookupPostcode("postCode")

	if postcodeResult.Status != 200 {
		t.Error("er")
	}
}

func TestPostCodeClient_GetNearestPostPostcode(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	mockHttpClientWrapper := mock.NewMockHttpClientWrapper(mockCtrl)
	text := "{\"status\": 200}"
	bytes := []byte(text)
	mockHttpClientWrapper.EXPECT().Get("?lon=a&lat=b").Return(bytes)
	configuration := &configuration{HttpClientWrapper: mockHttpClientWrapper}
	newPostCode := NewPostCode(configuration)
	result := newPostCode.GetNearestPostPostcode("a", "b")

	if result.Status != 200 {
		t.Error("er")
	}
}

func TestPostCodeClient_BulkLookupPostcode(t *testing.T) {

	t.Parallel()
	mockCtrl := gomock.NewController(t)
	mockHttpClientWrapper := mock.NewMockHttpClientWrapper(mockCtrl)
	text := "{\"Status\": 200}"
	bytes := []byte(text)
	request := &model.BulkPostCodeRequest{}
	configuration := &configuration{HttpClientWrapper: mockHttpClientWrapper}

	mockHttpClientWrapper.EXPECT().Post(request).Return(bytes)
	newPostCode := NewPostCode(configuration)
	result := newPostCode.BulkLookupPostcode(request)

	if result.Status != 200 {
		t.Error("er")
	}
}

func TestPostCodeClient_BulkReverseGeocoding(t *testing.T) {

	t.Parallel()
	mockCtrl := gomock.NewController(t)
	mockHttpClientWrapper := mock.NewMockHttpClientWrapper(mockCtrl)
	text := "{\"Status\": 200}"
	bytes := []byte(text)
	request := &model.BulkReverseGeocodingRequest{}
	configuration := &configuration{HttpClientWrapper: mockHttpClientWrapper}

	mockHttpClientWrapper.EXPECT().Post(request).Return(bytes)
	newPostCode := NewPostCode(configuration)
	result := newPostCode.BulkReverseGeocoding(request)

	if result.Status != 200 {
		t.Error("er")
	}
}

