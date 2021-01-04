package flag

// GetInt64Slice is the reexport of pflag.FlagSet.GetInt64Slice()
//
// GetInt64Slice return the []int64 value of a flag with the given name
func (f *FlagSet) GetInt64Slice(name string) ([]int64, error) {
	return f.fs.GetInt64Slice(name)
}

// Int64SliceVar is the reexport of pflag.FlagSet.Int64SliceVar()
//
// Int64SliceVar defines a int64Slice flag with specified name, default value, and usage string.
// The argument p points to a []int64 variable in which to store the value of the flag.
func (f *FlagSet) Int64SliceVar(p *[]int64, name string, value []int64, usage string) {
	f.fs.Int64SliceVar(p, name, value, usage)
}

// Int64SliceVarP is the reexport of pflag.FlagSet.Int64SliceVarP()
//
// Int64SliceVarP is like Int64SliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int64SliceVarP(p *[]int64, name, shorthand string, value []int64, usage string) {
	f.fs.Int64SliceVarP(p, name, shorthand, value, usage)
}

// Int64Slice is the reexport of pflag.FlagSet.Int64Slice()
//
// Int64Slice defines a []int64 flag with specified name, default value, and usage string.
// The return value is the address of a []int64 variable that stores the value of the flag.
func (f *FlagSet) Int64Slice(name string, value []int64, usage string) *[]int64 {
	p := []int64{}
	f.Int64SliceVarP(&p, name, "", value, usage)
	return &p
}

// Int64SliceP is the reexport of pflag.FlagSet.Int64SliceP()
//
// Int64SliceP is like Int64Slice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int64SliceP(name, shorthand string, value []int64, usage string) *[]int64 {
	p := []int64{}
	f.Int64SliceVarP(&p, name, shorthand, value, usage)
	return &p
}