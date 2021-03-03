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
func CreateCollection(host string, s SolrCollection) (SolrCollectionAPIResponse, error) {
	r := SolrCollectionAPIResponse{}
	queryParams := url.Values{}
	queryParams.Set("wt", "json")

	queryURL := host + "/solr/admin/collections?action=CREATE&"
	if s.Collection != "" {
		queryParams.Add("name", s.Collection)
	} else {
		return r, errors.New("Error: Collection name required")
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
		return SolrCollectionAPIResponse{}, err
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

// DeleteCollection delete a SolrCluster collection
func DeleteCollection(host string, collection string) (SolrCollectionAPIResponse, error) {
	r := SolrCollectionAPIResponse{}
	queryParams := url.Values{}
	queryParams.Set("wt", "json")

	queryURL := host + "/solr/admin/collections?action=DELETE&"
	if collection != "" {
		queryParams.Add("name", collection)
	} else {
		return r, errors.New("Error: Collection name to delete required")
	}

	resp, err := Get(queryURL + queryParams.Encode())
	if err != nil {
		return SolrCollectionAPIResponse{}, err
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

// RenameCollection rename a SolrCluster collection
// https://solr.apache.org/guide/8_7/collection-management.html#rename-command-parameters
func RenameCollection(host string, existingName string, targetName string) (SolrCollectionAPIResponse, error) {
	r := SolrCollectionAPIResponse{}
	queryParams := url.Values{}
	queryParams.Set("wt", "json")

	queryURL := host + "/solr/admin/collections?action=RENAME&"
	if existingName != "" {
		queryParams.Add("name", existingName)
	} else {
		return r, errors.New("Error: Collection name to rename required required")
	}

	if targetName != "" {
		queryParams.Add("name", targetName)
	} else {
		return r, errors.New("Error: Collection target name to rename to required required")
	}

	resp, err := Get(queryURL + queryParams.Encode())
	if err != nil {
		return SolrCollectionAPIResponse{}, err
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
