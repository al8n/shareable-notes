package flag

// GetInt32 is the reexport of pflag.FlagSet.GetInt32()
//
// GetInt32 return the int32 value of a flag with the given name
func (f *FlagSet) GetInt32(name string) (int32, error) {
	return f.fs.GetInt32(name)
}

// Int32Var is the reexport of pflag.FlagSet.Int32Var()
//
// Int32Var defines an int32 flag with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the flag.
func (f *FlagSet) Int32Var(p *int32, name string, value int32, usage string) {
	f.fs.Int32Var(p, name, value, usage)
}

// Int32VarP is the reexport of pflag.FlagSet.Int32VarP()
//
// Int32VarP is like Int32Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int32VarP(p *int32, name, shorthand string, value int32, usage string) {
	f.fs.Int32VarP(p, name, shorthand, value, usage)
}

// Int32 is the reexport of pflag.FlagSet.Int32()
//
// Int32 defines an int32 flag with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the flag.
func (f *FlagSet) Int32(name string, value int32, usage string) *int32 {
	p := new(int32)
	f.Int32VarP(p, name, "", value, usage)
	return p
}

// Int32P is the reexport of pflag.FlagSet.Int32P()
//
// Int32P is like Int32, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int32P(name, shorthand string, value int32, usage string) *int32 {
	p := new(int32)
	f.Int32VarP(p, name, shorthand, value, usage)
	return p
}