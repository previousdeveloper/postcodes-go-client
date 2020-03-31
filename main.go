package main

import (
	"fmt"
	"github.com/previousdeveloper/postcodes-go-client/model"
	"github.com/previousdeveloper/postcodes-go-client/postcode"
	"time"
)

func main() {
	config := postcode.DefaultConfig()
	config.Timeout = 1001 * time.Millisecond
	postCodeClient := postcode.NewPostCode(config)
	//postcodeResult := postCodeClient.LookupPostcode("NE1%204LF")
	a := postCodeClient.BulkLookupPostcode(&model.BulkPostCodeRequest{Postcodes: []string{"OX49 5NU"}})
	fmt.Println(a.Result[0].Result)
}
