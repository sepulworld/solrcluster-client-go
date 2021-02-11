package solrcluster

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCollections - Returns list of collections with responseheader data
func (c *Client) GetCollections() ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/solr/admin/collections?action=LIST", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	collections := SolrCollectionList{}
	err = json.Unmarshal(body, &collections)
	if err != nil {
		return nil, err
	}

	return collections.Collections, nil

}
