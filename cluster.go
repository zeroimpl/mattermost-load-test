package ltops

type ClusterConfig struct {
	Name                  string
	AppInstanceType       string
	AppInstanceCount      int
	DBInstanceType        string
	DBInstanceCount       int
	LoadtestInstanceCount int
	WorkingDirectory      string
}

// Represents an active cluster
type Cluster interface {
	// Returns the name of the cluster
	Name() string

	// Returns the current configuration of the cluster
	Configuration() *ClusterConfig

	// Returns the SSH private key to connect to the cluster's instances
	SSHKey() []byte

	// Returns the siteURL to connect to the cluster
	SiteURL() string

	// Retuns a slice of the IP addresses of the app server instances in this cluster
	GetAppInstancesAddrs() ([]string, error)

	// Retuns a slice of the IP addresses of the loadtest instances in this cluster
	GetLoadtestInstancesAddrs() ([]string, error)

	// Retuns a slice of the IP addresses of the proxy instances in this cluster
	GetProxyInstancesAddrs() ([]string, error)

	// Returns the master databame connection string
	DBConnectionString() string

	// Returns a list of all the read-replica database connection strings
	DBReaderConnectionStrings() []string

	// Deploys a mattermost package to the cluster.
	DeployMattermost(mattermostFile string, licenceFile string) error

	// Deploys a loadtest package to the cluster.
	DeployLoadtests(loadtestsFile string) error

	// Runs a loadtest
	Loadtest() error

	// Destroys the cluster
	Destroy() error

	// Modifies the configuration of an active Mattermost deployment
	//ModifyMattermostConfig(cluster Cluster, mattermostConfig string) error

	// Runs loadtests against the cluster. Must have deployed mattermost and loadtests
	//ModifyMattermostConfig(cluster Cluster, mattermostConfig string) error
}
