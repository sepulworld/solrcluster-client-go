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
func CreateCollection(host string, s SolrCollection) (SolrCollectionAPIResponse, SolrAsyncResponse, error) {
	r := SolrCollectionAPIResponse{}
	a := SolrAsyncResponse{}
	queryParams := url.Values{}
	queryParams.Set("wt", "json")

	queryURL := host + "/solr/admin/collections?action=CREATE&"
	if s.Collection != "" {
		queryParams.Add("name", s.Collection)
	} else {
		return r, a, errors.New("error: collection name required")
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

	if s.Shards != "" {
		queryParams.Add("shards", s.Shards)
	}

	if s.ReplicationFactor != 0 {
		numReplicationFactor := strconv.Itoa(s.ReplicationFactor)
		queryParams.Add("replicationFactor", numReplicationFactor)
	}

	if s.NRTReplicas != 0 {
		numNRTReplicas := strconv.Itoa(s.NRTReplicas)
		queryParams.Add("nrtReplicas", numNRTReplicas)
	}

	if s.TlogReplicas != 0 {
		numTlogReplicas := strconv.Itoa(s.TlogReplicas)
		queryParams.Add("tlogReplicas", numTlogReplicas)

	}

	if s.PullReplicas != 0 {
		numPullReplicas := strconv.Itoa(s.PullReplicas)
		queryParams.Add("pullReplicas", numPullReplicas)
	}

	if s.MaxShardsPerNode != 0 {
		numMaxShardsPerNode := strconv.Itoa(s.MaxShardsPerNode)
		queryParams.Add("maxShardsPerNode", numMaxShardsPerNode)
	}

	if s.CreateNodeSet != "" {
		queryParams.Add("createNodeSet", s.CreateNodeSet)
	}

	if s.CreateNodeSetShuffle {
		strCreateNodeSetShuffle := strconv.FormatBool(s.CreateNodeSetShuffle)
		queryParams.Add("createNodeSet.shuffle", strCreateNodeSetShuffle)
	}

	if s.RouterField != "" {
		queryParams.Add("router.field", s.RouterField)
	}

	if len(s.CorePropertySettings) > 0 {
		for _, prop := range s.CorePropertySettings {
			queryParams.Add("property."+prop.Name, prop.Value)
		}
	}

	if s.AutoAddReplicas {
		strAutoAddReplicas := strconv.FormatBool(s.AutoAddReplicas)
		queryParams.Add("autoAddReplicas", strAutoAddReplicas)
	}

	if s.Policy != "" {
		queryParams.Add("policy", s.Policy)
	}

	if s.WaitForFinalState {
		queryParams.Add("waitforFinalState", strconv.FormatBool(s.WaitForFinalState))
	}

	if s.WithCollection != "" {
		queryParams.Add("withCollection", s.WithCollection)
	}

	if s.RouterName != "" {
		queryParams.Add("router.name", s.RouterName)
	}

	if s.Alias != "" {
		queryParams.Add("alias", s.Alias)
	}

	resp, err := Get(queryURL + queryParams.Encode())
	if err != nil {
		return SolrCollectionAPIResponse{}, SolrAsyncResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body. ", err, "query", queryURL+queryParams.Encode())
	}

	if s.Async != "" {
		err = json.Unmarshal(body, &a)
		if err != nil {
			return r, a, err
		}
	} else {
		err = json.Unmarshal(body, &r)
		if err != nil {
			return r, a, err
		}
	}

	return r, a, nil
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
		return r, errors.New("error: collection name to delete required")
	}

	resp, err := Get(queryURL + queryParams.Encode())
	if err != nil {
		return SolrCollectionAPIResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading response body. ", err)
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
		return r, errors.New("error: collection name to rename required required")
	}

	if targetName != "" {
		queryParams.Add("name", targetName)
	} else {
		return r, errors.New("error: collection target name to rename to required required")
	}

	resp, err := Get(queryURL + queryParams.Encode())
	if err != nil {
		return SolrCollectionAPIResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading response body. ", err)
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

// ReloadCollection The RELOAD action is used when you have changed a configuration in ZooKeeper.
func ReloadCollection(host string, collection string, asyncID string) (SolrCollectionAPIResponse, SolrAsyncResponse, error) {
	r := SolrCollectionAPIResponse{}
	a := SolrAsyncResponse{}
	queryParams := url.Values{}
	queryParams.Set("wt", "json")

	queryURL := host + "/solr/admin/collections?action=RELOAD&"
	if collection != "" {
		queryParams.Add("name", collection)
	} else {
		return r, a, errors.New("error: collection name to reload required")
	}

	if asyncID != "" {
		queryParams.Add("async", asyncID)
	}

	resp, err := Get(queryURL + queryParams.Encode())
	if err != nil {
		return SolrCollectionAPIResponse{}, SolrAsyncResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading response body. ", err)
	}

	if asyncID != "" {
		err = json.Unmarshal(body, &a)
		if err != nil {
			return r, a, err
		}
	} else {
		err = json.Unmarshal(body, &r)
		if err != nil {
			return r, a, err
		}
	}

	return r, a, err
}

// ModifyCollection It’s possible to edit multiple attributes at a time. Changing these values only updates the z-node on ZooKeeper,
/*
   they do not change the topology of the collection. For instance, increasing replicationFactor will not automatically add more replicas to the collection but will allow more ADDREPLICA commands to succeed.
   An attribute can be deleted by passing an empty value. For example, yet_another_attribute_name= (with no value) will delete the yet_another_attribute_name parameter from the collection.
   At least one attribute parameter is required.

   The attributes that can be modified are:
   maxShardsPerNode
   replicationFactor
   autoAddReplicas
   collection.configName
   policy
   withCollection
   readOnly
   other custom properties that use a property. prefix
*/
func ModifyCollection(host, collection string, collectionAttributes []CollectionAttribute) (SolrCollectionAPIResponse, error) {
	r := SolrCollectionAPIResponse{}
	queryParams := url.Values{}
	queryParams.Set("wt", "json")

	queryURL := host + "/solr/admin/collections?action=MODIFYCOLLECTION&"
	if collection != "" {
		queryParams.Add("name", collection)
	} else {
		return r, errors.New("error: collection name to modify required")
	}

	if len(collectionAttributes) > 0 {
		for _, attr := range collectionAttributes {
			queryParams.Add(attr.Attribute, attr.Value)
		}
	} else {
		return r, errors.New("error: must provide atleast one attribute to modify")
	}
	resp, err := Get(queryURL + queryParams.Encode())
	if err != nil {
		return r, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading response body. ", err)
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, err
}
