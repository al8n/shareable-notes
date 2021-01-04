package flag

import "net"

// GetIPNetSlice is the reexport of pflag.FlagSet.GetIPNetSlice()
//
// GetIPNetSlice returns the []net.IPNet value of a flag with the given name
func (f *FlagSet) GetIPNetSlice(name string) ([]net.IPNet, error) {
	return f.fs.GetIPNetSlice(name)
}

// IPNetSliceVar is the reexport of pflag.FlagSet.IPNetSliceVar()
//
// IPNetSliceVar defines a ipNetSlice flag with specified name, default value, and usage string.
// The argument p points to a []net.IPNet variable in which to store the value of the flag.
func (f *FlagSet) IPNetSliceVar(p *[]net.IPNet, name string, value []net.IPNet, usage string) {
	f.fs.IPNetSliceVar(p, name, value, usage)
}

// IPNetSliceVarP is the reexport of pflag.FlagSet.IPNetSliceVarP()
//
// IPNetSliceVarP is like IPNetSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPNetSliceVarP(p *[]net.IPNet, name, shorthand string, value []net.IPNet, usage string) {
	f.fs.IPNetSliceVarP(p, name, shorthand, value, usage)
}

// IPNetSlice is the reexport of pflag.FlagSet.IPNetSlice()
//
// IPNetSlice defines a []net.IPNet flag with specified name, default value, and usage string.
// The return value is the address of a []net.IPNet variable that stores the value of that flag.
func (f *FlagSet) IPNetSlice(name string, value []net.IPNet, usage string) *[]net.IPNet {
	p := []net.IPNet{}
	f.IPNetSliceVarP(&p, name, "", value, usage)
	return &p
}

// IPNetSliceP is the reexport of pflag.FlagSet.IPNetSliceP()
//
// IPNetSliceP is like IPNetSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPNetSliceP(name, shorthand string, value []net.IPNet, usage string) *[]net.IPNet {
	p := []net.IPNet{}
	f.IPNetSliceVarP(&p, name, shorthand, value, usage)
	return &p
}