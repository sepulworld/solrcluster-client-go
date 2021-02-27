package solrcluster

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
)

// GetCollections - Returns collection list and response header status and query time
func GetCollections(host string) (SolrCollectionList, error) {
	collections := SolrCollectionList{}

	queryURL := host + "/solr/admin/collections?action=LIST"

	resp, err := Get(queryURL)
	if err != nil {
		return SolrCollectionList{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body. ", err)
	}

	err = json.Unmarshal(body, &collections)
	if err != nil {
		return SolrCollectionList{}, err
	}

	return collections, nil
}

// CreateCollection - Returns collection API request status
func CreateCollection(host string, s SolrCollection) (SolrCollectionCreateResponse, error) {
	r := SolrCollectionCreateResponse{}
	queryParams := url.Values{}
	queryParams.Set("wt", "json")

	queryURL := host + "/solr/admin/collections?action=CREATE&"
	if s.Collection == "" {
		return r, errors.New("Error: Collection name required")
	} else {
		queryParams.Add("name", s.Collection)
	}

	if s.CollectionConfigName != "" {
		queryParams.Add("collection.configName", s.CollectionConfigName)
	}

	if s.RouterName != "" {
		queryParams.Add("router.name", s.RouterName)
	}

	if s.NumShards != 0 {
		numShardsString := strconv.Itoa(s.NumShards)
		queryParams.Add("numShards", numShardsString)
	}

	if s.RouterName != "" {
		queryParams.Add("router.name", s.RouterName)
	}

	resp, err := Get(queryURL + queryParams.Encode())
	if err != nil {
		return SolrCollectionCreateResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body. ", err)
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
