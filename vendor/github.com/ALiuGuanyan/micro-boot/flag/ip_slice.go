package flag

import "net"

// GetIPSlice is the reexport of pflag.FlagSet.GetIPSlice()
//
// GetIPSlice returns the []net.IP value of a flag with the given name
func (f *FlagSet) GetIPSlice(name string) ([]net.IP, error) {
	return f.fs.GetIPSlice(name)
}

// IPSliceVar is the reexport of pflag.FlagSet.IPSliceVar()
//
// IPSliceVar defines a ipSlice flag with specified name, default value, and usage string.
// The argument p points to a []net.IP variable in which to store the value of the flag.
func (f *FlagSet) IPSliceVar(p *[]net.IP, name string, value []net.IP, usage string) {
	f.fs.IPSliceVar(p, name,value, usage)
}

// IPSliceVarP is the reexport of pflag.FlagSet.IPSliceVarP()
//
// IPSliceVarP is like IPSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPSliceVarP(p *[]net.IP, name, shorthand string, value []net.IP, usage string) {
	f.fs.IPSliceVarP(p, name, shorthand, value, usage)
}

// IPSlice is the reexport of pflag.FlagSet.IPSlice()
//
// IPSlice defines a []net.IP flag with specified name, default value, and usage string.
// The return value is the address of a []net.IP variable that stores the value of that flag.
func (f *FlagSet) IPSlice(name string, value []net.IP, usage string) *[]net.IP {
	p := []net.IP{}
	f.IPSliceVarP(&p, name, "", value, usage)
	return &p
}

// IPSliceP is the reexport of pflag.FlagSet.IPSliceP()
//
// IPSliceP is like IPSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPSliceP(name, shorthand string, value []net.IP, usage string) *[]net.IP {
	p := []net.IP{}
	f.IPSliceVarP(&p, name, shorthand, value, usage)
	return &p
}