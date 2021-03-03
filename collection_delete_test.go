package solrcluster

import (
	"bytes"
	"encoding/json"
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

func TestCollectionDelete(t *testing.T) {
	jsonData := `{
		"responseHeader":{
		  "status":0,
		  "QTime":875},
		"success":{
		  "example-solrcloud-0.example-solrcloud-headless.default:8983_solr":{
			"responseHeader":{
			  "status":0,
			  "QTime":315}}}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))
	host := "http://localhost"
	collection := "testCollection"

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, err := DeleteCollection(host, collection)
	if err != nil {
		fmt.Println(err)
	}
	bs, _ := json.Marshal(resp.Success)
	var respHeader struct {
		Status int "json:\"status\""
		QTime  int "json:\"QTime\""
	} = resp.ResponseHeader
	assert.Contains(t, string(bs), "example-solrcloud-0.example-solrcloud-headless.default")
	assert.Equal(t, respHeader.Status, 0)
}
