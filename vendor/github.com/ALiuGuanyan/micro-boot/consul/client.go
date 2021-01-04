package consul

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"github.com/hashicorp/consul/api"
	"time"
)

var (
	defaultClientFlagsPrefix = defaultConsulFlagsPrefix + "-client"
	defaultAddress = ""
	defaultScheme = "http"
	defaultDatacenter = ""
	defaultAuthUsername = ""
	defaultAuthPassword = ""
	defaultToken = ""
	defaultTokenFile = ""
	defaultNamespace = ""

	defaultWaitTime = time.Millisecond * 0
)

func SetDefaultClientFlagsPrefix(val string)  {
	defaultClientFlagsPrefix = val
}

func SetDefaultAddress(val string)  {
	defaultAddress = val
}

func SetDefaultScheme(val string)  {
	defaultScheme = val
}

func SetDefaultDatacenter(val string)  {
	defaultDatacenter = val
}

func SetDefaultAuthUsername(val string)  {
	defaultAuthUsername = val
}

func SetDefaultAuthPassword(val string)  {
	defaultAuthPassword = val
}

func SetDefaultWaitTime(val time.Duration)  {
	defaultWaitTime = val
}

func SetDefaultToken(val string)  {
	defaultToken = val
}

func SetDefaultTokenFile(val string)  {
	defaultTokenFile = val
}

func SetDefaultNamespace(val string)  {
	defaultNamespace = val
}

// HttpBasicAuth is used to authenticate http client with HTTP Basic Authentication
type HttpBasicAuth struct {
	// Username to use for HTTP Basic Authentication
	Username string  `json:"username" yaml:"username"`

	// Password to use for HTTP Basic Authentication
	Password string `json:"password" yaml:"password"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}


func (hba *HttpBasicAuth) BindFlags(fs *bootflag.FlagSet)  {
	if hba.CustomBindFlagsFunc != nil {
		hba.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&hba.Username,
		utils.BuildFlagName(defaultClientFlagsPrefix,  "auth-username"),
		defaultAuthUsername,
		"use for HTTP Basic Authentication username")

	fs.StringVar(
		&hba.Password,
		utils.BuildFlagName(defaultClientFlagsPrefix,  "auth-password"),
		defaultAuthPassword,
		"use for HTTP Basic Authentication password")
}

func (hba *HttpBasicAuth) Parse() (err error) {
	if hba.CustomParseFunc != nil {
		return hba.CustomParseFunc()
	}

	return nil
}

func (hba *HttpBasicAuth) Standardize() *api.HttpBasicAuth {
	return &api.HttpBasicAuth{
		Username: hba.Username,
		Password: hba.Password,
	}
}


// ClientConfig is used to configure the creation of a client
type ClientConfig struct {
	// Address is the address of the Consul server
	Address string  `json:"address" yaml:"address"`

	// Scheme is the URI scheme for the Consul server
	Scheme string `json:"scheme" yaml:"scheme"`

	// Datacenter to use. If not provided, the default agent datacenter is used.
	Datacenter string `json:"datacenter" yaml:"datacenter"`

	// HttpAuth is the auth info to use for http access.
	HttpAuth HttpBasicAuth `json:"auth" yaml:"auth"`

	// WaitTime limits how long a Watch will block. If not provided,
	// the agent default values will be used.
	WaitTime time.Duration `json:"wait-time" yaml:"wait-time"`

	// Token is used to provide a per-request ACL token
	// which overrides the agent's default token.
	Token string `json:"token" yaml:"token"`

	// TokenFile is a file containing the current token to use for this client.
	// If provided it is read once at startup and never again.
	TokenFile string `json:"token-file" yaml:"token-file"`

	// Namespace is the name of the namespace to send along for the request
	// when no other Namespace is present in the QueryOptions
	Namespace string `json:"namespace" yaml:"namespace"`

	TLSConfig TLSConfig `json:"tls" yaml:"tls"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (c *ClientConfig) BindFlags(fs *bootflag.FlagSet)  {
	if c.CustomBindFlagsFunc != nil {
		c.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&c.Address,
		utils.BuildFlagName(defaultClientFlagsPrefix, "address"),
		defaultAddress,
		" the address of the Consul server")

	fs.StringVar(
		&c.Scheme,
		utils.BuildFlagName(defaultClientFlagsPrefix, "scheme"),
		defaultScheme,
		"the URI scheme for the Consul server")

	fs.StringVar(
		&c.Datacenter,
		utils.BuildFlagName(defaultClientFlagsPrefix, "datacenter"),
		defaultDatacenter,
		"specify the datacenter to use")

	c.HttpAuth.BindFlags(fs)

	fs.DurationVar(
		&c.WaitTime,
		utils.BuildFlagName(defaultClientFlagsPrefix, "wait-time"),
		defaultWaitTime,
		"limits how long a Watch will block")

	fs.StringVar(
		&c.Token,
		utils.BuildFlagName(defaultClientFlagsPrefix, "token"),
		defaultToken,
		"used to provide a per-request ACL token which overrides the agent's default token")

	fs.StringVar(
		&c.TokenFile,
		utils.BuildFlagName(defaultClientFlagsPrefix, "token-file"),
		defaultTokenFile,
		"a file containing the current token to use for this client.")

	fs.StringVar(
		&c.Namespace,
		utils.BuildFlagName(defaultClientFlagsPrefix, "namespace"),
		defaultNamespace,
		"the name of the namespace to send along for the request when no other Namespace is present in the QueryOptions")

	c.TLSConfig.BindFlags(fs)
}

func (c *ClientConfig) Parse()  (err error) {
	if c.CustomParseFunc != nil {
		return c.CustomParseFunc()
	}

	if err = c.HttpAuth.Parse(); err != nil {
		return err
	}

	if err = c.TLSConfig.Parse(); err != nil {
		return err
	}

	return nil
}

func (c *ClientConfig) Standardize() *api.Config {
	return &api.Config{
		Address:    c.Address,
		Scheme:     c.Scheme,
		Datacenter: c.Datacenter,
		HttpAuth:   c.HttpAuth.Standardize(),
		WaitTime:   c.WaitTime,
		Token:      c.Token,
		TokenFile:  c.TokenFile,
		Namespace:  c.Namespace,
		TLSConfig:  c.TLSConfig.Standardize(),
	}
}

