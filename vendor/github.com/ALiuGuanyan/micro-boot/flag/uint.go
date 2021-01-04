package flag

// GetUint is the reexport of pflag.FlagSet.GetUint()
//
// GetUint return the uint value of a flag with the given name
func (f *FlagSet) GetUint(name string) (uint, error) {
	return f.fs.GetUint(name)
}

// UintVar is the reexport of pflag.FlagSet.UintVar()
//
// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
	f.fs.UintVar(p, name, value, usage)
}

// UintVarP is the reexport of pflag.FlagSet.UintVarP()
//
// UintVarP is like UintVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintVarP(p *uint, name, shorthand string, value uint, usage string) {
	f.fs.UintVarP(p, name,shorthand,value,usage)
}

// Uint is the reexport of pflag.FlagSet.Uint()
//
// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVarP(p, name, "", value, usage)
	return p
}

// UintP is the reexport of pflag.FlagSet.UintP()
//
// UintP is like Uint, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintP(name, shorthand string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVarP(p, name, shorthand, value, usage)
	return p
}