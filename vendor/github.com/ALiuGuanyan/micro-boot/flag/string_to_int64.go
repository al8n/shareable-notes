package flag

// GetStringToInt64 is the reexport of pflag.FlagSet.GetStringToInt64()
//
// GetStringToInt64 return the map[string]int64 value of a flag with the given name
func (f *FlagSet) GetStringToInt64(name string) (map[string]int64, error) {
	return f.fs.GetStringToInt64(name)
}

// StringToInt64Var is the reexport of pflag.FlagSet.StringToInt64Var()
//
// StringToInt64Var defines a string flag with specified name, default value, and usage string.
// The argument p point64s to a map[string]int64 variable in which to store the values of the multiple flags.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToInt64Var(p *map[string]int64, name string, value map[string]int64, usage string) {
	f.fs.StringToInt64Var(p, name, value, usage)
}

// StringToInt64VarP is the reexport of pflag.FlagSet.StringToInt64VarP()
//
// StringToInt64VarP is like StringToInt64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToInt64VarP(p *map[string]int64, name, shorthand string, value map[string]int64, usage string) {
	f.fs.StringToInt64VarP(p, name, shorthand, value, usage)
}

// StringToInt64 is the reexport of pflag.FlagSet.StringToInt64()
//
// StringToInt64 defines a string flag with specified name, default value, and usage string.
// The return value is the address of a map[string]int64 variable that stores the value of the flag.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToInt64(name string, value map[string]int64, usage string) *map[string]int64 {
	p := map[string]int64{}
	f.StringToInt64VarP(&p, name, "", value, usage)
	return &p
}

// StringToInt64P is the reexport of pflag.FlagSet.StringToInt64P()
//
// StringToInt64P is like StringToInt64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToInt64P(name, shorthand string, value map[string]int64, usage string) *map[string]int64 {
	p := map[string]int64{}
	f.StringToInt64VarP(&p, name, shorthand, value, usage)
	return &p
}