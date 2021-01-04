package flag

import "time"

// GetDuration is the reexport of pflag.FlagSet.GetDuration()
//
// GetDuration return the duration value of a flag with the given name
func (f *FlagSet) GetDuration(name string) (time.Duration, error) {
	return f.fs.GetDuration(name)
}

// DurationVar is the reexport of pflag.FlagSet.DurationVar()
//
// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.fs.DurationVar(p, name, value, usage)
}

// DurationVarP is the reexport of pflag.FlagSet.DurationVarP()
//
// DurationVarP is like DurationVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) DurationVarP(p *time.Duration, name, shorthand string, value time.Duration, usage string) {
	f.fs.DurationVarP(p, name, shorthand,value,usage)
}

// Duration is the reexport of pflag.FlagSet.Duration()
//
// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVarP(p, name, "", value, usage)
	return p
}

// DurationP is the reexport of pflag.FlagSet.DurationP()
//
// DurationP is like Duration, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) DurationP(name, shorthand string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVarP(p, name, shorthand, value, usage)
	return p
}