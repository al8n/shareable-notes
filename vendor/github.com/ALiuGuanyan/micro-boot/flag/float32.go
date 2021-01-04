package flag

// GetFloat32 is the reexport of pflag.FlagSet.GetFloat32()
//
// GetFloat32 return the float32 value of a flag with the given name
func (f *FlagSet) GetFloat32(name string) (float32, error) {
	return f.fs.GetFloat32(name)
}

// Float32Var is the reexport of pflag.FlagSet.Float32Var()
//
// Float32Var defines a float32 flag with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the flag.
func (f *FlagSet) Float32Var(p *float32, name string, value float32, usage string) {
	f.fs.Float32Var(p,name,value,usage)
}

// Float32VarP is the reexport of pflag.FlagSet.Float32VarP()
//
// Float32VarP is like Float32Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float32VarP(p *float32, name, shorthand string, value float32, usage string) {
	f.fs.Float32VarP(p, name,shorthand, value, usage)
}

// Float32 is the reexport of pflag.FlagSet.Float32()
//
// Float32 defines a float32 flag with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the flag.
func (f *FlagSet) Float32(name string, value float32, usage string) *float32 {
	p := new(float32)
	f.Float32VarP(p, name, "", value, usage)
	return p
}

// Float32P is the reexport of pflag.FlagSet.Float32P()
//
// Float32P is like Float32, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float32P(name, shorthand string, value float32, usage string) *float32 {
	p := new(float32)
	f.Float32VarP(p, name, shorthand, value, usage)
	return p
}
