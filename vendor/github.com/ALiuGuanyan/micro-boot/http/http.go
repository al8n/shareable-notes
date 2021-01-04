package http

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"strconv"
	"time"
)

var (
	defaultHTTPPrefix = "http"
	defaultHTTPName = ""
	defaultHTTPRunnable = true
	defaultHTTPPort = ""
	defaultHTTPReadTimeout = 15 * time.Second
	defaultHTTPReadHeaderTimeout = 15 * time.Second
	defaultHTTPWriteTimeout = 15 * time.Second
	defaultHTTPIdleTimeout time.Duration = 0
	defaultHTTPMaxHeaderBytes = 0
)

func SetDefaultHTTPFlagsPrefix(prefix string)  {
	defaultHTTPPrefix = prefix
}

func SetDefaultHTTPName(name string)  {
	defaultHTTPName = name
}

func SetDefaultHTTPRunnable(b bool)  {
	defaultHTTPRunnable = b
}

func SetDefaultHTTPPort(val string)  {
	defaultHTTPPort = val
}

func SetDefaultHTTPReadTimeout(val time.Duration)  {
	defaultHTTPReadTimeout = val
}

func SetDefaultHTTPReadHeaderTimeout(val time.Duration)  {
	defaultHTTPReadHeaderTimeout = val
}

func SetDefaultHTTPWriterTimeout(val time.Duration)  {
	defaultHTTPWriteTimeout = val
}

func SetDefaultHTTPIdleTimeout(val time.Duration)  {
	defaultHTTPIdleTimeout = val
}

func SetDefaultHTTPMaxHeaderBytes(val int)  {
	defaultHTTPMaxHeaderBytes = val
}

type HTTP struct {
	Name string `json:"name" yaml:"name"`
	Runnable          bool          `mapstructure:"http" json:"runnable" yaml:"runnable"`
	Port              string        `json:"port" yaml:"port"`
	ReadTimeout       time.Duration `json:"read-timeout" yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout" yaml:"write-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout" yaml:"idle-timeout"`
	MaxHeaderBytes    int           `json:"max-header-bytes" yaml:"max-header-bytes"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (http *HTTP) GetIntPort() (port int) {
	var port64 int64
	port64, _ = strconv.ParseInt(http.Port, 10, 64)
	return int(port64)
}

func (http *HTTP) BindFlags(fs *bootflag.FlagSet)  {
	if http.CustomBindFlagsFunc != nil {
		http.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&http.Name,
		utils.BuildFlagName(defaultHTTPPrefix, "name"),
		defaultHTTPName,
		"set the http service name")

	fs.BoolVar(&http.Runnable, utils.BuildFlagName(defaultHTTPPrefix, "runnable"), defaultHTTPRunnable, "run HTTP server or not (default is true)")
	fs.IPv4PortStringVar(&http.Port,  utils.BuildFlagName(defaultHTTPPrefix, "port"), defaultHTTPPort, `HTTPS server listening port` )

	fs.DurationVar(&http.ReadTimeout, utils.BuildFlagName(defaultHTTPPrefix,"read-timeout"), defaultHTTPReadTimeout, "the maximum duration for reading the entire request, including the body")

	fs.DurationVar(&http.ReadHeaderTimeout,utils.BuildFlagName(defaultHTTPPrefix,"read-header-timeout"), defaultHTTPReadHeaderTimeout, "the amount of time allowed to read request headers")
	fs.DurationVar(&http.WriteTimeout, utils.BuildFlagName(defaultHTTPPrefix,"write-timeout"), defaultHTTPWriteTimeout, "the maximum duration before timing out writes of the response")
	fs.DurationVar(&http.IdleTimeout,utils.BuildFlagName(defaultHTTPPrefix,"idle-timeout"), defaultHTTPIdleTimeout, "the maximum amount of time to wait for the next request when keep-alives are enabled.")
	fs.IntVar(&http.MaxHeaderBytes, utils.BuildFlagName(defaultHTTPPrefix,"max-header-bytes"), defaultHTTPMaxHeaderBytes, "the maximum number of bytes the server will read parsing the request.")
}

func (http *HTTP) Parse() (err error) {
	if http.CustomParseFunc != nil {
		return http.CustomParseFunc()
	}
	return nil
}