package api

import (
	"errors"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"github.com/sony/gobreaker"
	"strings"
	"time"
)

var (
	ErrorNoAPIName = errors.New("API name is required")
	ErrorNoAPIMakeFunc = errors.New("no endpoint make function is specified")

	defaultAPIFlagsPrefix = ""
	defaultAPIRateLimitFlagsPrefix = defaultAPIFlagsPrefix + "-ratelimit"
	defaultAPIBreakerFlagsPrefix =  defaultAPIFlagsPrefix + "-breaker"

	defaultAPILogKVs = map[string]string{}

	defaultAPIRateLimitDelta = 1000
	defaultAPIRateLimitDuration = time.Second

	defaultAPIBreakerMaxRequests uint32 = 1
	defaultAPIBreakerTimeout = 60 * time.Second
	defaultAPIBreakerInterval time.Duration = 0
)

func SetDefaultAPIRateLimitFlagsPrefix(prefix string)  {
	defaultAPIRateLimitFlagsPrefix = prefix
}

func SetDefaultAPIBreakerFlagsPrefix(prefix string)  {
	defaultAPIBreakerFlagsPrefix = prefix
}

func SetDefaultAPIFlagsPrefix(prefix string)  {
	defaultAPIFlagsPrefix = prefix
}

func SetDefaultAPILogKVs(logkvs map[string]string)  {
	defaultAPILogKVs = logkvs
}

func SetDefaultAPIRateLimitDelta(val int)  {
	defaultAPIRateLimitDelta = val
}

func SetDefaultAPIRateLimitDuration(val time.Duration)  {
	defaultAPIRateLimitDuration = val
}

func SetDefaultAPIBreakerMaxRequests(val uint32)  {
	defaultAPIBreakerMaxRequests = val
}

func SetDefaultAPIBreakerTimeout(val time.Duration)  {
	defaultAPIBreakerTimeout = val
}

func SetDefaultAPIBreakerInterval(val time.Duration)  {
	defaultAPIBreakerInterval = val
}

type RateLimit struct {
	Delta    int   `json:"delta" yaml:"delta"`
	Duration time.Duration `json:"duration" yaml:"duration"`
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`
	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (rl *RateLimit) BindFlags(fs *bootflag.FlagSet)  {
	if rl.CustomBindFlagsFunc != nil {
		rl.CustomBindFlagsFunc(fs)
		return
	}
	fs.IntVar(
		&rl.Delta,
		utils.BuildFlagName(defaultAPIRateLimitFlagsPrefix,"delta"),
		defaultAPIRateLimitDelta,
		"api rate limit delta")

	fs.DurationVar(
		&rl.Duration,
		utils.BuildFlagName(defaultAPIRateLimitFlagsPrefix,"duration"),
		defaultAPIRateLimitDuration,
		"api rate limit duration")
}

func (rl RateLimit) Parse() (err error) {
	if rl.CustomParseFunc != nil {
		return rl.CustomParseFunc()
	}
	return nil
}

type Breaker struct {
	Name        string        `json:"name" yaml:"name"`
	MaxRequests uint32        `json:"max" yaml:"max"`
	Interval    time.Duration `json:"interval" yaml:"interval"`
	Timeout     time.Duration  `json:"timeout" yaml:"timeout"`
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`
	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (b *Breaker) Standardize() gobreaker.Settings {
	return gobreaker.Settings{
		Name:          b.Name,
		MaxRequests:   b.MaxRequests,
		Interval:      b.Interval,
		Timeout:       b.Timeout,
	}
}

func (b *Breaker) BindFlags(fs *bootflag.FlagSet)  {
	if b.CustomBindFlagsFunc != nil {
		b.CustomBindFlagsFunc(fs)
	}

	fs.StringVar(&b.Name, utils.BuildFlagName(defaultAPIBreakerFlagsPrefix, "name"), "", "the name of the CircuitBreaker")
	fs.Uint32Var(&b.MaxRequests,utils.BuildFlagName(defaultAPIBreakerFlagsPrefix,"max-requests"), defaultAPIBreakerMaxRequests, "the maximum number of requests allowed to pass through when the CircuitBreaker is half-open.")
	fs.DurationVar(&b.Interval,utils.BuildFlagName(defaultAPIBreakerFlagsPrefix, "interval"), defaultAPIBreakerInterval, "the cyclic period of the closed state for the CircuitBreaker to clear the internal Counts.")
	fs.DurationVar(&b.Timeout,utils.BuildFlagName(defaultAPIBreakerFlagsPrefix, "duration"), defaultAPIBreakerTimeout, "the period of the open state, after which the state of the CircuitBreaker becomes half-open.")
}

func (b *Breaker) Parse() (err error) {
	if b.CustomParseFunc != nil {
		return b.CustomParseFunc()
	}
	return nil
}

type APIs map[string]API

func (as *APIs) Parse() (err error) {
	for name, api := range *as {
		err = api.Parse()
		if err != nil {
			return err
		}
		(*as)[name] = api
	}
	return nil
}

type Method uint8

const (
	GET Method = iota
	POST
	DELETE
	PUT
	PATCH
)

type API struct {
	Name       string            `json:"name" yaml:"name"`
	LogKVs     map[string]string `json:"log" yaml:"log"`
	Instrument []string          `json:"instrument" yaml:"instrument"`
	RateLimit  RateLimit         `json:"ratelimit" yaml:"ratelimit"`
	Breaker    Breaker           `json:"breaker" yaml:"breaker"`

	Path string `json:"path" yaml:"path"`
	Method string `json:"method" yaml:"method"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (a API) GetGoKitLoggerKVs() (kvs []interface{}) {
	for k, v := range a.LogKVs {
		kvs = append(kvs, k, v)
	}
	return
}

func (a *API) BindFlags(fs *bootflag.FlagSet)  {
	if a.CustomBindFlagsFunc != nil {
		a.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(&a.Name, utils.BuildFlagName(defaultAPIFlagsPrefix, "name"), "", "the name of the API")
	fs.StringSliceVar(&a.Instrument, utils.BuildFlagName(defaultAPIFlagsPrefix, "instrument"), []string{},"key value pair for instrument duration")

	fs.StringToStringVar(&a.LogKVs, utils.BuildFlagName(defaultAPIFlagsPrefix, "log"), defaultAPILogKVs, "key value pair for request log")

	a.RateLimit.BindFlags(fs)
	a.Breaker.BindFlags(fs)
}

func (a *API) Parse() (err error) {
	a.Name = strings.TrimSpace(a.Name)
	if a.Name == "" {
		return ErrorNoAPIName
	}
	if a.CustomParseFunc  != nil {
		return a.CustomParseFunc()
	}
	return nil
}
