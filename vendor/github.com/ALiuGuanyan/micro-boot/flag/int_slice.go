package flag

// GetIntSlice is the reexport of pflag.FlagSet.GetIntSlice()
//
// GetIntSlice return the []int value of a flag with the given name
func (f *FlagSet) GetIntSlice(name string) ([]int, error) {
	return f.fs.GetIntSlice(name)
}

// IntSliceVar is the reexport of pflag.FlagSet.IntSliceVar()
//
// IntSliceVar defines a intSlice flag with specified name, default value, and usage string.
// The argument p points to a []int variable in which to store the value of the flag.
func (f *FlagSet) IntSliceVar(p *[]int, name string, value []int, usage string) {
	f.fs.IntSliceVar(p, name, value, usage)
}

// IntSliceVarP is the reexport of pflag.FlagSet.IntSliceVarP()
//
// IntSliceVarP is like IntSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string) {
	f.fs.IntSliceVarP(p,name,shorthand,value,usage)
}

// IntSlice is the reexport of pflag.FlagSet.IntSlice()
//
// IntSlice defines a []int flag with specified name, default value, and usage string.
// The return value is the address of a []int variable that stores the value of the flag.
func (f *FlagSet) IntSlice(name string, value []int, usage string) *[]int {
	p := []int{}
	f.IntSliceVarP(&p, name, "", value, usage)
	return &p
}

// IntSliceP is the reexport of pflag.FlagSet.IntSliceP()
//
// IntSliceP is like IntSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntSliceP(name, shorthand string, value []int, usage string) *[]int {
	p := []int{}
	f.IntSliceVarP(&p, name, shorthand, value, usage)
	return &p
}