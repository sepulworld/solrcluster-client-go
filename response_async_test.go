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

func TestResponseStatus(t *testing.T) {
	json := `{
		"responseHeader":{
		  "status":0,
		  "QTime":38},
		"status":{
		  "state":"success",
		  "msg":"found 1000 in completed tasks"}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	host := "http://localhost"

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, err := GetRequestStatus(host, "1000")
	if err != nil {
		fmt.Println(err)
	}

	var asyncState string = resp.Status.AsyncState
	var respHeader struct {
		Status int "json:\"status\""
		QTime  int "json:\"QTime\""
	} = resp.ResponseHeader
	assert.Equal(t, asyncState, "success")
	assert.Equal(t, respHeader.Status, 0)
	assert.Equal(t, respHeader.QTime, 38)
}
