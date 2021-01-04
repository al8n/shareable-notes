package flag

// GetStringToString is the reexport of pflag.FlagSet.GetStringToString()
//
// GetStringToString return the map[string]string value of a flag with the given name
func (f *FlagSet) GetStringToString(name string) (map[string]string, error) {
	return f.fs.GetStringToString(name)
}

// StringToStringVar is the reexport of pflag.FlagSet.StringToStringVar()
//
// StringToStringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a map[string]string variable in which to store the values of the multiple flags.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToStringVar(p *map[string]string, name string, value map[string]string, usage string) {
	f.fs.StringToStringVar(p,name,value,usage)
}

// StringToStringVarP is the reexport of pflag.FlagSet.StringToStringVarP()
//
// StringToStringVarP is like StringToStringVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToStringVarP(p *map[string]string, name, shorthand string, value map[string]string, usage string) {
	f.fs.StringToStringVarP(p, name, shorthand, value, usage)
}

// StringToString is the reexport of pflag.FlagSet.StringToString()
//
// StringToString defines a string flag with specified name, default value, and usage string.
// The return value is the address of a map[string]string variable that stores the value of the flag.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToString(name string, value map[string]string, usage string) *map[string]string {
	p := map[string]string{}
	f.StringToStringVarP(&p, name, "", value, usage)
	return &p
}

// StringToStringP is the reexport of pflag.FlagSet.StringToStringP()
//
// StringToStringP is like StringToString, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToStringP(name, shorthand string, value map[string]string, usage string) *map[string]string {
	p := map[string]string{}
	f.StringToStringVarP(&p, name, shorthand, value, usage)
	return &p
}