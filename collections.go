package solrcluster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// GetCollections - Returns collection list and response header status and query time
func GetCollections(host string) (SolrCollectionList, error) {
	collections := SolrCollectionList{}

	resp, err := Get(fmt.Sprintf("%s/solr/admin/collections?action=LIST", host))
	if err != nil {
		return SolrCollectionList{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	err = json.Unmarshal(body, &collections)
	if err != nil {
		return SolrCollectionList{}, err
	}

	return collections, nil
}
