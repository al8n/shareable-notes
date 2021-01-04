package flag

import "net"

// GetIPNet is the reexport of pflag.FlagSet.GetIPNet()
//
// GetIPNet return the net.IPNet value of a flag with the given name
func (f *FlagSet) GetIPNet(name string) (net.IPNet, error) {
	return f.fs.GetIPNet(name)
}

// IPNetVar is the reexport of pflag.FlagSet.IPNetVar()
//
// IPNetVar defines an net.IPNet flag with specified name, default value, and usage string.
// The argument p points to an net.IPNet variable in which to store the value of the flag.
func (f *FlagSet) IPNetVar(p *net.IPNet, name string, value net.IPNet, usage string) {
	f.fs.IPNetVar(p, name, value, usage)
}

// IPNetVarP is the reexport of pflag.FlagSet.IPNetVarP()
//
// IPNetVarP is like IPNetVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPNetVarP(p *net.IPNet, name, shorthand string, value net.IPNet, usage string) {
	f.fs.IPNetVarP(p, name, shorthand, value, usage)
}

// IPNet is the reexport of pflag.FlagSet.IPNet()
//
// IPNet defines an net.IPNet flag with specified name, default value, and usage string.
// The return value is the address of an net.IPNet variable that stores the value of the flag.
func (f *FlagSet) IPNet(name string, value net.IPNet, usage string) *net.IPNet {
	p := new(net.IPNet)
	f.IPNetVarP(p, name, "", value, usage)
	return p
}

// IPNetP is the reexport of pflag.FlagSet.IPNetP()
//
// IPNetP is like IPNet, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPNetP(name, shorthand string, value net.IPNet, usage string) *net.IPNet {
	p := new(net.IPNet)
	f.IPNetVarP(p, name, shorthand, value, usage)
	return p
}