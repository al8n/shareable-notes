package config

import (
	"encoding/json"
	"errors"
	"github.com/al8n/shareable-notes/share/common"
	bootconsul "github.com/al8n/micro-boot/consul"
	bootflag "github.com/al8n/micro-boot/flag"
	bootgrpc "github.com/al8n/micro-boot/grpc"
	boothttp "github.com/al8n/micro-boot/http"
	bootmongo "github.com/al8n/micro-boot/mongo"
	bootprom "github.com/al8n/micro-boot/prometheus"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var (
	config    *Config
	once sync.Once
)

type Config struct {

	Address string `json:"address" yaml:"address"`
	Host string `json:"host" yaml:"host"`

	// Share service
	Service Share `json:"service" yaml:"service"`

	// HTTP
	HTTP boothttp.HTTP  `json:"http" yaml:"http"`

	// HTTPS
	HTTPS boothttp.HTTPS `json:"https" yaml:"https"`

	// GRPC
	GRPC bootgrpc.GRPC 	`json:"grpc" yaml:"grpc"`

	// Mongo
	Mongo bootmongo.ClientOptions `json:"mongo" yaml:"mongo"`

	// Prometheus
	Prom  bootprom.Config          `json:"prometheus" yaml:"prometheus"`

	Consul bootconsul.Config      `json:"consul" yaml:"consul"`
}

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}
	})
	return config
}

func (c *Config) Initialize(name string) (err error) {
	var (
		extension string
		file []byte
		newCfg Config
	)

	file, err = ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	extension = filepath.Ext(name)
	switch extension {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(file, &newCfg)
		if err != nil {
			return err
		}
	case ".json":
		err = json.Unmarshal(file, &newCfg)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported config file type")
	}

	err = mergo.Merge(config, &newCfg)
	if err != nil {
		return err
	}


	if !config.HTTP.Runnable && !config.HTTPS.Runnable && !config.GRPC.Runnable {
		return common.ErrorNoServicesConfig
	}

	err = config.Parse()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) BindFlags(fs *bootflag.FlagSet)  {
	fs.StringVar(&c.Address, "address", "", "specify the service address. e.g. 127.0.0.1:8080")
	fs.StringVar(&c.Host, "host", "", "specify the service host. e.g. 127.0.0.1")
	c.HTTP.BindFlags(fs)
	c.HTTPS.BindFlags(fs)
	c.GRPC.BindFlags(fs)
	c.Mongo.BindFlags(fs)
	c.Prom.BindFlags(fs)
	c.Service.BindFlags(fs)
	c.Consul.BindFlags(fs)
	return
}

func (c *Config) Parse() (err error) {
	if c.Address == "" {
		return errors.New("invalid address")
	}

	if c.Host == "" {
		return errors.New("invalid host")
	}

	err = c.Mongo.Parse()
	if err != nil {
		return err
	}

	err = c.Service.Parse()
	if err != nil {
		return err
	}

	err = c.Consul.Parse()
	if err != nil {
		return err
	}

	return nil
}





