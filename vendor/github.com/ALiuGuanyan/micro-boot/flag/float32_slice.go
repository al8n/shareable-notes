package flag

// GetFloat32Slice is the reexport of pflag.FlagSet.GetFloat32Slice()
//
// GetFloat32Slice return the []float32 value of a flag with the given name
func (f *FlagSet) GetFloat32Slice(name string) ([]float32, error) {
	return f.fs.GetFloat32Slice(name)
}

// Float32SliceVar is the reexport of pflag.FlagSet.Float32SliceVar()
//
// Float32SliceVar defines a float32Slice flag with specified name, default value, and usage string.
// The argument p points to a []float32 variable in which to store the value of the flag.
func (f *FlagSet) Float32SliceVar(p *[]float32, name string, value []float32, usage string) {
	f.fs.Float32SliceVar(p, name, value, usage)
}

// Float32SliceVarP is the reexport of pflag.FlagSet.Float32SliceVarP()
//
// Float32SliceVarP is like Float32SliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float32SliceVarP(p *[]float32, name, shorthand string, value []float32, usage string) {
	f.fs.Float32SliceVarP(p, name,shorthand,value,usage)
}

// Float32Slice is the reexport of pflag.FlagSet.Float32Slice()
//
// Float32Slice defines a []float32 flag with specified name, default value, and usage string.
// The return value is the address of a []float32 variable that stores the value of the flag.
func (f *FlagSet) Float32Slice(name string, value []float32, usage string) *[]float32 {
	p := []float32{}
	f.Float32SliceVarP(&p, name, "", value, usage)
	return &p
}

// Float32SliceP is the reexport of pflag.FlagSet.Float32SliceP()
//
// Float32SliceP is like Float32Slice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float32SliceP(name, shorthand string, value []float32, usage string) *[]float32 {
	p := []float32{}
	f.Float32SliceVarP(&p, name, shorthand, value, usage)
	return &p
}
