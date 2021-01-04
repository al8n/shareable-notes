package zipkin

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"github.com/openzipkin/zipkin-go/reporter"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"time"
)

var (
	defaultReporterPrefix = defaultZipkinFlagsPrefix + "-reporter"

	defaultReporterTimeout = 5 * time.Second
	defaultReporterBatchSize = 100
	defaultReporterMaxBacklog = 1000
	defaultReporterBatchInterval = 1 * time.Second
)

func SetDefaultReporterFlagsPrefix(prefix string)  {
	defaultReporterPrefix = prefix
}

func SetDefaultReporterBatchSize(num int)  {
	defaultReporterBatchSize = num
}

func SetDefaultReporterMaxBacklog(num int)  {
	defaultReporterMaxBacklog = num
}

func SetDefaultReporterTimeout(d time.Duration)  {
	defaultReporterTimeout = d
}

func SetDefaultReporterBatchInterval(d time.Duration)  {
	defaultReporterBatchInterval = d
}

// Reporter is the settings for official Zipkin Golang client reporter instance.
type Reporter struct {
	// URL enables Zipkin tracing via HTTP reporter URL e.g. http://localhost:9411/api/v2/spans
	URL string `json:"url" yaml:"url"`

	// Timeout sets maximum timeout for the http request through its context.
	Timeout time.Duration `json:"timeout" yaml:"timeout"`

	// BatchSize sets the maximum batch size, after which a collect will be
	// triggered. The default batch size is 100 traces.
	BatchSize int `json:"batch-size" yaml:"batch-size"`

	// MaxBacklog sets the maximum backlog size. When batch size reaches this
	// threshold, spans from the beginning of the batch will be disposed.
	MaxBacklog int `json:"max-backlog" yaml:"max-backlog"`

	// BatchInterval sets the maximum duration we will buffer traces before
	// emitting them to the collector. The default batch interval is 1 second.
	BatchInterval time.Duration `json:"batch-interval" yaml:"batch-interval"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (r *Reporter) Parse() (err error) {
	if r.CustomParseFunc != nil {
		return r.CustomParseFunc()
	}
	return nil
}

func (r *Reporter) BindFlags(fs *bootflag.FlagSet)  {
	if r.CustomBindFlagsFunc != nil {
		r.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(&r.URL, utils.BuildFlagName(defaultReporterPrefix, "url"), "", "Zipkin tracing via HTTP reporter URL")
	fs.DurationVar(&r.Timeout, utils.BuildFlagName(defaultReporterPrefix, "timeout"), defaultReporterTimeout, "sets maximum timeout for the http request through its context")
	fs.IntVar(&r.BatchSize, utils.BuildFlagName(defaultReporterPrefix, "batch-size"), defaultReporterBatchSize, "Zipkin tracing via HTTP reporter URL")
	fs.IntVar(&r.MaxBacklog, utils.BuildFlagName(defaultReporterPrefix, "max-backlog"), defaultReporterMaxBacklog, "sets the maximum backlog size")
	fs.DurationVar(&r.BatchInterval, utils.BuildFlagName(defaultReporterPrefix, "batch-interval"), defaultReporterBatchInterval, "sets the maximum duration we will buffer traces before emitting them to the collector")
}

// Timeout sets maximum timeout for the http request through its context.
func(r Reporter) WithTimeout() httpreporter.ReporterOption {
	return httpreporter.Timeout(r.Timeout)
}

// BatchSize sets the maximum batch size, after which a collect will be
// triggered. The default batch size is 100 traces.
func(r Reporter) WithBatchSize() httpreporter.ReporterOption {
	return httpreporter.BatchSize(r.BatchSize)
}

// MaxBacklog sets the maximum backlog size. When batch size reaches this
// threshold, spans from the beginning of the batch will be disposed.
func(r Reporter) WithMaxBacklog() httpreporter.ReporterOption {
	return httpreporter.MaxBacklog(r.MaxBacklog)
}

// BatchInterval sets the maximum duration we will buffer traces before
// emitting them to the collector. The default batch interval is 1 second.
func(r Reporter) WithBatchInterval() httpreporter.ReporterOption {
	return httpreporter.BatchInterval(r.BatchInterval)
}

func (r Reporter) Standardize() reporter.Reporter {
	return zipkinhttp.NewReporter(r.URL, r.WithTimeout(), r.WithBatchInterval(), r.WithBatchSize(), r.WithMaxBacklog())
}