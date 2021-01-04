package flag

// GetInt32Slice is the reexport of pflag.FlagSet.GetInt32Slice()
//
// GetInt32Slice return the []int32 value of a flag with the given name
func (f *FlagSet) GetInt32Slice(name string) ([]int32, error) {
	return f.fs.GetInt32Slice(name)
}

// Int32SliceVar is the reexport of pflag.FlagSet.Int32SliceVar()
//
// Int32SliceVar defines a int32Slice flag with specified name, default value, and usage string.
// The argument p points to a []int32 variable in which to store the value of the flag.
func (f *FlagSet) Int32SliceVar(p *[]int32, name string, value []int32, usage string) {
	f.fs.Int32SliceVar(p, name, value, usage)
}

// Int32SliceVarP is the reexport of pflag.FlagSet.Int32SliceVarP()
//
// Int32SliceVarP is like Int32SliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int32SliceVarP(p *[]int32, name, shorthand string, value []int32, usage string) {
	f.fs.Int32SliceVarP(p, name,shorthand,value,usage)
}

// Int32Slice is the reexport of pflag.FlagSet.Int32Slice()
//
// Int32Slice defines a []int32 flag with specified name, default value, and usage string.
// The return value is the address of a []int32 variable that stores the value of the flag.
func (f *FlagSet) Int32Slice(name string, value []int32, usage string) *[]int32 {
	p := []int32{}
	f.Int32SliceVarP(&p, name, "", value, usage)
	return &p
}

// Int32SliceP is the reexport of pflag.FlagSet.Int32SliceP()
//
// Int32SliceP is like Int32Slice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int32SliceP(name, shorthand string, value []int32, usage string) *[]int32 {
	p := []int32{}
	f.Int32SliceVarP(&p, name, shorthand, value, usage)
	return &p
}
