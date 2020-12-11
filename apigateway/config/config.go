package config

import (
	"encoding/json"
	"errors"
	"github.com/ALiuGuanyan/margin/share/common"

	boothttp "github.com/ALiuGuanyan/micro-boot/http"
	bootzipkin "github.com/ALiuGuanyan/micro-boot/zipkin"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"

	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
)

var (
	config *Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}
	})
	return config
}

type Config struct {
	HTTP boothttp.HTTP `json:"http" yaml:"http"`
	HTTPS boothttp.HTTPS `json:"https" yaml:"https"`

	Zipkin bootzipkin.Config `json:"zipkin" yaml:"zipkin"`

	ShareSVC ShareService `json:"share-svc" yaml:"share-svc"`

	ConsulAddr string `json:"consul-addr" yaml:"consul-addr"`
	RetryMax int `json:"retry-max" yaml:"retry-max"`
	RetryTimeout time.Duration `json:"retry-timeout" yaml:"retry-timeout"`
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


	if !config.HTTP.Runnable && !config.HTTPS.Runnable {
		return common.ErrorNoServicesConfig
	}

	err = config.Parse()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) BindFlags(fs *bootflag.FlagSet)  {
	c.HTTP.BindFlags(fs)
	c.HTTPS.BindFlags(fs)
	c.Zipkin.BindFlags(fs)

	fs.StringVar(&c.ConsulAddr, "consul-addr", "", "Consul agent address")
	fs.IntVar(&c.RetryMax, "retry-max", 3, "per-request retries to different instances")
	fs.DurationVar(&c.RetryTimeout, "retry-timeout", 500 * time.Millisecond, "per-request timeout, including retries")
}

func (c *Config) Parse() (err error) {
	return nil
}
