package model

type BulkReverseGeocodingRequest struct {
	Geolocations []struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
		Radius    int     `json:"radius,omitempty"`
		Limit     int     `json:"limit,omitempty"`
	} `json:"geolocations"`
}
