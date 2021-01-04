package mongo

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
)

var (
	defaultCredentialPrefix = defaultMongoFlagsPrefix + "-auth"
	defaultAuthMechanism = ""
	defaultAuthMechanismProperties = make(map[string]string)
	defaultAuthSource = "admin"
	defaultUsername = ""
	defaultPassword = ""
	defaultPasswordSet = false
)

const (
	defaultAuthMechanismUsage = "the mechanism to use for authentication"
	defaultAuthMechanismPropertiesUsage = "specify additional configuration options for certain mechanisms"
	defaultAuthSourceUsage = "the name of the database to use for authentication"
	defaultUsernameUsage = "the username for authentication"
	defaultPasswordUsage = "the password for authentication"
	defaultPasswordSetUsage = "For GSSAPI mechanism. Other mechanisms, this field is ignored"
	)

func SetDefaultCredentialFlagsPrefix(val string)  {
	defaultCredentialPrefix = val
}

func SetDefaultAuthMechanism (val string)  {
	defaultAuthMechanism = val
}

func SetDefaultAuthMechanismProperties(val map[string]string)  {
	defaultAuthMechanismProperties = val
}

func SetDefaultAuthSource (val string)  {
	defaultAuthSource = val
}

func SetDefaultAuthUsername (val string)  {
	defaultUsername = val
}

func SetDefaultAuthPassword (val string)  {
	defaultPassword = val
}

func SetDefaultAuthPasswordSet (val bool)  {
	defaultPasswordSet = val
}

// Credential can be used to provide authentication options when configuring a Client.
//
// See official go-mongo-driver doc https://godoc.org/go.mongodb.org/mongo-driver/mongo/options#Credential for details.
type Credential struct {
	AuthMechanism           string               `json:"mechanism" yaml:"mechanism"`
	AuthMechanismProperties map[string]string    `json:"mechanism-properties" yaml:"mechanism-properties"`
	AuthSource              string				`json:"source" yaml:"source"`
	Username                string              `json:"username" yaml:"username"`
	Password                string              `json:"password" yaml:"password"`
	PasswordSet             bool                `json:"password-set" yaml:"password-set"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (c *Credential) BindFlags(fs *bootflag.FlagSet)  {
	if c.CustomBindFlagsFunc != nil {
		c.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&c.AuthMechanism,
		utils.BuildFlagName(defaultCredentialPrefix, "mechanism"),
		defaultAuthMechanism,
		defaultAuthMechanismUsage)

	fs.StringToStringVar(
		&c.AuthMechanismProperties,
		utils.BuildFlagName(defaultCredentialPrefix, "mechanism-properties"),
		defaultAuthMechanismProperties,
		defaultAuthMechanismPropertiesUsage)

	fs.StringVar(
		&c.AuthSource,
		utils.BuildFlagName(defaultCredentialPrefix, "source"),
		defaultAuthSource,
		defaultAuthSourceUsage)

	fs.StringVar(
		&c.Username,
		utils.BuildFlagName(defaultCredentialPrefix, "username"),
		defaultUsername,
		defaultUsernameUsage)

	fs.StringVar(
		&c.Password,
		utils.BuildFlagName(defaultCredentialPrefix, "password"),
		defaultPassword,
		defaultPasswordUsage)

	fs.BoolVar(
		&c.PasswordSet,
		utils.BuildFlagName(defaultCredentialPrefix, "password-set"),
		defaultPasswordSet,
		defaultPasswordSetUsage)
}

func (c *Credential) Parse() (err error) {
	if c.CustomParseFunc != nil {
		return c.CustomParseFunc()
	}

	return nil
}

