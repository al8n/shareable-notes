package flag

// GetCount is the reexport of pflag.FlagSet.GetCount()
//
// GetCount return the int value of a flag with the given name
func (f *FlagSet) GetCount(name string) (int, error)  {
	return f.fs.GetCount(name)
}

// CountVar is the reexport of pflag.FlagSet.CountVar()
//
// CountVar defines a count flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
// A count flag will add 1 to its value every time it is found on the command line
func (f *FlagSet) CountVar(p *int, name string, usage string) {
	f.fs.CountVarP(p, name, "", usage)
}

// CountVarP is the reexport of pflag.FlagSet.CountVarP()
//
// CountVarP is like CountVar only take a shorthand for the flag name.
func (f *FlagSet) CountVarP(p *int, name, shorthand string, usage string) {
	f.fs.CountVarP(p, name, shorthand, usage)
}

// Count is the reexport of pflag.FlagSet.Count()
//
// Count defines a count flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
// A count flag will add 1 to its value every time it is found on the command line
func (f *FlagSet) Count(name string, usage string) *int {
	p := new(int)
	f.CountVarP(p, name, "", usage)
	return p
}

// CountP is the reexport of pflag.FlagSet.CountP()
//
// CountP is like Count only takes a shorthand for the flag name.
func (f *FlagSet) CountP(name, shorthand string, usage string) *int {
	p := new(int)
	f.CountVarP(p, name, shorthand, usage)
	return p
}