package flag

// GetString is the reexport of pflag.FlagSet.GetString()
//
// GetString return the string value of a flag with the given name
func (f *FlagSet) GetString(name string) (string, error) {
	return f.fs.GetString(name)
}

// StringVar is the reexport of pflag.FlagSet.StringVar
//
// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)  {
	f.fs.StringVar(p, name, value, usage)
}

// StringVarP is the reexport of pflag.FlagSet.StringVarP()
//
// StringVarP is like StringVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringVarP(p *string, name, shorthand string, value string, usage string) {
	f.fs.StringVarP(p, name, shorthand, value, usage)
}

// String is the reexport of pflag.FlagSet.String()
//
// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) String(name string, value string, usage string) *string {
	p := new(string)
	f.StringVarP(p, name, "", value, usage)
	return p
}

// StringP is the reexport of pflag.FlagSet.StringP()
//
// StringP is like String, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringP(name, shorthand string, value string, usage string) *string {
	p := new(string)
	f.StringVarP(p, name, shorthand, value, usage)
	return p
}

