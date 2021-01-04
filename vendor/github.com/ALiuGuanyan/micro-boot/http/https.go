package http

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"strconv"
	"time"
)

var (
	defaultHTTPSPrefix = "https"
	defaultHTTPSName = ""
	defaultHTTPSRunnable = false
	defaultHTTPSPort = ""
	defaultHTTPSReadTimeout = 15 * time.Second
	defaultHTTPSReadHeaderTimeout = 15 * time.Second
	defaultHTTPSWriteTimeout = 15 * time.Second
	defaultHTTPSIdleTimeout time.Duration = 0
	defaultHTTPSMaxHeaderBytes = 0
)

func SetDefaultHTTPSFlagsPrefix(prefix string)  {
	defaultHTTPSPrefix = prefix
}

func SetDefaultHTTPSName(name string)  {
	defaultHTTPSName = name
}

func SetDefaultHTTPSRunnable(b bool)  {
	defaultHTTPSRunnable = b
}

func SetDefaultHTTPSPort(val string)  {
	defaultHTTPSPort = val
}

func SetDefaultHTTPSReadTimeout(val time.Duration)  {
	defaultHTTPSReadTimeout = val
}

func SetDefaultHTTPSReadHeaderTimeout(val time.Duration)  {
	defaultHTTPSReadHeaderTimeout = val
}

func SetDefaultHTTPSWriterTimeout(val time.Duration)  {
	defaultHTTPSWriteTimeout = val
}

func SetDefaultHTTPSIdleTimeout(val time.Duration)  {
	defaultHTTPSIdleTimeout = val
}

func SetDefaultHTTPSMaxHeaderBytes(val int)  {
	defaultHTTPSMaxHeaderBytes = val
}

type HTTPS struct {
	Name string `json:"name" yaml:"name"`
	Runnable          bool          `json:"runnable" yaml:"runnable"`
	Port  string `json:"port" yaml:"port"`
	ReadTimeout time.Duration `json:"read-timeout" yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout time.Duration  `json:"write-timeout" yaml:"write-timeout"`
	IdleTimeout time.Duration  `json:"idle-timeout" yaml:"idle-timeout"`
	MaxHeaderBytes int `json:"max-header-bytes" yaml:"max-header-bytes"`

	Cert string `json:"cert" yaml:"cert"`
	Key string `json:"key" yaml:"key"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (https *HTTPS) GetIntPort() (port int) {
	var port64 int64
	port64, _ = strconv.ParseInt(https.Port, 10, 64)
	return int(port64)
}

func (https *HTTPS) BindFlags(fs *bootflag.FlagSet)  {
	if https.CustomBindFlagsFunc != nil {
		https.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&https.Name,
		utils.BuildFlagName(defaultHTTPSPrefix, "name"),
		defaultHTTPSName,
		"set the https service name")

	fs.BoolVar(&https.Runnable, utils.BuildFlagName(defaultHTTPSPrefix, "runnable"), defaultHTTPSRunnable, "run HTTPS server or not (default is false)")

	fs.IPv4PortStringVar(&https.Port,  utils.BuildFlagName(defaultHTTPSPrefix, "port"), defaultHTTPSPort, `HTTPS server listening port` )

	fs.DurationVar(&https.ReadTimeout, utils.BuildFlagName(defaultHTTPSPrefix,"read-timeout"), defaultHTTPSReadTimeout, "the maximum duration for reading the entire request, including the body")

	fs.DurationVar(&https.ReadHeaderTimeout,utils.BuildFlagName(defaultHTTPSPrefix,"read-header-timeout"), defaultHTTPSReadHeaderTimeout, "the amount of time allowed to read request headers")

	fs.DurationVar(&https.WriteTimeout, utils.BuildFlagName(defaultHTTPSPrefix,"write-timeout"), defaultHTTPSWriteTimeout, "the maximum duration before timing out writes of the response")

	fs.DurationVar(&https.IdleTimeout,utils.BuildFlagName(defaultHTTPSPrefix,"idle-timeout"), defaultHTTPSIdleTimeout, "the maximum amount of time to wait for the next request when keep-alives are enabled.")

	fs.IntVar(&https.MaxHeaderBytes, utils.BuildFlagName(defaultHTTPSPrefix,"max-header-bytes"), defaultHTTPSMaxHeaderBytes, "the maximum number of bytes the server will read parsing the request.")

	fs.StringVar(&https.Cert, utils.BuildFlagName(defaultHTTPSPrefix,"cert"), "", "HTTPS cert file")

	fs.StringVar(&https.Key, utils.BuildFlagName(defaultHTTPSPrefix,"key"), "", "HTTPS key file")
}

func (https *HTTPS) Parse() (err error) {
	if https.CustomParseFunc != nil {
		return https.CustomParseFunc()
	}
	return nil
}