package flag

// GetInt16 is the reexport of pflag.FlagSet.GetInt16()
//
// GetInt16 return the int8 value of a flag with the given name
func (f *FlagSet) GetInt16(name string) (int16, error) {
	return f.fs.GetInt16(name)
}

// Int16Var is the reexport of pflag.FlagSet.Int16Var()
//
// Int16Var defines an int8 flag with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the flag.
func (f *FlagSet) Int16Var(p *int16, name string, value int16, usage string) {
	f.fs.Int16Var(p, name, value, usage)
}

// Int16VarP is the reexport of pflag.FlagSet.Int16VarP()
//
// Int16VarP is like Int16Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int16VarP(p *int16, name, shorthand string, value int16, usage string) {
	f.fs.Int16VarP(p, name, shorthand, value, usage)
}

// Int16 is the reexport of pflag.FlagSet.Int16()
//
// Int16 defines an int8 flag with specified name, default value, and usage string.
// The return value is the address of an Int16 variable that stores the value of the flag.
func (f *FlagSet) Int16(name string, value int16, usage string) *int16 {
	p := new(int16)
	f.Int16VarP(p, name, "", value, usage)
	return p
}

// Int16P is the reexport of pflag.FlagSet.Int16P()
//
// Int16P is like Int16, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int16P(name, shorthand string, value int16, usage string) *int16 {
	p := new(int16)
	f.Int16VarP(p, name, shorthand, value, usage)
	return p
}