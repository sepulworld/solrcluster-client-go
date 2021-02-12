package solrcluster

import (
	"fmt"
	"testing"
)

func TestCollection(t *testing.T) {
	host := "http://default-example-solrcloud.ing.local.domain"
	resp, err := GetCollections(host)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Collections", resp.Collections)
	fmt.Println("Resp Headers", resp.ResponseHeader)
}
