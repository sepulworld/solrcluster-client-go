package solrcluster

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sepulworld/solrcluster-client-go/utils/mocks"
)

func ExampleGetCollections() {
	json := `{
		"responseHeader":{
		  "status":0,
		  "QTime":2011},
		"collections":["collection1",
		  "example1",
		  "example2"]}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	host := "http://default-example-solrcloud.ing.local.domain"

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	resp, err := GetCollections(host)
	if err != nil {
		log.Fatal("Error connecting to solrcloud. ", err)
	}
	fmt.Println(resp.Collections)

	// Output: [collection1 example1 example2]
}
