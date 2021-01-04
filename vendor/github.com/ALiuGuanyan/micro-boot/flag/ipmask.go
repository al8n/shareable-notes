package flag

import "net"

// GetIPv4Mask is the reexport of pflag.FlagSet.GetIPv4Mask()
//
// GetIPv4Mask return the net.IPv4Mask value of a flag with the given name
func (f *FlagSet) GetIPv4Mask(name string) (net.IPMask, error) {
	return f.fs.GetIPv4Mask(name)
}

// IPMaskVar is the reexport of pflag.FlagSet.IPMaskVar()
//
// IPMaskVar defines an net.IPMask flag with specified name, default value, and usage string.
// The argument p points to an net.IPMask variable in which to store the value of the flag.
func (f *FlagSet) IPMaskVar(p *net.IPMask, name string, value net.IPMask, usage string) {
	f.fs.IPMaskVar(p, name, value, usage)
}

// IPMaskVarP is the reexport of pflag.FlagSet.IPMaskVarP()
//
// IPMaskVarP is like IPMaskVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPMaskVarP(p *net.IPMask, name, shorthand string, value net.IPMask, usage string) {
	f.fs.IPMaskVarP(p, name, shorthand,value, usage)
}

// IPMask is the reexport of pflag.FlagSet.IPMask()
//
// IPMask defines an net.IPMask flag with specified name, default value, and usage string.
// The return value is the address of an net.IPMask variable that stores the value of the flag.
func (f *FlagSet) IPMask(name string, value net.IPMask, usage string) *net.IPMask {
	p := new(net.IPMask)
	f.IPMaskVarP(p, name, "", value, usage)
	return p
}

// IPMaskP is the reexport of pflag.FlagSet.IPMaskP()
//
// IPMaskP is like IPMask, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPMaskP(name, shorthand string, value net.IPMask, usage string) *net.IPMask {
	p := new(net.IPMask)
	f.IPMaskVarP(p, name, shorthand, value, usage)
	return p
}