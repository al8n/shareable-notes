package flag

// GetInt64 is the reexport of pflag.FlagSet.GetInt64()
//
// GetInt64 return the int64 value of a flag with the given name
func (f *FlagSet) GetInt64(name string) (int64, error) {
	return f.fs.GetInt64(name)
}

// Int64Var is the reexport of pflag.FlagSet.Int64Var()
//
// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
	f.fs.Int64Var(p, name, value, usage)
}

// Int64VarP is the reexport of pflag.FlagSet.Int64VarP()
//
// Int64VarP is like Int64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int64VarP(p *int64, name, shorthand string, value int64, usage string) {
	f.fs.Int64VarP(p, name, shorthand, value, usage)
}

// Int64 is the reexport of pflag.FlagSet.Int64()
//
// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (f *FlagSet) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64VarP(p, name, "", value, usage)
	return p
}

// Int64P is the reexport of pflag.FlagSet.Int64P()
//
// Int64P is like Int64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int64P(name, shorthand string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64VarP(p, name, shorthand, value, usage)
	return p
}
