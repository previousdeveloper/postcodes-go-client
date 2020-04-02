package postcode

import "testing"
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

func TestPostCodeClient_Bla(t *testing.T) {
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
