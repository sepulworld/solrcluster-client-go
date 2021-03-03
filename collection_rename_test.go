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

func TestCollectionRename(t *testing.T) {
	jsonData := `{
		"responseHeader":{
		  "status":0,
		  "QTime":875}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))
	host := "http://localhost"
	collection := "testCollection"
	targetName := "newCollectionName"

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, err := RenameCollection(host, collection, targetName)
	if err != nil {
		fmt.Println(err)
	}
	var respHeader struct {
		Status int "json:\"status\""
		QTime  int "json:\"QTime\""
	} = resp.ResponseHeader
	assert.Equal(t, respHeader.Status, 0)
}
