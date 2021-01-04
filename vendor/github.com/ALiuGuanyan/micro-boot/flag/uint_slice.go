package flag

// GetUintSlice is the reexport of pflag.FlagSet.GetUintSlice()
//
// GetUintSlice returns the []uint value of a flag with the given name.
func (f *FlagSet) GetUintSlice(name string) ([]uint, error) {
	return f.fs.GetUintSlice(name)
}

// UintSliceVar is the reexport of pflag.FlagSet.UintSliceVar()
//
// UintSliceVar defines a uintSlice flag with specified name, default value, and usage string.
// The argument p points to a []uint variable in which to store the value of the flag.
func (f *FlagSet) UintSliceVar(p *[]uint, name string, value []uint, usage string) {
	f.fs.UintSliceVar(p, name,value,usage)
}

// UintSliceVarP is the reexport of pflag.FlagSet.UintSliceVarP()
//
// UintSliceVarP is like UintSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintSliceVarP(p *[]uint, name, shorthand string, value []uint, usage string) {
	f.fs.UintSliceVarP(p, name,shorthand,value,usage)
}

// UintSlice is the reexport of pflag.FlagSet.UintSlice()
//
// UintSlice defines a []uint flag with specified name, default value, and usage string.
// The return value is the address of a []uint variable that stores the value of the flag.
func (f *FlagSet) UintSlice(name string, value []uint, usage string) *[]uint {
	p := []uint{}
	f.UintSliceVarP(&p, name, "", value, usage)
	return &p
}

// UintSliceP is the reexport of pflag.FlagSet.UintSliceP()
//
// UintSliceP is like UintSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintSliceP(name, shorthand string, value []uint, usage string) *[]uint {
	p := []uint{}
	f.UintSliceVarP(&p, name, shorthand, value, usage)
	return &p
}