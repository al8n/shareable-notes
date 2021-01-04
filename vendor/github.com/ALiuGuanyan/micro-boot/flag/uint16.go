package flag

// GetUint16 is the reexport of pflag.FlagSet.GetUint16()
//
// GetUint16 return the uint16 value of a flag with the given name
func (f *FlagSet) GetUint16(name string) (uint16, error) {
	return f.fs.GetUint16(name)
}

// Uint16Var is the reexport of pflag.FlagSet.Uint16Var()
//
// Uint16Var defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) Uint16Var(p *uint16, name string, value uint16, usage string) {
	f.fs.Uint16Var(p, name, value, usage)
}

// Uint16VarP is the reexport of pflag.FlagSet.Uint16VarP()
//
// Uint16VarP is like Uint16Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string) {
	f.fs.Uint16VarP(p, name, shorthand, value,usage)
}

// Uint16 is the reexport of pflag.FlagSet.Uint16()
//
// Uint16 defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) Uint16(name string, value uint16, usage string) *uint16 {
	p := new(uint16)
	f.Uint16VarP(p, name, "", value, usage)
	return p
}

// Uint16P is the reexport of pflag.FlagSet.Uint16P()
//
// Uint16P is like Uint16, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint16P(name, shorthand string, value uint16, usage string) *uint16 {
	p := new(uint16)
	f.Uint16VarP(p, name, shorthand, value, usage)
	return p
}
