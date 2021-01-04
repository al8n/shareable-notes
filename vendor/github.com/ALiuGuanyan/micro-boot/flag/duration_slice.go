package flag

import "time"

// GetDurationSlice is the reexport of pflag.FlagSet.GetDurationSlice()
//
// GetDurationSlice returns the []time.Duration value of a flag with the given name
func (f *FlagSet) GetDurationSlice(name string) ([]time.Duration, error) {
	return f.fs.GetDurationSlice(name)
}

// DurationSliceVar is the reexport of pflag.FlagSet.DurationSliceVar()
//
// DurationSliceVar defines a durationSlice flag with specified name, default value, and usage string.
// The argument p points to a []time.Duration variable in which to store the value of the flag.
func (f *FlagSet) DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string) {
	f.fs.DurationSliceVar(p, name, value, usage)
}

// DurationSliceVarP is the reexport of pflag.FlagSet.DurationSliceVarP()
//
// DurationSliceVarP is like DurationSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) DurationSliceVarP(p *[]time.Duration, name, shorthand string, value []time.Duration, usage string) {
	f.fs.DurationSliceVarP(p, name, shorthand, value, usage)
}

// DurationSlice is the reexport of pflag.FlagSet.DurationSlice()
//
// DurationSlice defines a []time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a []time.Duration variable that stores the value of the flag.
func (f *FlagSet) DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration {
	p := []time.Duration{}
	f.DurationSliceVarP(&p, name, "", value, usage)
	return &p
}

// DurationSliceP is the reexport of pflag.FlagSet.DurationSliceP()
//
// DurationSliceP is like DurationSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) DurationSliceP(name, shorthand string, value []time.Duration, usage string) *[]time.Duration {
	p := []time.Duration{}
	f.DurationSliceVarP(&p, name, shorthand, value, usage)
	return &p
}
