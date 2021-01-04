package consul

import bootflag "github.com/ALiuGuanyan/micro-boot/flag"

var (
	defaultConsulFlagsPrefix = "consul"
)

type Config struct {
	Client ClientConfig `json:"client" yaml:"client"`
	Agent  AgentServiceRegistration `json:"agent-service-registration" yaml:"agent-service-registration"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (c *Config) BindFlags(fs *bootflag.FlagSet)  {
	if c.CustomBindFlagsFunc != nil {
		c.CustomBindFlagsFunc(fs)
		return
	}

	c.Client.BindFlags(fs)
	c.Agent.BindFlags(fs)
}

func (c *Config) Parse() ( err error) {
	if err = c.Client.Parse(); err != nil {
		return err
	}

	if err = c.Agent.Parse(); err != nil {
		return err
	}

	if c.CustomParseFunc != nil{
		return c.CustomParseFunc()
	}

	return nil
}