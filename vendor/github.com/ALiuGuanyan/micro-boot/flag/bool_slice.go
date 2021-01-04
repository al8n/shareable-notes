package flag


// GetBoolSlice is the reexport of pflag.FlagSet.GetBoolSlice()
//
// GetBoolSlice returns the []bool value of a flag with the given name.
func (f *FlagSet) GetBoolSlice(name string) ([]bool, error) {
	return f.fs.GetBoolSlice(name)
}

// BoolSliceVar is the reexport of pflag.FlagSet.BoolSliceVar()
//
// BoolSliceVar defines a boolSlice flag with specified name, default value, and usage string.
// The argument p points to a []bool variable in which to store the value of the flag.
func (f *FlagSet) BoolSliceVar(p *[]bool, name string, value []bool, usage string) {
	f.fs.BoolSliceVar(p, name, value, usage)
}

// BoolSliceVarP is the reexport of pflag.FlagSet.BoolSliceVarP()
//
// BoolSliceVarP is like BoolSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolSliceVarP(p *[]bool, name, shorthand string, value []bool, usage string) {
	f.fs.BoolSliceVarP(p, name, shorthand, value, usage)
}

// BoolSlice is the reexport of pflag.FlagSet.BoolSlice()
//
// BoolSlice defines a []bool flag with specified name, default value, and usage string.
// The return value is the address of a []bool variable that stores the value of the flag.
func (f *FlagSet) BoolSlice(name string, value []bool, usage string) *[]bool {
	p := []bool{}
	f.BoolSliceVarP(&p, name, "", value, usage)
	return &p
}

// BoolSliceP is the reexport of pflag.FlagSet.BoolSliceP()
//
// BoolSliceP is like BoolSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolSliceP(name, shorthand string, value []bool, usage string) *[]bool {
	p := []bool{}
	f.BoolSliceVarP(&p, name, shorthand, value, usage)
	return &p
}