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


#### Create Collection

```go
        host := "http://default-example-solrcloud.ing.local.domain"
        collection := solrcluster.SolrCollection{}
        collection.Collection = "power-house-search-collection""
        collection.CollectionConfigName = "_default"
        collection.RouterName = "compositeId"
        collection.NumShards = 2

        createCollection, err := solrcluster.CreateCollection(host, collection)
        if err != nil {
                log.Fatal("Error creating solrcloud. ", err)
        }
        fmt.Println("Collection Created: ", createCollection)
```
