package flag

// GetStringArray is the reexport of pflag.FlagSet.GetStringArray()
//
// GetStringArray return the []string value of a flag with the given name
func (f *FlagSet) GetStringArray(name string) ([]string, error) {
	return f.fs.GetStringArray(name)
}

// StringArrayVar is the reexport of pflag.FlagSet.StringArrayVar()
//
// StringArrayVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the values of the multiple flags.
// The value of each argument will not try to be separated by comma. Use a StringSlice for that.
func (f *FlagSet) StringArrayVar(p *[]string, name string, value []string, usage string) {
	f.fs.StringArrayVar(p, name,value, usage)
}

// StringArrayVarP is the reexport of pflag.FlagSet.StringArrayVarP()
//
// StringArrayVarP is like StringArrayVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringArrayVarP(p *[]string, name, shorthand string, value []string, usage string) {
	f.fs.StringArrayVarP(p, name, shorthand, value, usage)
}

// StringArray is the reexport of pflag.FlagSet.StringArray()
//
// StringArray defines a string flag with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the flag.
// The value of each argument will not try to be separated by comma. Use a StringSlice for that.
func (f *FlagSet) StringArray(name string, value []string, usage string) *[]string {
	p := []string{}
	f.StringArrayVarP(&p, name, "", value, usage)
	return &p
}

// StringArrayP is the reexport of pflag.FlagSet.StringArrayP()
//
// StringArrayP is like StringArray, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringArrayP(name, shorthand string, value []string, usage string) *[]string {
	p := []string{}
	f.StringArrayVarP(&p, name, shorthand, value, usage)
	return &p
}