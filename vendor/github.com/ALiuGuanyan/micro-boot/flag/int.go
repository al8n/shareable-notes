package flag

// GetInt is the reexport of pflag.FlagSet.GetInt()
//
// GetInt return the int value of a flag with the given name
func (f *FlagSet) GetInt(name string) (int, error) {
	return f.fs.GetInt(name)
}

// IntVar is the reexport of pflag.FlagSet.IntVar()
//
// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
	f.fs.IntVar(p, name,value,usage)
}


// IntVarP is the reexport of pflag.FlagSet.IntVarP()
//
// IntVarP is like IntVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntVarP(p *int, name, shorthand string, value int, usage string) {
	f.fs.IntVarP(p,name,shorthand,value,usage)
}

// Int is the reexport of pflag.FlagSet.Int()
//
// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (f *FlagSet) Int(name string, value int, usage string) *int {
	p := new(int)
	f.IntVarP(p, name, "", value, usage)
	return p
}

// IntP is the reexport of pflag.FlagSet.IntP()
//
// IntP is like Int, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntP(name, shorthand string, value int, usage string) *int {
	p := new(int)
	f.IntVarP(p, name, shorthand, value, usage)
	return p
}