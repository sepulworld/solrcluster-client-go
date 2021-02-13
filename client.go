package solrcluster

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client interface
var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// Post sends a post request to the URL with the body
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return Client.Do(request)
}

// Get request to the URL
func Get(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
		return nil, err
	}
	resp, err := Client.Do(request)
	if err != nil {
		log.Fatal("Error reading response. ", err)
		return nil, err
	}

	return resp, err
}
