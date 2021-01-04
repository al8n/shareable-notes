package flag

// GetUint8 is the reexport of pflag.FlagSet.GetUint8()
//
// GetUint8 return the uint8 value of a flag with the given name
func (f *FlagSet) GetUint8(name string) (uint8, error) {
	return f.fs.GetUint8(name)
}

// Uint8Var is the reexport of pflag.FlagSet.Uint8Var()
//
// Uint8Var defines a uint8 flag with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the flag.
func (f *FlagSet) Uint8Var(p *uint8, name string, value uint8, usage string) {
	f.fs.Uint8Var(p, name, value, usage)
}

// Uint8VarP is the reexport of pflag.FlagSet.Uint8VarP()
//
// Uint8VarP is like Uint8Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint8VarP(p *uint8, name, shorthand string, value uint8, usage string) {
	f.fs.Uint8VarP(p, name, shorthand, value,usage)
}

// Uint8 is the reexport of pflag.FlagSet.Uint8()
//
// Uint8 defines a uint8 flag with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the flag.
func (f *FlagSet) Uint8(name string, value uint8, usage string) *uint8 {
	p := new(uint8)
	f.Uint8VarP(p, name, "", value, usage)
	return p
}

// Uint8P is the reexport of pflag.FlagSet.Uint8P()
//
// Uint8P is like Uint8, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint8P(name, shorthand string, value uint8, usage string) *uint8 {
	p := new(uint8)
	f.Uint8VarP(p, name, shorthand, value, usage)
	return p
}