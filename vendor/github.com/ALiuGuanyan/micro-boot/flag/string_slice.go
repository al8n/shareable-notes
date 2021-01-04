package flag

// GetStringSlice is the reexport of pflag.FlagSet.GetStringSlice()
//
// GetStringSlice return the []string value of a flag with the given name
func (f *FlagSet) GetStringSlice(name string) ([]string, error) {
	return f.fs.GetStringSlice(name)
}

// StringSliceVar is the reexport of pflag.FlagSet.StringSliceVar()
//
// StringSliceVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag.
// Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly.
// For example:
//   --ss="v1,v2" --ss="v3"
// will result in
//   []string{"v1", "v2", "v3"}
func (f *FlagSet) StringSliceVar(p *[]string, name string, value []string, usage string) {
	f.fs.StringSliceVar(p, name, value, usage)
}

// StringSliceVarP is the reexport of pflag.FlagSet.StringSliceVarP()
//
// StringSliceVarP is like StringSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringSliceVarP(p *[]string, name, shorthand string, value []string, usage string) {
	f.fs.StringSliceVarP(p, name,shorthand,value,usage)
}

// StringSlice is the reexport of pflag.FlagSet.StringSlice()
//
// StringSlice defines a string flag with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the flag.
// Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly.
// For example:
//   --ss="v1,v2" --ss="v3"
// will result in
//   []string{"v1", "v2", "v3"}
func (f *FlagSet) StringSlice(name string, value []string, usage string) *[]string {
	p := []string{}
	f.StringSliceVarP(&p, name, "", value, usage)
	return &p
}

// StringSliceP is the reexport of pflag.FlagSet.StringSliceP()
//
// StringSliceP is like StringSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringSliceP(name, shorthand string, value []string, usage string) *[]string {
	p := []string{}
	f.StringSliceVarP(&p, name, shorthand, value, usage)
	return &p
}