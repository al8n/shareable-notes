package mongo

import (
	"fmt"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"time"
)

// string values
var (
	defaultMongoFlagsPrefix = "mongo"
	defaultReplicaSet = ""
	defaultAppName = ""
	defaultDatabase = ""
	defaultCollection = ""
	defaultURI = ""
)

func SetDefaultMongoFlagsPrefix(val string)  {
	defaultMongoFlagsPrefix = val
}

func SetDefaultAppName(val string)  {
	defaultAppName = val
}

func SetDefaultDatabase(val string)  {
	defaultDatabase = val
}

func SetDefaultCollection(val string)  {
	defaultCollection = val
}

func SetDefaultReplicaSet(val string)  {
	defaultReplicaSet = val
}

func SetDefaultURI(val string)  {
	defaultURI = val
}

// number values
var (
	defaultMaxPoolSize uint64 = 100
	defaultMinPoolSize uint64 = 0
	defaultZlibLevel = -1
	defaultZstdLevel = 6
)

func SetDefaultMaxPoolSize(val uint64)  {
	defaultMaxPoolSize =val
}

func SetDefaultMinPoolSize(val uint64)  {
	defaultMinPoolSize =val
}

func SetDefaultZlibLevel(val int)  {
	defaultZlibLevel = val
}

func SetDefaultZstdLevel(val int)  {
	defaultZstdLevel = val
}

// string slice values
var (
	defaultMongoCompressors = []string{}
	defaultMongoDBHosts = []string{}
)

func SetDefaultCompressors(val []string)  {
	defaultMongoCompressors = val
}

func SetDefaultHosts(val []string)  {
	defaultMongoDBHosts = val
}

// bool values
var (
	defaultDisableOCSPEndpointCheck = false
	defaultDirect = false
	defaultRetryReads = true
	defaultRetryWrites = true
)

func SetDefaultDirect(val bool)  {
	defaultDirect = val
}

func SetDefaultDisableOCSPEndpointCheck(val bool)  {
	defaultDisableOCSPEndpointCheck = val
}

func SetDefaultRetryReads(val bool)  {
	defaultRetryReads = val
}

func SetDefaultRetryWrites(val bool)  {
	defaultRetryWrites = val
}

// duration values
var (
	defaultMongoConnectionTimeout = 30 * time.Second
	defaultHeartbeatInterval = 10 * time.Second
	defaultLocalThreshold = 15 * time.Millisecond
	defaultMaxConnIdleTime time.Duration = 0
	defaultSeverSelectTimeout = 30 * time.Second
	defaultSocketTimeout time.Duration = 0
)

func SetDefaultConnectionTimeout(val time.Duration)  {
	defaultMongoConnectionTimeout = val
}

func SetDefaultHeartbeatInterval(val time.Duration)  {
	defaultHeartbeatInterval = val
}

func SetDefaultLocalThreshold(val time.Duration)  {
	defaultLocalThreshold = val
}

func SetDefaultMaxConnIdleTime(val time.Duration)  {
	defaultMaxConnIdleTime = val
}

func SetDefaultSeverSelectTimeout(val time.Duration)  {
	defaultSeverSelectTimeout = val
}

func SetDefaultSocketTimeout(val time.Duration)  {
	defaultSocketTimeout = val
}

