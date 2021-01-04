package flag

// GetInt8 is the reexport of pflag.FlagSet.GetInt8()
//
// GetInt8 return the int8 value of a flag with the given name
func (f *FlagSet) GetInt8(name string) (int8, error) {
	return f.fs.GetInt8(name)
}

// Int8Var is the reexport of pflag.FlagSet.Int8Var()
//
// Int8Var defines an int8 flag with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the flag.
func (f *FlagSet) Int8Var(p *int8, name string, value int8, usage string) {
	f.fs.Int8Var(p, name, value, usage)
}

// Int8VarP is the reexport of pflag.FlagSet.Int8VarP()
//
// Int8VarP is like Int8Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int8VarP(p *int8, name, shorthand string, value int8, usage string) {
	f.fs.Int8VarP(p, name, shorthand, value, usage)
}

// Int8 is the reexport of pflag.FlagSet.Int8()
//
// Int8 defines an int8 flag with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the flag.
func (f *FlagSet) Int8(name string, value int8, usage string) *int8 {
	p := new(int8)
	f.Int8VarP(p, name, "", value, usage)
	return p
}

// Int8P is the reexport of pflag.FlagSet.Int8P()
//
// Int8P is like Int8, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int8P(name, shorthand string, value int8, usage string) *int8 {
	p := new(int8)
	f.Int8VarP(p, name, shorthand, value, usage)
	return p
}