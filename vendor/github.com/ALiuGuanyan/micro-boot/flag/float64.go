package flag

// GetFloat64 is the reexport of pflag.FlagSet.GetFloat64()
//
// GetFloat64 return the float64 value of a flag with the given name
func (f *FlagSet) GetFloat64(name string) (float64, error) {
	return f.fs.GetFloat64(name)
}


// Float64Var is the reexport of pflag.FlagSet.Float64Var()
//
// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string) {
	f.fs.Float64Var(p, name, value, usage)
}

// Float64VarP is the reexport of pflag.FlagSet.Float64VarP()
//
// Float64VarP is like Float64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64VarP(p *float64, name, shorthand string, value float64, usage string) {
	f.fs.Float64VarP(p, name, shorthand, value, usage)
}


// Float64 is the reexport of pflag.FlagSet.Float64()
//
// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (f *FlagSet) Float64(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64VarP(p, name, "", value, usage)
	return p
}

// Float64P is the reexport of pflag.FlagSet.Float64P()
//
// Float64P is like Float64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64P(name, shorthand string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64VarP(p, name, shorthand, value, usage)
	return p
}
