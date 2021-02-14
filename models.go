package solrcluster

// SolrAsyncResponse for async requests
type SolrAsyncResponse struct {
	ResponseHeader SolrResponseHeader `json:"responseHeader"`

	RequestID string `json:"requestId"`

	Status SolrAsyncStatus `json:"status"`
}

// SolrAsyncStatus state and message
type SolrAsyncStatus struct {
	// Possible states can be found here: https://github.com/apache/lucene-solr/blob/1d85cd783863f75cea133fb9c452302214165a4d/solr/solrj/src/java/org/apache/solr/client/solrj/response/RequestStatusState.java
	AsyncState string `json:"state"`

	Message string `json:"msg"`
}

// SolrResponseHeader http header response status and query time
type SolrResponseHeader struct {
	Status int `json:"status"`

	QTime int `json:"QTime"`
}

// SolrCollectionList defines collection list api response
type SolrCollectionList struct {
	ResponseHeader SolrResponseHeader `json:"responseHeader"`
	Collections    []string           `json:"collections"`
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
	// ImplicitRouter route type
	ImplicitRouter CollectionRouterName = "implicit"

	// CompositeIDRouter route type
	CompositeIDRouter CollectionRouterName = "compositeId"
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
