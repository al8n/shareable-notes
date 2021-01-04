package prometheus

import (
	"errors"
	"fmt"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

var (
	// Handler Opts
	defaultHandlerErrorHandling = promhttp.HTTPErrorOnError
	defaultHandlerDisableCompression = false
	defaultHandlerMaxRequestsInFlight = 0
	defaultHandlerTimeout time.Duration = 0

	// OptionsDefault
	defaultOptionsConstLabels = map[string]string{}
	defaultOptionsLabelNames = []string{}

	// Counter
	defaultCounterConstLabels = map[string]string{}
	defaultCounterLabelNames = []string{}

	// Gauge
	defaultGaugeConstLabels = map[string]string{}
	defaultGaugeLabelNames = []string{}

	// Summary
	defaultSummaryConstLabels = map[string]string{}
	defaultSummaryLabelNames = []string{}
	defaultSummaryObjectives = map[float64]float64{}
	defaultSummaryMaxAge = 10 * time.Minute
	defaultSummaryBufCap uint32 = 500
	defaultSummaryAgeBuckets uint32 = 5

	// Histogram
	defaultHistogramConstLabels = map[string]string{}
	defaultHistogramLabelNames = []string{}
	defaultHistogramBuckets = []float64{}
)

func SetDefaultHandlerErrorHandling(n promhttp.HandlerErrorHandling)  {
	defaultHandlerErrorHandling = n
}

func SetDefaultHandlerDisableCompression(n bool)  {
	defaultHandlerDisableCompression = n
}

func SetDefaultHandlerMaxRequestsInFlight(n int)  {
	defaultHandlerMaxRequestsInFlight = n
}

func SetDefaultHandlerTimeout(d time.Duration)  {
	defaultHandlerTimeout = d
}

func SetDefaultOptionsLabelNames(labels []string)  {
	defaultOptionsLabelNames = labels
}

func SetDefaultOptionsConstLabels(labels map[string]string)  {
	defaultOptionsConstLabels = labels
}

func SetDefaultCounterLabelNames(labels []string)  {
	defaultCounterLabelNames = labels
}

func SetDefaultCounterConstLabels(labels map[string]string)  {
	defaultCounterConstLabels = labels
}

func SetDefaultGaugeLabelNames(labels []string)  {
	defaultGaugeLabelNames = labels
}

func SetDefaultGaugeConstLabels(labels map[string]string)  {
	defaultGaugeConstLabels = labels
}

func SetDefaultSummaryObjectives(objectives map[float64]float64)  {
	defaultSummaryObjectives = objectives
}

func SetDefaultSummaryLabelNames(labels []string)  {
	defaultSummaryLabelNames = labels
}

func SetDefaultSummaryConstLabels(labels map[string]string)  {
	defaultSummaryConstLabels = labels
}

func SetDefaultSummaryMaxAge(d time.Duration)  {
	defaultSummaryMaxAge = d
}

func SetDefaultSummaryBufCap(num uint32)  {
	defaultSummaryBufCap = num
}

func SetDefaultSummaryAgeBuckets(num uint32)  {
	defaultSummaryAgeBuckets = num
}

func SetDefaultHistogramBuckets(buckets []float64)  {
	defaultHistogramBuckets = buckets
}

func SetDefaultHistogramLabelNames(labels []string)  {
	defaultHistogramLabelNames = labels
}

func SetDefaultHistogramConstLabels(labels map[string]string)  {
	defaultHistogramConstLabels = labels
}


// HandlerOptions is a bridge between promhttp.HandlerOpts and boot configuration
//
// See the official prometheus Golang client docs https://godoc.org/github.com/prometheus/client_golang/prometheus/promhttp#HandlerOpts for more details
type HandlerOptions struct {
	DisableCompression bool `json:"compression" yaml:"compression"`

	MaxRequestsInFlight int `json:"max-requests-in-flight" yaml:"max-requests-in-flight"`

	Timeout time.Duration `json:"timeout" yaml:"timeout"`

	ErrorHandling promhttp.HandlerErrorHandling `json:"error-handling" yaml:"error-handling"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (o *HandlerOptions) Parse() (err error) {
	if o.CustomParseFunc != nil {
		return o.CustomParseFunc()
	}
	return nil
}

// BindFlags binds the fields of HandlerOptions with command line flags
func (o *HandlerOptions) BindFlags(fs *bootflag.FlagSet) {
	if o.CustomBindFlagsFunc != nil {
		o.CustomBindFlagsFunc(fs)
		return
	}

	fs.BoolVar(&o.DisableCompression, utils.BuildFlagName(defaultPrometheusFlagsPrefix,  "disable-compression"), defaultHandlerDisableCompression, "compress prometheus response or not (default is false)")

	fs.IntVar(&o.MaxRequestsInFlight, utils.BuildFlagName(defaultPrometheusFlagsPrefix, "max-requests-inflight"), defaultHandlerMaxRequestsInFlight, "the number of concurrent HTTP prometheus requests limitation")

	fs.DurationVar(&o.Timeout, utils.BuildFlagName(defaultPrometheusFlagsPrefix, "http-timeout"), defaultHandlerTimeout, "prometheus will handle a request takes longer than timeout")

	fs.PrometheusHandlerErrorHandlingVar(&o.ErrorHandling, utils.BuildFlagName(defaultPrometheusFlagsPrefix, "error-handling"), defaultHandlerErrorHandling, "defines how errors are handled by prometheus")
}

// Standardize returns a standard promhttp.HandlerOpts.
func (o HandlerOptions) Standardize() (opts promhttp.HandlerOpts) {
	return promhttp.HandlerOpts{
		MaxRequestsInFlight: o.MaxRequestsInFlight,
		DisableCompression: o.DisableCompression,
		Timeout: o.Timeout,
		ErrorHandling: o.ErrorHandling,
	}
}

// Options is a bridge between prometheus.Opts and boot configuration
//
// See the official prometheus Golang client docs https://godoc.org/github.com/prometheus/client_golang/prometheus#Opts for more details
type Options struct {
	Namespace string 	`json:"namespace" yaml:"namespace"`
	Subsystem string 	`json:"subsystem" yaml:"subsystem"`
	Name      string	`json:"name" yaml:"name"`
	Help string			`json:"help" yaml:"help"`
	ConstLabels map[string]string `json:"labels" yaml:"labels"`
	LabelNames []string `json:"label-names" yaml:"label-names"`
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

// Parse will check if Options is a valid configuration for prometheus.Opts
func (o *Options) Parse() (err error) {
	if err = parse(o.Name, "Options"); err != nil {
		return err
	}
	if o.CustomParseFunc != nil {
		return o.CustomParseFunc()
	}
	return nil
}

// BindFlags binds the fields of Options with command line flags
func (o *Options) BindFlags(fs *bootflag.FlagSet)  {
	if o.CustomBindFlagsFunc != nil {
		o.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(&o.Name,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "opts-name"), "", "the name of prometheus options")

	fs.StringVar(&o.Subsystem,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"opts-subsystem"), "", "the subsystem of prometheus options")

	fs.StringVar(&o.Namespace, utils.BuildFlagName(defaultPrometheusFlagsPrefix,"opts-namespace"), "", "the namespace of prometheus options")

	fs.StringVar(&o.Help,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"opts-help"), "", "the help of prometheus options")

	fs.StringToStringVar(&o.ConstLabels,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"opts-const-labels"), defaultOptionsConstLabels,"the const labels of prometheus options")

	fs.StringSliceVar(&o.LabelNames,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"opts-labels"), defaultOptionsLabelNames, "the labels of prometheus options")
}

// Standardize returns a standard prometheus.Opts
func (o *Options) Standardize() prometheus.Opts {
	return prometheus.Opts{
		Namespace:   o.Namespace,
		Subsystem:   o.Subsystem,
		Name:        o.Name,
		Help:        o.Help,
		ConstLabels: o.ConstLabels,
	}
}

// CounterOptions is a bridge between prometheus.CounterOpts and boot configuration
//
// See the official prometheus Golang client docs
// https://godoc.org/github.com/prometheus/client_golang/prometheus#CounterOpts for more details
type CounterOptions Options

// Parse will check if CounterOptions is a valid configuration for prometheus.CounterOpts
func (o *CounterOptions) Parse() (err error) {
	if err = parse(o.Name, "CounterOptions"); err != nil {
		return err
	}
	if o.CustomParseFunc != nil {
		return o.CustomParseFunc()
	}
	return nil
}

// Standardize returns a standard prometheus.CounterOpts
func (o *CounterOptions) Standardize() prometheus.CounterOpts {
	return prometheus.CounterOpts{
		Namespace:   o.Namespace,
		Subsystem:   o.Subsystem,
		Name:        o.Name,
		Help:        o.Help,
		ConstLabels: o.ConstLabels,
	}
}

// BindFlags binds the fields of CounterOptions with command line flags
func (o *CounterOptions) BindFlags(fs *bootflag.FlagSet)  {
	if o.CustomBindFlagsFunc != nil {
		o.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(&o.Name,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "ctr-name"), "", "the name of prometheus counter options")

	fs.StringVar(&o.Subsystem,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"ctr-subsystem"), "", "the subsystem of prometheus counter options")

	fs.StringVar(&o.Namespace, utils.BuildFlagName(defaultPrometheusFlagsPrefix,"ctr-namespace"), "", "the namespace of prometheus counter options")

	fs.StringVar(&o.Help,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"ctr-help"), "", "the help of prometheus  counter options")

	fs.StringToStringVar(&o.ConstLabels,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"ctr-const-labels"), defaultCounterConstLabels,"the const labels of prometheus counter options")

	fs.StringSliceVar(&o.LabelNames,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"ctr-labels"), defaultCounterLabelNames, "the labels of prometheus counter options")
}

// GaugeOptions is a bridge between prometheus.CounterOpts and boot configuration
//
// See the official prometheus Golang client docs
// https://godoc.org/github.com/prometheus/client_golang/prometheus#GaugeOpts for more details
type GaugeOptions Options

// Parse will check if GaugeOptions is a valid configuration for prometheus.GaugeOpts
func (o *GaugeOptions) Parse() (err error) {
	if err = parse(o.Name, "GaugeOptions"); err != nil {
		return err
	}
	if o.CustomParseFunc != nil {
		return o.CustomParseFunc()
	}
	return nil
}

// Standardize returns a standard prometheus.GaugeOpts
func (o *GaugeOptions) Standardize() prometheus.GaugeOpts {
	return prometheus.GaugeOpts{
		Namespace:   o.Namespace,
		Subsystem:   o.Subsystem,
		Name:        o.Name,
		Help:        o.Help,
		ConstLabels: o.ConstLabels,
	}
}

// BindFlags binds the fields of GaugeOptions with command line flags
func (o *GaugeOptions) BindFlags(fs *bootflag.FlagSet)  {
	if o.CustomBindFlagsFunc != nil {
		o.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(&o.Name,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "gauge-name"), "", "the name of prometheus gauge options")

	fs.StringVar(&o.Subsystem,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"gauge-subsystem"), "", "the subsystem of prometheus gauge options")

	fs.StringVar(&o.Namespace, utils.BuildFlagName(defaultPrometheusFlagsPrefix,"gauge-namespace"), "", "the namespace of prometheus gauge options")

	fs.StringVar(&o.Help,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"gauge-help"), "", "the help of prometheus  gauge options")

	fs.StringToStringVar(&o.ConstLabels,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"gauge-const-labels"), defaultGaugeConstLabels,"the const labels of prometheus gauge options")

	fs.StringSliceVar(&o.LabelNames,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"gauge-labels"), defaultGaugeLabelNames, "the labels of prometheus gauge options")
}

// SummaryOptions is a bridge between prometheus.SummaryOpts and boot configuration
//
// See the official prometheus Golang client docs
// https://godoc.org/github.com/prometheus/client_golang/prometheus#SummaryOpts for more details
type SummaryOptions struct {
	Namespace string 	`json:"namespace" yaml:"namespace"`
	Subsystem string 	`json:"subsystem" yaml:"subsystem"`
	Name      string	`json:"name" yaml:"name"`
	Help 	  string	`json:"help" yaml:"help"`
	ConstLabels map[string]string `json:"labels" yaml:"labels"`
	LabelNames []string `json:"label-names" yaml:"label-names"`
	Objectives map[float64]float64 `json:"objectives" yaml:"objectives"`
	MaxAge time.Duration `json:"max-age" yaml:"max-age"`
	AgeBuckets uint32 `json:"age-buckets" yaml:"age-buckets"`
	BufCap uint32 `json:"buf-cap" yaml:"buf-cap"`
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

// Parse will check if SummaryOptions is a valid configuration for prometheus.SummaryOpts
func (o *SummaryOptions) Parse() (err error) {
	if err = parse(o.Name, "SummaryOptions"); err != nil {
		return err
	}
	if o.CustomParseFunc != nil {
		return o.CustomParseFunc()
	}
	return nil
}

// BindFlags binds the fields of SummaryOptions with command line flags
func (o *SummaryOptions) BindFlags(fs *bootflag.FlagSet)  {
	if o.CustomBindFlagsFunc != nil {
		o.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(&o.Name,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "summary-name"), "", "the name of prometheus summary options")

	fs.StringVar(&o.Subsystem,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "summary-subsystem"), "", "the subsystem of prometheus summary options")

	fs.StringVar(&o.Namespace,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "summary-namespace"), "", "the namespace of prometheus summary options")

	fs.StringVar(&o.Help,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "summary-help"), "", "the help of prometheus summary options")

	fs.StringToStringVar(&o.ConstLabels,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "summary-const-labels"), defaultSummaryConstLabels, "the const labels of prometheus summary options")

	fs.StringSliceVar(&o.LabelNames,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "summary-labels"), defaultSummaryLabelNames, "the labels of prometheus summary options")

	fs.Float64ToFloat64Var( &o.Objectives,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"summary-objectives"), defaultSummaryObjectives, "the objectives of prometheus summary options")

	fs.DurationVar(&o.MaxAge,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"summary-maxage"), defaultSummaryMaxAge, "the max age of prometheus summary options")

	fs.Uint32Var(&o.AgeBuckets,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"summary-age-buckets"), defaultSummaryAgeBuckets, "the number of age buckets of prometheus summary options")

	fs.Uint32Var(&o.BufCap,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"summary-age-buckets"), defaultSummaryBufCap, "the bufcap of prometheus summary options")
}


// Standardize returns a standard prometheus.SummaryOpts
func (o *SummaryOptions) Standardize() prometheus.SummaryOpts {
	return prometheus.SummaryOpts{
		Namespace:   o.Namespace,
		Subsystem:   o.Subsystem,
		Name:        o.Name,
		Help:        o.Help,
		ConstLabels: o.ConstLabels,
		Objectives: o.Objectives,
		MaxAge: o.MaxAge,
		AgeBuckets: o.AgeBuckets,
		BufCap: o.BufCap,
	}
}

// HistogramOptions is a bridge between prometheus.HistogramOpts and boot configuration
//
// See the official prometheus Golang client docs
// https://godoc.org/github.com/prometheus/client_golang/prometheus#HistogramOpts for more details
type HistogramOptions struct {
	Namespace string 	`json:"namespace" yaml:"namespace"`
	Subsystem string 	`json:"subsystem" yaml:"subsystem"`
	Name      string	`json:"name" yaml:"name"`
	Help string			`json:"help" yaml:"help"`
	ConstLabels map[string]string `json:"labels" yaml:"labels"`
	LabelNames []string `json:"label-names" yaml:"label-names"`
	Buckets []float64 `json:"buckets" yaml:"buckets"`
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

// Parse will check if HistogramOptions is a valid configuration for prometheus.HistogramOpts
func (o *HistogramOptions) Parse() (err error) {
	if err = parse(o.Name, "HistogramOptions"); err != nil {
		return err
	}
	if o.CustomParseFunc != nil {
		return o.CustomParseFunc()
	}
	return nil
}

// BindFlags binds the fields of HistogramOptions with command line flags
func (o *HistogramOptions) BindFlags(fs *bootflag.FlagSet)  {
	if o.CustomBindFlagsFunc != nil {
		o.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(&o.Name,utils.BuildFlagName(defaultPrometheusFlagsPrefix, "prom-histo-name"), "", "the name of prometheus histogram options")

	fs.StringVar(&o.Subsystem,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"prom-histo-subsystem"), "", "the subsystem of prometheus histogram options")

	fs.StringVar(&o.Namespace,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"prom-histo-namespace"), "", "the namespace of prometheus histogram options")

	fs.StringVar(&o.Help,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"prom-histo-help"), "", "the help of prometheus histogram options")

	fs.StringToStringVar(&o.ConstLabels,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"prom-histo-const-labels"), defaultHistogramConstLabels, "the const labels of prometheus histogram options")

	fs.StringSliceVar(&o.LabelNames,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"prom-histo-labels"), defaultHistogramLabelNames, "the labels of prometheus histogram options")

	fs.Float64SliceVar(&o.Buckets,utils.BuildFlagName(defaultPrometheusFlagsPrefix,"prom-histo-labels"), defaultHistogramBuckets, "the labels of prometheus histogram options")
}

// Standardize returns a standard prometheus.HistogramOpts
func (o *HistogramOptions) Standardize() prometheus.HistogramOpts {
	return prometheus.HistogramOpts{
		Namespace:   o.Namespace,
		Subsystem:   o.Subsystem,
		Name:        o.Name,
		Help:        o.Help,
		ConstLabels: o.ConstLabels,
		Buckets: o.Buckets,
	}
}

func parse(name, typ string) (err error)  {
	if name == "" {
		return errors.New(fmt.Sprintf("name of %s is mandatory.", typ))
	}
	return nil
}
