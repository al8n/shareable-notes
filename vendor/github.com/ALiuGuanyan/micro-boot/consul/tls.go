package consul

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"github.com/hashicorp/consul/api"
)

var (
	defaultTlSFlagsPrefix = defaultClientFlagsPrefix + "-tls"
	defaultTlSAddress = ""
	defaultTlSCAFile = ""
	defaultTlSCAPath = ""
	defaultTlSCAPem = []byte{}

	defaultTlSCertFile = ""
	defaultTlSCertPem = []byte{}

	defaultTlSKeyFile = ""
	defaultTlSKeyPem = []byte{}
	defaultTlSInsecureSkipVerify = false
)

func SetDefaultTLSFlagsPrefix(val string)  {
	defaultTlSFlagsPrefix = val
}

func SetDefaultTLSAddress(val string)  {
	defaultTlSAddress = val
}

func SetDefaultTLSCertFile(val string)  {
	defaultTlSCertFile = val
}

func SetDefaultTLSCAPath(val string)  {
	defaultTlSCAPath = val
}

func SetDefaultTLSCAFile(val string)  {
	defaultTlSCAFile = val
}

func SetDefaultTLSKeyFile(val string)  {
	defaultTlSKeyFile = val
}

func SetDefaultTLSInsecureSkipVerify(val bool)  {
	defaultTlSInsecureSkipVerify = val
}

// TLSConfig is used to generate a TLSClientConfig that's useful for talking to
// Consul using TLS.
type TLSConfig struct {
	// Address is the optional address of the Consul server. The port, if any
	// will be removed from here and this will be set to the ServerName of the
	// resulting config.
	Address string  `json:"address" yaml:"address"`

	// CAFile is the optional path to the CA certificate used for Consul
	// communication, defaults to the system bundle if not specified.
	CAFile string  `json:"ca-file" yaml:"ca-file"`

	// CAPath is the optional path to a directory of CA certificates to use for
	// Consul communication, defaults to the system bundle if not specified.
	CAPath string  `json:"ca-path" yaml:"ca-path"`

	// CAPem is the optional PEM-encoded CA certificate used for Consul
	// communication, defaults to the system bundle if not specified.
	CAPem []byte    `json:"ca-pem" yaml:"ca-pem"`

	// CertFile is the optional path to the certificate for Consul
	// communication. If this is set then you need to also set KeyFile.
	CertFile string `json:"cert-file" yaml:"cert-file"`

	// CertPEM is the optional PEM-encoded certificate for Consul
	// communication. If this is set then you need to also set KeyPEM.
	CertPEM []byte  `json:"cert-pem" yaml:"cert-pem"`

	// KeyFile is the optional path to the private key for Consul communication.
	// If this is set then you need to also set CertFile.
	KeyFile string  `json:"key-file" yaml:"key-file"`

	// KeyPEM is the optional PEM-encoded private key for Consul communication.
	// If this is set then you need to also set CertPEM.
	KeyPEM []byte   `json:"key-pem" yaml:"key-pem"`

	// InsecureSkipVerify if set to true will disable TLS host verification.
	InsecureSkipVerify bool `json:"insecure-skip-verify" yaml:"insecure-skip-verify"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (c *TLSConfig) BindFlags(fs *bootflag.FlagSet) {
	if c.CustomBindFlagsFunc != nil {
		c.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&c.Address,
		utils.BuildFlagName(defaultTlSFlagsPrefix, "address"),
		defaultTlSAddress,
		"the optional address of the Consul server")

	fs.StringVar(
		&c.CAFile,
		utils.BuildFlagName(defaultTlSFlagsPrefix, "ca-file"),
		defaultTlSCAFile,
		"the optional path to the CA certificate used for Consul communication")

	fs.StringVar(
		&c.CAPath,
		utils.BuildFlagName(defaultTlSFlagsPrefix, "ca-path"),
		defaultTlSCAPath,
		"the optional path to a directory of CA certificates to use for Consul communication")

	fs.StringVar(
		&c.CertFile,
		utils.BuildFlagName(defaultTlSFlagsPrefix, "cert-file"),
		defaultTlSCertFile,
		"the optional path to the certificate for Consul communication")

	fs.StringVar(
		&c.KeyFile,
		utils.BuildFlagName(defaultTlSFlagsPrefix, "key-file"),
		defaultTlSKeyFile,
		"the optional path to the private key for Consul communication")

	fs.BoolVar(
		&c.InsecureSkipVerify,
		utils.BuildFlagName(defaultTlSFlagsPrefix, "insecure-skip-verify"),
		defaultTlSInsecureSkipVerify,
		"if set to true will disable TLS host verification")
}

func (c *TLSConfig) Parse() (err error) {
	if c.CustomParseFunc != nil {
		return c.CustomParseFunc()
	}

	return nil
}

func (c *TLSConfig) Standardize() api.TLSConfig {
	return api.TLSConfig{
		Address: c.Address,
		KeyFile: c.KeyFile,
		CAPath: c.CAPath,
		CAFile:  c.CAFile,
		CertFile: c.CertFile,
		InsecureSkipVerify: c.InsecureSkipVerify,
	}
}