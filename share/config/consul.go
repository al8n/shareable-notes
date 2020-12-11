package config

import (
	bootconsul "github.com/ALiuGuanyan/micro-boot/consul"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"time"
)

type Agent struct {
	IP       string  `json:"ip" yaml:"ip"`
	Tags     []string  `json:"tags" yaml:"tags"`
	Port     int  `json:"port" yaml:"port"`
	DeregisterCriticalServiceAfter time.Duration `json:"" yaml:""`
	Interval  time.Duration `json:"interval" yaml:"interval"`
}

func (a *Agent) BindFlags(fs *bootflag.FlagSet)  {

	fs.StringVar(&a.IP, "consul-agent-ip", "", "specify the ip for consul agent")
	fs.IntVar(&a.Port, "consul-agent-port", 8500, "specify the port for consul agent")
	fs.DurationVar(&a.Interval, "consul-agent-interval", 10 * time.Second, "specify the port for consul agent")
}

type Consul struct {
	Agent Agent `json:"agent" yaml:"agent"`
	Client bootconsul.ClientConfig `json:"client" yaml:"client"`
}

func (c *Consul) BindFlags(fs *bootflag.FlagSet)  {
	c.Agent.BindFlags(fs)
	c.Client.BindFlags(fs)
}

func (c *Consul) Parse() (err error) {
	return c.Client.Parse()
}