package flag

// GetFloat64Slice is the reexport of pflag.FlagSet.GetFloat64Slice()
//
// GetFloat64Slice return the []float64 value of a flag with the given name
func (f *FlagSet) GetFloat64Slice(name string) ([]float64, error) {
	return f.fs.GetFloat64Slice(name)
}

// Float64SliceVar is the reexport of pflag.FlagSet.Float64SliceVar()
//
// Float64SliceVar defines a float64Slice flag with specified name, default value, and usage string.
// The argument p points to a []float64 variable in which to store the value of the flag.
func (f *FlagSet) Float64SliceVar(p *[]float64, name string, value []float64, usage string) {
	f.fs.Float64SliceVar(p, name, value, usage)
}

// Float64SliceVarP is the reexport of pflag.FlagSet.Float64SliceVarP()
//
// Float64SliceVarP is like Float64SliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64SliceVarP(p *[]float64, name, shorthand string, value []float64, usage string) {
	f.fs.Float64SliceVarP(p, name, shorthand, value, usage)
}

// Float64Slice is the reexport of pflag.FlagSet.Float64Slice()
//
// Float64Slice defines a []float64 flag with specified name, default value, and usage string.
// The return value is the address of a []float64 variable that stores the value of the flag.
func (f *FlagSet) Float64Slice(name string, value []float64, usage string) *[]float64 {
	p := []float64{}
	f.Float64SliceVarP(&p, name, "", value, usage)
	return &p
}

// Float64SliceP is the reexport of pflag.FlagSet.Float64SliceP()
//
// Float64SliceP is like Float64Slice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64SliceP(name, shorthand string, value []float64, usage string) *[]float64 {
	p := []float64{}
	f.Float64SliceVarP(&p, name, shorthand, value, usage)
	return &p
}

