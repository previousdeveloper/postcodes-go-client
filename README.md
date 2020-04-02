# postcodes-go-client

=======================

Go library for interacting with the excellent [Postcodes.io][1] service.

### Usage

    config := postcode.DefaultConfig()
    postCodeClient := postcode.NewPostCode(config)
    a := postCodeClient.BulkLookupPostcode(&model.BulkPostCodeRequest{Postcodes: []string{"OX49 5NU"}})
	fmt.Println(a.Result[0].Result)

----------

Installation
-------------

After installing, you can use `go doc` to get documentation:

```bash
go get github.com/previousdeveloper/postcodes-go-client
```
### Contributions

Feel free to make contributions via a pull request. Please keep the tests current (and add, if necessary).


[1]: http://postcodes.io/
