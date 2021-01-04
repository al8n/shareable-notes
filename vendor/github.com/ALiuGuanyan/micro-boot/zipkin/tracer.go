package zipkin

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/idgenerator"
)

var (
	defaultTracerFlagsPrefix = defaultZipkinFlagsPrefix + "-tracer"

	defaultAlwaysSample = true
	defaultNeverSample = false

	defaultNoopSpan = false
	defaultSharedSpans = true

	defaultTracerGenerator = bootflag.Random64
	defaultTracerTags = map[string]string{}
	defaultTracerNoop = false
	defaultCountingSamplerRate = 0.0
	defaultBoundarySamplerRate = 0.0
	defaultBoundarySamplerSalt int64 = 0
	defaultModuloSamplerMod uint64 = 0
)

func SetDefaultTracerFlagsPrefix(prefix string)  {
	defaultTracerFlagsPrefix = prefix
}

func SetDefaultTracerAlwaysSample(b bool) {
	defaultAlwaysSample = b
}

func SetDefaultTracerNeverSample(b bool) {
	defaultNeverSample = b
}

func SetDefaultTracerNoopSpan(b bool)  {
	defaultNoopSpan = b
}

func SetDefaultTracerSharedSpans(b bool)  {
	defaultSharedSpans = b
}

func SetDefaultTracerTags(tags map[string]string)  {
	defaultTracerTags = tags
}

func SetDefaultTracerGenerator(generator bootflag.ZipkinIDGenerator)  {
	defaultTracerGenerator = generator
}

func SetDefaultTracerNoop(b bool)  {
	defaultTracerNoop = b
}

func SetDefaultCountingSamplerRate(rate float64)  {
	defaultCountingSamplerRate = rate
}

func SetDefaultBoundarySamplerRate(rate float64)  {
	defaultBoundarySamplerRate = rate
}

func SetDefaultBoundarySamplerSalt(salt int64)  {
	defaultBoundarySamplerSalt = salt
}

func SetDefaultModuloSamplerMod(mod uint64)  {
	defaultModuloSamplerMod = mod
}

type Span struct {
	Noop bool `json:"noop-span" yaml:"noop-span"`
	SharedSpans bool `json:"shared-spans" yaml:"shared-spans"`
}

// Sampler specifies a sample policy if a Zipkin span should be sampled, based on its traceID.
type Sampler struct {
	// AlwaysSample will always return true. If used by a service it will always start
	// traces if no upstream trace has been propagated. If an incoming upstream trace
	// is not sampled the service will adhere to this and only propagate the context.
	AlwaysSample bool `json:"always" yaml:"always"`

	// NeverSample will always return false. If used by a service it will not allow
	// the service to start traces but will still allow the service to participate
	// in traces started upstream.
	NeverSample bool `json:"never" yaml:"never"`

	// BoundarySampler is the settings of boundary sampler for
	// creating a standard zipkin boundary sampler
	BoundarySampler BoundarySampler `json:"boundary" yaml:"boundary"`

	// ModuloSampler is the settings of modulo sampler for creating a standard zipkin
	// module sampler
	ModuloSampler ModuloSampler `json:"modulo" yaml:"modulo"`

	// CountingSampler is the settings of counting sampler for creating a standard zipkin
	// counting sampler
	CountingSampler CountingSampler `json:"counting" yaml:"counting"`
}

