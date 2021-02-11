package solrcluster

// SolrCollectionList
type SolrCollectionList struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
	} `json:"responseHeader"`
	Collections []string `json:"collections"`
}

// SolrCollection defines the desired state of SolrCollection
type SolrCollection struct {
	// The name of the collection to perform the action on
	Collection string `json:"collection"`

	// Define a configset to use for the collection. Use '_default' if you don't have a custom configset
	CollectionConfigName string `json:"collectionConfigName"`

	// The router name that will be used. The router defines how documents will be distributed
	// +optional
	RouterName CollectionRouterName `json:"routerName,omitempty"`

	// If this parameter is specified, the router will look at the value of the field in an input document
	// to compute the hash and identify a shard instead of looking at the uniqueKey field.
	// If the field specified is null in the document, the document will be rejected.
	// +optional
	RouterField string `json:"routerField,omitempty"`

	// The num of shards to create, used if RouteName is compositeId
	// +optional
	NumShards int64 `json:"numShards,omitempty"`

	// The replication factor to be used
	// +optional
	ReplicationFactor int64 `json:"replicationFactor,omitempty"`

	// Max shards per node
	// +optional
	MaxShardsPerNode int64 `json:"maxShardsPerNode,omitempty"`

	// A comma separated list of shard names, e.g., shard-x,shard-y,shard-z. This is a required parameter when the router.name is implicit
	// +optional
	Shards string `json:"shards,omitempty"`

	// When set to true, enables automatic addition of replicas when the number of active replicas falls below the value set for replicationFactor
	// +optional
	AutoAddReplicas bool `json:"autoAddReplicas,omitempty"`
}

// CollectionRouterName is a string enumeration type that enumerates the ways that documents can be routed for a collection.
type CollectionRouterName string

const (
	// The Implicit router
	ImplicitRouter CollectionRouterName = "implicit"

	// The CompositeId router
	CompositeIdRouter CollectionRouterName = "compositeId"
)

// SolrCollectionAlias defines the desired state of SolrCollectionAlias
type SolrCollectionAlias struct {
	// A reference to the SolrCloud to create alias on
	SolrCloud string `json:"solrCloud"`

	// Collections is a list of collections to apply alias to
	Collections []string `json:"collections"`

	// AliasType is a either standard or routed, right now we support standard
	AliasType string `json:"aliasType"`
}
