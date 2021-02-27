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
	Collection string

	/*
		    If not provided, Solr will use the configuration of _default configset to create a new (and mutable) configset named <collectionName>.AUTOCREATED and will use it for the new collection.
			When such a collection is deleted, its autocreated configset will be deleted by default when it is not in use by any other collection.
	*/
	CollectionConfigName string

	// The router name that will be used. The router defines how documents will be distributed
	RouterName string

	// If this parameter is specified, the router will look at the value of the field in an input document
	// to compute the hash and identify a shard instead of looking at the uniqueKey field.
	// If the field specified is null in the document, the document will be rejected.
	RouterField string

	// The num of shards to create, used if RouteName is compositeId
	NumShards int

	// The replication factor to be used
	ReplicationFactor int64

	// Max shards per node
	MaxShardsPerNode int64

	// A comma separated list of shard names, e.g., shard-x,shard-y,shard-z. This is a required parameter when the router.name is implicit
	Shards string

	// When set to true, enables automatic addition of replicas when the number of active replicas falls below the value set for replicationFactor
	AutoAddReplicas bool

	// The number of NRT (Near-Real-Time) replicas to create for this collection. This type of replica maintains a transaction log and updates its index locally. If you want all of your replicas to be of this type, you can simply use replicationFactor instead.
	NRTReplicas int64

	// The number of TLOG replicas to create for this collection. This type of replica maintains a transaction log but only updates its index via replication from a leader. See the section Types of Replicas for more information about replica types.
	TlogReplicas int64

	// The number of PULL replicas to create for this collection. This type of replica does not maintain a transaction log and only updates its index via replication from a leader. This type is not eligible to become a leader and should not be the only type of replicas in the collection. See the section Types of Replicas for more information about replica types.
	PullReplicas int64

	/*
			Allows defining the nodes to spread the new collection across.
		    The format is a comma-separated list of node_names, such as localhost:8983_solr,localhost:8984_solr,localhost:8985_solr.
		    If not provided, the CREATE operation will create shard-replicas spread across all live Solr nodes.
			Alternatively, use the special value of EMPTY to initially create no shard-replica within the new collection and then later use the ADDREPLICA operation to add shard-replicas when and where required.
	*/
	CreateNodeSet string

	/*
		Controls whether or not the shard-replicas created for this collection will be assigned to the nodes specified by the createNodeSet in a sequential manner, or if the list of nodes should be shuffled prior to creating individual replicas.
		A false value makes the results of a collection creation predictable and gives more exact control over the location of the individual shard-replicas, but true can be a better choice for ensuring replicas are distributed evenly across nodes. The default is true.
		This parameter is ignored if createNodeSet is not also specified.
	*/
	CreateNodeSetShuffle bool

	// Set core property name to value. See the section Defining core.properties for details on supported properties and values.
	CorePropertySettings []SolrCoreProperty

	// Request ID to track this action which will be processed asynchronously.
	Async string

	// Name of the collection-level policy
	Policy string

	/*
		If true, the request will complete only when all affected replicas become active.
		The default is false, which means that the API will return the status of the single action, which may be before the new replica is online and active.
	*/
	WaitForFinalState bool

	// The name of the collection with which all replicas of this collection must be co-located. The collection must already exist and must have a single shard named shard1.
	WithCollection string

	/*
		Starting with version 8.1 when a collection is created additionally an alias can be created that points to this collection.
		This parameter allows specifying the name of this alias
	*/
	Alias string
}

// SolrCoreProperty The core.properties file is a simple Java Properties file where each line is just a key=value pair, e.g., name=core1
type SolrCoreProperty struct {
	Name  string
	Value string
}

// SolrCollectionCreateResponse response interface for collection API CREATE
type SolrCollectionCreateResponse struct {
	ResponseHeader SolrResponseHeader `json:"responseHeader"`
	Success        SolrCoreSuccess    `json:"success,omitempty"`
	Warning        string             `json:"warning,omitempty"`
	Exception      SolrException      `json:"exception,omitempty"`
	Error          SolrError          `json:"error,omitempty"`
}

// SolrCoreResponse response from Solr Core creation
type SolrCoreResponse struct {
	ResponseHeader SolrResponseHeader `json:"responseHeader"`
	Core           string             `json:"core"`
}

// SolrCoreSuccess response from API call
type SolrCoreSuccess struct {
	SolrCoreResponses []SolrCoreResponse
}

// SolrException exception response from API
type SolrException struct {
	Msg     string `json:"msg"`
	RspCode int    `json:"rspCode"`
}

// SolrError error message from API
type SolrError struct {
	Metadata []string `json:"metadata"`
	Msg      string   `json:"msg"`
	Code     int      `json:"code"`
}

// SolrCollectionAlias defines the desired state of SolrCollectionAlias
type SolrCollectionAlias struct {
	// A reference to the SolrCloud to create alias on
	SolrCloud string

	// Collections is a list of collections to apply alias to
	Collections []string

	// AliasType is a either standard or routed, right now we support standard
	AliasType string
}
