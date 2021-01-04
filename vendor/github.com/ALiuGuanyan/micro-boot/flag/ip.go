package flag

import "net"

// GetIP is the reexport of pflag.FlagSet.GetIP()
//
// GetIP return the net.IP value of a flag with the given name
func (f *FlagSet) GetIP(name string) (net.IP, error) {
	return f.fs.GetIP(name)
}

// IPVar is the reexport of pflag.FlagSet.IPVar()
//
// IPVar defines an net.IP flag with specified name, default value, and usage string.
// The argument p points to an net.IP variable in which to store the value of the flag.
func (f *FlagSet) IPVar(p *net.IP, name string, value net.IP, usage string) {
	f.fs.IPVar(p, name, value, usage)
}

// IPVarP is the reexport of pflag.FlagSet.IPVarP()
//
// IPVarP is like IPVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPVarP(p *net.IP, name, shorthand string, value net.IP, usage string) {
	f.fs.IPVarP(p, name,shorthand,value, usage)
}

// IP is the reexport of pflag.FlagSet.IP()
//
// IP defines an net.IP flag with specified name, default value, and usage string.
// The return value is the address of an net.IP variable that stores the value of the flag.
func (f *FlagSet) IP(name string, value net.IP, usage string) *net.IP {
	p := new(net.IP)
	f.IPVarP(p, name, "", value, usage)
	return p
}

// IPP is the reexport of pflag.FlagSet.IPP()
//
// IPP is like IP, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPP(name, shorthand string, value net.IP, usage string) *net.IP {
	p := new(net.IP)
	f.IPVarP(p, name, shorthand, value, usage)
	return p
}
