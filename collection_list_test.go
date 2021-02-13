package solrcluster

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/sepulworld/solrcluster-client-go/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func init() {
	Client = &mocks.MockClient{}
}

func TestCollectionList(t *testing.T) {
	json := `{
		"responseHeader":{
		  "status":0,
		  "QTime":2011},
		"collections":["collection1",
		  "example1",
		  "example2"]}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	host := "http://localhost"

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, err := GetCollections(host)
	if err != nil {
		fmt.Println(err)
	}

	var collections []string = resp.Collections
	var respHeader struct {
		Status int "json:\"status\""
		QTime  int "json:\"QTime\""
	} = resp.ResponseHeader
	assert.Contains(t, collections, "collection1")
	assert.Contains(t, collections, "example1")
	assert.Contains(t, collections, "example2")
	assert.Equal(t, respHeader.Status, 0)
	assert.Equal(t, respHeader.QTime, 2011)
}
