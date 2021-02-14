# solrcluster-client-go

![Go](https://github.com/sepulworld/solrcluster-client-go/workflows/Go/badge.svg?branch=main)

[![Go Reference](https://pkg.go.dev/badge/github.com/sepulworld/solrcluster-client-go.svg)](https://pkg.go.dev/github.com/sepulworld/solrcluster-client-go)

Go client for managing SolrClusters

### Collection API

#### List Collections

```go
package main

import (
	"fmt"
	"log"

	solrcluster "github.com/sepulworld/solrcluster-client-go"
)

// Get list of collections on a SolrCluster
func main() {
	host := "http://default-example-solrcloud.ing.local.domain"
	resp, err := solrcluster.GetCollections(host)
	if err != nil {
		log.Fatal("Error connecting to solrcloud. ", err)
	}
	fmt.Println("Collections:", resp.Collections)
}
```
