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
package prometheus

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
)

var (
	// prefix
	defaultPrometheusFlagsPrefix = "prom"

	defaultPrometheusPath = "/metrics"
)

func SetDefaultPrometheusFlagsPrefix(val string)  {
	defaultPrometheusFlagsPrefix = val
}

func SetDefaultPrometheusPath(val string)  {
	defaultPrometheusPath = val
}

// Config to boot prometheus
type Config struct {
	Path           string         `json:"path" yaml:"path"`
	HandlerOptions HandlerOptions `json:"handler-options" yaml:"handler-options"`

	CounterOptions map[string]CounterOptions     `json:"counter-options" yaml:"counter-options"`
	GaugeOptions map[string]GaugeOptions         `json:"gauge-options" yaml:"gauge-options"`
	SummaryOptions map[string]SummaryOptions     `json:"summary-options" yaml:"summary-options"`
	HistogramOptions map[string]HistogramOptions `json:"histogram-options" yaml:"histogram-options"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (c *Config) BindFlags(fs *bootflag.FlagSet)  {
	if c.CustomBindFlagsFunc != nil {
		c.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&c.Path,
		utils.BuildFlagName(defaultPrometheusFlagsPrefix, "path"),
		defaultPrometheusPath,
		"specify the metrics http route")

	c.HandlerOptions.BindFlags(fs)
}

func (c *Config) Parse() (err error) {
	if c.CustomParseFunc != nil {
		return c.CustomParseFunc()
	}

	return nil
}