type ClientOptions struct {
	AppName                  string      `json:"app-name" yaml:"app-name"`
	Auth                     Credential `json:"auth" yaml:"auth"`

	//AutoEncryptionOptions    *AutoEncryptionOptions  `json:"auto-encryption" yaml:"auto-encryption"`

	ConnectTimeout           time.Duration    `json:"connect-timeout" yaml:"connect-timeout"`
	Compressors              []string			`json:"compressors" yaml:"compressors"`
	Collection 				string        `json:"collection" yaml:"collection"`
	DB         				string        `json:"database" yaml:"database"`
	Direct                   bool				`json:"direct" yaml:"direct"`
	DisableOCSPEndpointCheck bool				`json:"disable-ocsp-endpoint-check" yaml:"disable-ocsp-endpoint-check"`
	HeartbeatInterval        time.Duration     `json:"heartbeat-interval" yaml:"heartbeat-interval"`
	Hosts 					[]string 			`json:"hosts" yaml:"hosts"`
	LocalThreshold           time.Duration    `json:"local-threshold" yaml:"local-threshold"`
	MaxConnIdleTime          time.Duration    `json:"max-conn-idle-time" yaml:"max-conn-idle-time"`
	MaxPoolSize              uint64 			`json:"max-pool-size" yaml:"max-pool-size"`
	MinPoolSize              uint64 			`json:"min-pool-size" yaml:"min-pool-size"`

	// ReadConcern for replica sets and replica set shards determines which data to return from a query.
	ReadConcern              ReadConcern `json:"read-concern" yaml:"read-concern"`

	// ReadPreference determines which servers are considered suitable for read operations.
	ReadPreference           ReadPref `json:"read-preference" yaml:"read-preference"`

	ReplicaSet               string  			`json:"replica-set" yaml:"replica-set"`

	RetryReads               bool  			`json:"retry-reads" yaml:"retry-reads"`

	RetryWrites              bool 				`json:"retry-writes" yaml:"retry-writes"`

	ServerSelectionTimeout   time.Duration     `json:"server-selection-timeout" yaml:"server-selection-timeout"`

	SocketTimeout            time.Duration      `json:"socket-timeout" yaml:"socket-timeout"`

	URI 					string 				`json:"uri" yaml:"uri"`

	ZlibLevel                int 				`json:"zlib-level" yaml:"zlib-yaml"`

	ZstdLevel                int 				`json:"zstd-level" yaml:"zstd-yaml"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (m *ClientOptions) BindFlags(fs *bootflag.FlagSet) {
	if m.CustomBindFlagsFunc != nil {
		m.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&m.AppName,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "app-name"),
		defaultAppName,
		"specifies an application name that is sent to the server when creating new connections")

	fs.StringVar(
		&m.DB,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "database"),
		defaultDatabase,
		"specify the database used by micro services")

	fs.StringVar(
		&m.Collection,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "collection"),
		defaultCollection,
		"specify the collection used by micro services")

	m.Auth.BindFlags(fs)

	fs.DurationVar(
		&m.ConnectTimeout,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "connect-timeout"),
		defaultMongoConnectionTimeout,
		"specifies a timeout that is used for creating connections to the server")

	fs.StringSliceVar(
		&m.Compressors,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "compressors"),
		defaultMongoCompressors,
		"sets the compressors that can be used when communicating with a server")

	fs.BoolVar(
		&m.Direct,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "direct"),
		defaultDirect, "specifies whether or not a direct connect should be made")

	fs.BoolVar(
		&m.DisableOCSPEndpointCheck,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "disable-ocsp-endpoint-check"),
		defaultDisableOCSPEndpointCheck,
		"specifies whether or not the driver should reach out to OCSP responders to verify")

	fs.DurationVar(
		&m.HeartbeatInterval,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "heartbeat-interval"),
		defaultHeartbeatInterval,
		"specifies the amount of time to wait between periodic background server checks")

	fs.StringSliceVar(
		&m.Hosts,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "hosts"),
		defaultMongoDBHosts,
		"specifies a list of host names or IP addresses for servers in a cluster")

	fs.DurationVar(
		&m.LocalThreshold,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "local-threshold"),
		defaultLocalThreshold,
		"specifies the width of the 'latency window'")

	fs.DurationVar(
		&m.MaxConnIdleTime,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "max-conn-idle-time"),
		defaultMaxConnIdleTime,
		"specifies the maximum amount of time that a connection will remain idle in a connection pool")

	fs.Uint64Var(
		&m.MaxPoolSize,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "max-pool-size"),
		defaultMaxPoolSize,
		"specifies that maximum number of connections allowed in the driver's connection pool to each server")

	fs.Uint64Var(
		&m.MinPoolSize,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "min-pool-size"),
		defaultMinPoolSize,
		"specifies that minimum number of connections allowed in the driver's connection pool to each server")

	m.ReadConcern.BindFlags(fs)

	m.ReadPreference.BindFlags(fs)

	fs.StringVar(
		&m.URI,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "uri"),
		defaultURI,
		"specify uri of servers")

	fs.StringVar(
		&m.ReplicaSet,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "replica-set"),
		defaultReplicaSet,
		"specifies the replica set name for the cluster")

	fs.BoolVar(
		&m.RetryReads,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "retry-reads"),
		defaultRetryReads, "specifies whether supported read operations should be retried once on certain errors")

	fs.BoolVar(
		&m.RetryWrites,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "retry-writes"),
		defaultRetryWrites, "specifies whether supported write operations should be retried once on certain errors")

	fs.DurationVar(
		&m.ServerSelectionTimeout,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "server-select-timeout"),
		defaultSeverSelectTimeout,
		"specifies how long the driver will wait to find an available, suitable server to execute an operation")

	fs.DurationVar(
		&m.SocketTimeout,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "socket-timeout"),
		defaultSocketTimeout,
		"specifies how long the driver will wait for a socket read or write to return before returning a network error")

	fs.IntVar(
		&m.ZlibLevel,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "zlib-level"),
		defaultZlibLevel,
		"specifies the level for the zlib compressor")

	fs.IntVar(
		&m.ZstdLevel,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "zstd-level"),
		defaultZstdLevel,
		"sets the level for the zstd compressor")

}

func (m *ClientOptions) Parse() (err error) {
	if m.URI != "" {
		_, err = connstring.Parse(m.URI)
		if err != nil {
			return fmt.Errorf("%s is not a valid MongoDB URI", m.URI)
		}
	}

	return nil
}

func (m *ClientOptions) Standardize() (opts *options.ClientOptions, err error) {
	opts = &options.ClientOptions{}

	opts.SetAppName(m.AppName)



	if m.Auth.Username != "" && m.Auth.Password != "" {

		auth := options.Credential{
			Username:                m.Auth.Username,
			Password:                m.Auth.Password,
			PasswordSet:             m.Auth.PasswordSet,
		}

		if m.Auth.AuthMechanism != "" {
			auth.AuthMechanism = m.Auth.AuthMechanism
		}

		if m.Auth.AuthMechanismProperties != nil {
			auth.AuthMechanismProperties = m.Auth.AuthMechanismProperties
		}

		if m.Auth.AuthSource != "" {
			auth.AuthSource = m.Auth.AuthSource
		}

		opts.SetAuth(auth)
	}


	opts.SetCompressors(m.Compressors)
	opts.SetConnectTimeout(m.ConnectTimeout)
	opts.SetDirect(m.Direct)
	opts.SetDisableOCSPEndpointCheck(m.DisableOCSPEndpointCheck)
	opts.SetHeartbeatInterval(m.HeartbeatInterval)
	opts.SetHosts(m.Hosts)
	opts.SetLocalThreshold(m.LocalThreshold)
	opts.SetMaxConnIdleTime(m.MaxConnIdleTime)
	opts.SetMaxPoolSize(m.MaxPoolSize)
	opts.SetMinPoolSize(m.MinPoolSize)
	opts.SetReadConcern(readconcern.New(readconcern.Level(m.ReadConcern.Level)))

	pref, err := m.ReadPreference.Standardize()
	if err != nil {
		return nil, err
	}
	opts.SetReadPreference(pref)

	opts.SetReplicaSet(m.ReplicaSet)
	opts.SetRetryReads(m.RetryReads)
	opts.SetRetryWrites(m.RetryWrites)
	opts.SetServerSelectionTimeout(m.ServerSelectionTimeout)
	opts.SetSocketTimeout(m.SocketTimeout)
	opts.SetZlibLevel(m.ZlibLevel)
	opts.SetZstdLevel(m.ZstdLevel)

	if m.URI != "" {
		opts.ApplyURI(m.URI)
	}

	return
}