// Tracer is the settings for Golang official Zipkin client tracer implementation.
type Tracer struct {
	// ExtractFailurePolicy deals with Extraction errors:
	//
	// 0 for ExtractFailurePolicyRestart
	//
	// 1 for ExtractFailurePolicyError
	//
	// 2 for ExtractFailurePolicyTagAndRestart
	ExtractFailurePolicy zipkin.ExtractFailurePolicy `json:"extract-failure-policy" yaml:"extract-failure-policy"`

	// LocalEndpoint is the settings for creating a standard Zipkin local endpoint
	LocalEndpoint LocalEndpoint `json:"endpoint" yaml:"endpoint"`

	// Generator interface can be used to provide the Zipkin Tracer with custom
	// implementations to generate Span and Trace IDs.
	Generator idgenerator.IDGenerator `json:"generator" yaml:"generator"`

	// Tags allows one to set default tags to be added to each created span
	Tags map[string]string `json:"tags" yaml:"tags"`

	// Noop allows one to start the zipkin tracer as noop implementation.
	Noop bool `json:"noop" yaml:"noop"`

	// Sampler is the sample policy for a zipkin tracer
	Sampler Sampler `json:"sampler" yaml:"sampler"`

	Span Span `json:"span" yaml:"span"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (t *Tracer) Parse() (err error) {
	if t.CustomParseFunc != nil {
		return t.CustomParseFunc()
	}
	return nil
}

func (t *Tracer) BindFlags (fs *bootflag.FlagSet)  {
	if t.CustomBindFlagsFunc != nil {
		t.CustomBindFlagsFunc(fs)
		return
	}

	fs.ZipkinExtractFailurePolicyVar(
		&t.ExtractFailurePolicy,
		utils.BuildFlagName(defaultTracerFlagsPrefix, "extract-failure-policy"),
		defaultTracerExtractFailurePolicy,
		"deals with Extraction errors (restart, error, or tagAndRestart)")

	fs.StringVar(
		&t.LocalEndpoint.Name,
		utils.BuildFlagName(defaultTracerFlagsPrefix,"endpoint-name"),
		"",
		"the service name which used to create a local endpoint")

	fs.StringVar(
		&t.LocalEndpoint.HostPort,
		utils.BuildFlagName(defaultTracerFlagsPrefix,"endpoint-address"),
		"",
		"the address which used to create a local endpoint. eg. (localhost:80)")

	fs.ZipkinIDGeneratorVar( &t.Generator, utils.BuildFlagName(defaultTracerFlagsPrefix,"generator"), defaultTracerGenerator, "set a custom ID Generator")

	fs.StringToStringVar(
		&t.Tags,
		utils.BuildFlagName(defaultTracerFlagsPrefix, "tags"), defaultTracerTags,
		"tags allows one to set default tags to be added to each created span")

	fs.BoolVar(
		&t.Noop,
		utils.BuildFlagName(defaultTracerFlagsPrefix, "noop"), defaultTracerNoop,
		"allows one to start the Tracer as Noop implementation.")

	fs.BoolVar(
		&t.Span.Noop,
		utils.BuildFlagName(defaultTracerFlagsPrefix, "noop-span"),
		defaultNoopSpan,
		"if true tracer will switch to a NoopSpan implementation if the trace is not sampled.")

	fs.BoolVar(
		&t.Span.SharedSpans,
		utils.BuildFlagName(defaultTracerFlagsPrefix, "shared-spans"),
		defaultSharedSpans,
		"allows to place client-side and server-side annotations for a RPC call")

	fs.BoolVar(
		&t.Sampler.AlwaysSample,
		utils.BuildFlagName(defaultTracerFlagsPrefix, "always-sample"),
		defaultAlwaysSample,
		"zipkin tracer will always sample")

	fs.BoolVar(
		&t.Sampler.NeverSample,
		utils.BuildFlagName(defaultTracerFlagsPrefix, "never-sample"),
		defaultNeverSample,
		"zipkin tracer will never sample")

	fs.Float64Var(
		&t.Sampler.CountingSampler.Rate,
		utils.BuildFlagName(defaultTracerFlagsPrefix,"counting-rate"),
		defaultCountingSamplerRate,
		"the rate for zipkin counting sampler (default is 0.0, never sample)")

	fs.Float64Var(
		&t.Sampler.BoundarySampler.Rate,
		utils.BuildFlagName(defaultTracerFlagsPrefix,"boundary-rate"),
		defaultBoundarySamplerRate,
		"the rate for zipkin boundary sampler (default is 0.0, never sample)")

	fs.Int64Var(
		&t.Sampler.BoundarySampler.Salt,
		utils.BuildFlagName(defaultTracerFlagsPrefix,"boundary-salt"),
		defaultBoundarySamplerSalt,
		"the salt for zipkin boundary sampler (default is 0)")

	fs.Uint64Var(
		&t.Sampler.ModuloSampler.Mod,
		utils.BuildFlagName(defaultTracerFlagsPrefix,"modulo-mod"),
		defaultModuloSamplerMod,
		"the mod for zipkin modulo sampler (default is 0, always sample)")
}

// WithIDGenerator allows one to set a custom ID Generator
func (t Tracer) WithIDGenerator() zipkin.TracerOption {
	return zipkin.WithIDGenerator(t.Generator)
}

// WithLocalEndpoint sets the local endpoint of the tracer 
// or returns a error if fail to create an Endpoint instance
func (t Tracer) WithLocalEndpoint() (zipkin.TracerOption, error) {
	ep, err := zipkin.NewEndpoint(t.LocalEndpoint.Name, t.LocalEndpoint.HostPort)
	if err != nil {
		return nil, err
	}
	return zipkin.WithLocalEndpoint(ep), nil
}

// WithExtractFailurePolicy allows one to set the ExtractFailurePolicy.
func (t Tracer) WithExtractFailurePolicy() zipkin.TracerOption {
	return zipkin.WithExtractFailurePolicy(t.ExtractFailurePolicy)
}

// WithAlwaysSampler is the same as zipkin.WithSampler(zipkin.AlwaysSampler) in official Golang client.
func (t Tracer) WithAlwaysSampler() zipkin.TracerOption {
	if t.Sampler.AlwaysSample {
		return zipkin.WithSampler(zipkin.AlwaysSample)
	}
	return nil
}

// WithNeverSampler is the same as zipkin.WithSampler(zipkin.NeverSample) in official Golang client.
func (t Tracer) WithNeverSampler() zipkin.TracerOption {
	if t.Sampler.NeverSample {
		return zipkin.WithSampler(zipkin.NeverSample)
	}
	return nil
}

// WithBoundarySampler is the same as zipkin.WithSampler(zipkin.NewBoundarySampler(rate, salt)) in official Golang client.
func (t Tracer) WithBoundarySampler() (zipkin.TracerOption, error) {
	//if t.Sampler.BoundarySampler == nil {
	//	return nil, errors.New("no enough parameters provided for boundary sampler")
	//}

	s, err := zipkin.NewBoundarySampler(t.Sampler.BoundarySampler.Rate, t.Sampler.BoundarySampler.Salt)
	if err != nil {
		return nil, err
	}
	return zipkin.WithSampler(s), nil
}

// WithModuloSampler is the same as zipkin.WithSampler(zipkin.NewModuloSampler(mod)) in official Golang client.
func (t Tracer) WithModuloSampler() (zipkin.TracerOption, error) {
	//if t.Sampler.ModuloSampler == nil {
	//	return nil, errors.New("no enough parameters provided for modulo sampler")
	//}
	return zipkin.WithSampler(zipkin.NewModuloSampler(t.Sampler.ModuloSampler.Mod)), nil
}

// WithCountingSampler is the same as zipkin.WithSampler(zipkin.NewModuloSampler(mod)) in official Golang client.
func (t Tracer) WithCountingSampler() (zipkin.TracerOption, error ){
	//if t.Sampler.CountingSampler == nil {
	//	return nil, errors.New("no enough parameters provided for counting sampler")
	//}

	s, err := zipkin.NewCountingSampler(t.Sampler.CountingSampler.Rate)
	if err != nil {
		return nil, err
	}
	return zipkin.WithSampler(s), nil
}

// WithSharedSpans allows to place client-side and server-side annotations
// for a RPC call in the same span (Zipkin V1 behavior) or different spans
// (more in line with other tracing solutions). By default this Tracer
// uses shared host spans (so client-side and server-side in the same span).
func (t Tracer) WithSharedSpans() zipkin.TracerOption {
	return zipkin.WithSharedSpans(t.Span.SharedSpans)
}

// WithNoopSpan if set to true will switch to a NoopSpan implementation
// if the trace is not sampled.
func (t Tracer) WithNoopSpan() zipkin.TracerOption {
	return zipkin.WithNoopSpan(t.Span.Noop)
}