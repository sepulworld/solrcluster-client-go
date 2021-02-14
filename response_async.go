package solrcluster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// GetRequestStatus Request the status and response of an already submitted Asynchronous Collection API call. This call is also used to clear up the stored statuses.
func GetRequestStatus(host string, requestID string) (SolrAsyncResponse, error) {
	asyncResponse := SolrAsyncResponse{}

	resp, err := Get(fmt.Sprintf("%s/solr/admin/collections?action=REQUESTSTATUS&requestid=1000&wt=json", host))
	if err != nil {
		return SolrAsyncResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body. ", err)
	}

	err = json.Unmarshal(body, &asyncResponse)
	if err != nil {
		return SolrAsyncResponse{}, err
	}

	return asyncResponse, nil

}
