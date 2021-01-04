package flag

// GetUint64 is the reexport of pflag.FlagSet.GetUint64()
//
// GetUint64 return the uint64 value of a flag with the given name
func (f *FlagSet) GetUint64(name string) (uint64, error) {
	return f.fs.GetUint64(name)
}

// Uint64Var is the reexport of pflag.FlagSet.Uint64Var()
//
// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.fs.Uint64Var(p, name,value,usage)
}

// Uint64VarP is the reexport of pflag.FlagSet.Uint64VarP()
//
// Uint64VarP is like Uint64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string) {
	f.fs.Uint64VarP(p, name,shorthand,value,usage)
}

// Uint64 is the reexport of pflag.FlagSet.Uint64()
//
// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64VarP(p, name, "", value, usage)
	return p
}

// Uint64P is the reexport of pflag.FlagSet.Uint64P()
//
// Uint64P is like Uint64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint64P(name, shorthand string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64VarP(p, name, shorthand, value, usage)
	return p
}