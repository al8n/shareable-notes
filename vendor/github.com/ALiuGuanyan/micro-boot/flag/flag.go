// flag defines some special flags for micro-service tools such as Prometheus, Zipkin, OpenTracer...
//
// Most of functions and basic structure are reexported from https://github.com/spf13/pflag,
// which means you do not need to learn a new flag package and all of usages of this package are the same as pflag.
//
//
// This package does not provide any parse method for FlagSet, because
// it should be used with boot.Command
//
//
//
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
package flag

import (
	"fmt"
	"github.com/spf13/pflag"
	"strings"
)

// ErrorHandling defines how to handle flag parsing errors.
type ErrorHandling int

const (
	// ContinueOnError will return an err from Parse() if an error is found
	ContinueOnError ErrorHandling = iota
	// ExitOnError will call os.Exit(2) if an error is found when parsing
	ExitOnError
	// PanicOnError will panic() if an error is found when parsing flags
	PanicOnError
)

// Flags defines a behaviour for custom config structure.
//
// If you want to have a custom config structure, the BindFlags method must be implemented
type Flags interface {
	// BindFlags binds Structure fields with the command line flags
	BindFlags(fs *FlagSet)

	// Parse is used to check the validity of the config structure,
	// which aims to find the weird values before boot the services.
	//
	//
	// If the config structure does not need to do any validity check,
	// you can just leave it with {return nil}.
	// See examples/nocheck and examples/check.
	//
	//
	// You can also do some value assignment in this method.
	// See examples/assign
	Parse() (err error)
}

type FlagSet struct {
	fs *pflag.FlagSet
}

// NewFlagSet returns a FlagSet instance
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
	switch errorHandling {
	case ContinueOnError:
		return &FlagSet{
			fs: pflag.NewFlagSet(name, pflag.ContinueOnError),
		}
	case ExitOnError:
		return &FlagSet{
			fs: pflag.NewFlagSet(name, pflag.ExitOnError),
		}
	case PanicOnError:
		return &FlagSet{
			fs: pflag.NewFlagSet(name, pflag.PanicOnError),
		}
	default:
		return &FlagSet{
			fs: pflag.NewFlagSet(name, pflag.ExitOnError),
		}
	}
}


// GetFlagSet returns the internal plag.FlagSet of FlagSet
func (f *FlagSet) GetFlagSet() *pflag.FlagSet {
	return f.fs
}

func (f *FlagSet) getFlagType(name string, ftype string, convFunc func(sval string) (interface{}, error)) (interface{}, error) {
	flag := f.fs.Lookup(name)
	if flag == nil {
		err := fmt.Errorf("flag accessed but not defined: %s", name)
		return nil, err
	}

	if flag.Value.Type() != ftype {
		err := fmt.Errorf("trying to get %s value of flag of type %s", ftype, flag.Value.Type())
		return nil, err
	}

	sval := flag.Value.String()
	result, err := convFunc(sval)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (f *FlagSet) Args() []string {
	return f.fs.Args()
}

func (f *FlagSet) NArgs() int {
	return f.fs.NArg()
}

// Value is the reexport of pflag.Value
//
// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
type Value interface {
	String() string
	Set(string) error
	Type() string
}

// Var is the reexport of pflag.FlagSet.Var()
//
// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func (f *FlagSet) Var(value Value, name string, usage string) {
	f.fs.VarP(value, name, "", usage)
}

// VarP is the reexport of pflag.VarP
//
// VarP is like Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) VarP(value Value, name, shorthand, usage string) {
	f.fs.VarPF(value, name, shorthand, usage)
}

// AddFlagSet is the reexport of pflag.AddFlagSet
//
// AddFlagSet adds one FlagSet to another. If a flag is already present in f
// the flag from newSet will be ignored.
func (f *FlagSet) AddFlagSet(newSet *FlagSet) {
	f.fs.AddFlagSet(newSet.fs)
}



var dashBlankReplacer = strings.NewReplacer(
	"-", "",
	" ", "",
	"_", "")


