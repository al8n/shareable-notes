package flag

// GetStringToInt is the reexport of pflag.FlagSet.GetStringToInt()
//
// GetStringToInt return the map[string]int value of a flag with the given name
func (f *FlagSet) GetStringToInt(name string) (map[string]int, error) {
	return f.fs.GetStringToInt(name)
}

// StringToIntVar is the reexport of pflag.FlagSet.StringToIntVar()
//
// StringToIntVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a map[string]int variable in which to store the values of the multiple flags.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToIntVar(p *map[string]int, name string, value map[string]int, usage string) {
	f.fs.StringToIntVar(p, name, value, usage)
}

// StringToIntVarP is the reexport of pflag.FlagSet.StringToIntVarP()
//
// StringToIntVarP is like StringToIntVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToIntVarP(p *map[string]int, name, shorthand string, value map[string]int, usage string) {
	f.fs.StringToIntVarP(p, name, shorthand, value, usage)
}

// StringToInt is the reexport of pflag.FlagSet.StringToInt()
//
// StringToInt defines a string flag with specified name, default value, and usage string.
// The return value is the address of a map[string]int variable that stores the value of the flag.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToInt(name string, value map[string]int, usage string) *map[string]int {
	p := map[string]int{}
	f.StringToIntVarP(&p, name, "", value, usage)
	return &p
}

// StringToIntP is the reexport of pflag.FlagSet.StringToIntP()
//
// StringToIntP is like StringToInt, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToIntP(name, shorthand string, value map[string]int, usage string) *map[string]int {
	p := map[string]int{}
	f.StringToIntVarP(&p, name, shorthand, value, usage)
	return &p
}
