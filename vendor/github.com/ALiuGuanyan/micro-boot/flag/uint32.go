package flag

// GetUint32 is the reexport of pflag.FlagSet.GetUint32()
//
// GetUint32 return the uint32 value of a flag with the given name
func (f *FlagSet) GetUint32(name string) (uint32, error) {
	return f.fs.GetUint32(name)
}

// Uint32Var is the reexport of pflag.FlagSet.Uint32Var()
//
// Uint32Var defines a uint32 flag with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the flag.
func (f *FlagSet) Uint32Var(p *uint32, name string, value uint32, usage string) {
	f.fs.Uint32Var(p, name, value, usage)
}

// Uint32VarP is the reexport of pflag.FlagSet.Uint32VarP()
//
// Uint32VarP is like Uint32Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint32VarP(p *uint32, name, shorthand string, value uint32, usage string) {
	f.fs.Uint32VarP(p, name, shorthand, value, usage)
}

// Uint32 is the reexport of pflag.FlagSet.Uint32()
//
// Uint32 defines a uint32 flag with specified name, default value, and usage string.
// The return value is the address of a uint32  variable that stores the value of the flag.
func (f *FlagSet) Uint32(name string, value uint32, usage string) *uint32 {
	p := new(uint32)
	f.Uint32VarP(p, name, "", value, usage)
	return p
}

// Uint32P is the reexport of pflag.FlagSet.Uint32P()
//
// Uint32P is like Uint32, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint32P(name, shorthand string, value uint32, usage string) *uint32 {
	p := new(uint32)
	f.Uint32VarP(p, name, shorthand, value, usage)
	return p
}
