package solrcluster

import (
	"fmt"
	"testing"
)

func TestCollection(t *testing.T) {
	type m interface{}
	host := "http://default-example-solrcloud.ing.local.domain"
	c, err := NewClient(&host)
	if err != nil {
		fmt.Println(err)
	}
	collections, err := c.GetCollections()
	fmt.Println("Collection", collections)
}
