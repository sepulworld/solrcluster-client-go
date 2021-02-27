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

func TestCollectionCreate(t *testing.T) {
	jsonData := `{
		"responseHeader":{
		  "status":0,
		  "QTime":6599},
		"success":{
		  "example-solrcloud-0.example-solrcloud-headless.default:8983_solr":{
			"responseHeader":{
			  "status":0,
			  "QTime":2152},
			"core":"newCollection_shard1_replica_n1"},
		  "example-solrcloud-1.example-solrcloud-headless.default:8983_solr":{
			"responseHeader":{
			  "status":0,
			  "QTime":5316},
			"core":"newCollection_shard2_replica_n3"}},
		"warning":"Using _default configset. Data driven schema functionality is enabled by default, which is NOT RECOMMENDED for production use. To turn it off: curl http://{host:port}/solr/newCollection/config -d '{\"set-user-property\": {\"update.autoCreateFields\":\"false\"}}'"}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))
	host := "http://localhost"
	collection := SolrCollection{}
	collection.Collection = "newcollection"
	collection.CollectionConfigName = "_default"
	collection.RouterName = "compositeId"
	collection.NumShards = 2

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, err := CreateCollection(host, collection)
	if err != nil {
		fmt.Println(err)
	}
	bs, _ := json.Marshal(resp.Success)
	var respHeader struct {
		Status int "json:\"status\""
		QTime  int "json:\"QTime\""
	} = resp.ResponseHeader
	assert.Contains(t, string(bs), "newCollection_shard1_replica_n1")
	assert.Equal(t, respHeader.Status, 0)
}
