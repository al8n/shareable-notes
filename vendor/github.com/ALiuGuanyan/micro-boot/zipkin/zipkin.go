// Copyright 2020 The micro-boot Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package zipkin

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
)

var (
	defaultZipkinFlagsPrefix = "zipkin"
)

func SetDefaultZipkinFlagsPrefix(prefix string)  {
	defaultZipkinFlagsPrefix = prefix
}

type Config struct {
	// Bridge uses Zipkin OpenTracing bridge instead of native implementation
	Bridge bool `json:"bridge" yaml:"bridge"`

	// Tracer is the settings for Golang official Zipkin client tracer implementation.
	Tracer Tracer `json:"tracer" yaml:"tracer"`

	Reporter Reporter `json:"reporter" yaml:"reporter"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}


func (c *Config) BindFlags (fs *bootflag.FlagSet)  {
	if c.CustomBindFlagsFunc != nil {
		c.CustomBindFlagsFunc(fs)
		return
	}

	fs.BoolVar(&c.Bridge, utils.BuildFlagName(defaultZipkinFlagsPrefix,"bridge"), true, "enables Zipkin tracing via HTTP reporter URL (default is false)")
	c.Tracer.BindFlags(fs)
	c.Reporter.BindFlags(fs)
}

func (c *Config) Parse() (err error) {
	if c.CustomParseFunc != nil {
		return c.CustomParseFunc()
	}
	return nil
}

// WithIDGenerator allows one to set a custom ID Generator
func (c Config) WithIDGenerator() zipkin.TracerOption {
	return c.Tracer.WithIDGenerator()
}

// WithLocalEndpoint sets the local endpoint of the tracer
// or returns a error if fail to create an Endpoint instance
func (c Config) WithLocalEndpoint() (zipkin.TracerOption, error) {
	return c.Tracer.WithLocalEndpoint()
}

// WithExtractFailurePolicy allows one to set the ExtractFailurePolicy.
func (c Config) WithExtractFailurePolicy() zipkin.TracerOption {
	return c.Tracer.WithExtractFailurePolicy()
}

// WithAlwaysSampler is the same as zipkin.WithSampler(zipkin.AlwaysSampler) in official Golang client.
func (c Config) WithAlwaysSampler() zipkin.TracerOption {

	return c.Tracer.WithAlwaysSampler()
}

// WithNeverSampler is the same as zipkin.WithSampler(zipkin.NeverSample) in official Golang client.
func (c Config) WithNeverSampler() zipkin.TracerOption {
	return c.Tracer.WithNeverSampler()
}

// WithBoundarySampler is the same as zipkin.WithSampler(zipkin.NewBoundarySampler(rate, salt)) in official Golang client.
func (c Config) WithBoundarySampler() (zipkin.TracerOption, error) {
	return c.Tracer.WithBoundarySampler()
}

// WithModuloSampler is the same as zipkin.WithSampler(zipkin.NewModuloSampler(mod)) in official Golang client.
func (c Config) WithModuloSampler() (zipkin.TracerOption, error) {
	return c.Tracer.WithModuloSampler()
}

// WithCountingSampler is the same as zipkin.WithSampler(zipkin.NewModuloSampler(mod)) in official Golang client.
func (c Config) WithCountingSampler() (zipkin.TracerOption, error ){
	return c.Tracer.WithCountingSampler()
}

// WithSharedSpans allows to place client-side and server-side annotations
// for a RPC call in the same span (Zipkin V1 behavior) or different spans
// (more in line with other tracing solutions). By default this Tracer
// uses shared host spans (so client-side and server-side in the same span).
func (c Config) WithSharedSpans() zipkin.TracerOption {
	return c.Tracer.WithSharedSpans()
}

// WithNoopSpan if set to true will switch to a NoopSpan implementation
// if the trace is not sampled.
func (c Config) WithNoopSpan() zipkin.TracerOption {
	return c.Tracer.WithNoopSpan()
}

func (c Config) StandardReporter() reporter.Reporter {
	return c.Reporter.Standardize()
}
