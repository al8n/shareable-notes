package flag

// GetBool is the reexport of pflag.FlagSet.GetBool()
//
// GetBool return the bool value of a flag with the given name
func (f *FlagSet) GetBool(name string) (bool, error) {
	return f.fs.GetBool(name)
}

// BoolVar is the reexport of pflag.FlagSet.BoolVar()
//
// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
	f.BoolVarP(p, name, "", value, usage)
}

// BoolVarP is the reexport of pflag.FlagSet.BoolVarP()
//
// BoolVarP is like BoolVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	f.fs.BoolVarP(p, name, shorthand, value, usage)
}

// Bool is the reexport of pflag.FlagSet.Bool()
//
// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
	return f.BoolP(name, "", value, usage)
}

// BoolP is the reexport of pflag.FlagSet.BoolP()
//
// BoolP is like Bool, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolP(name, shorthand string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVarP(p, name, shorthand, value, usage)
	return p
}

