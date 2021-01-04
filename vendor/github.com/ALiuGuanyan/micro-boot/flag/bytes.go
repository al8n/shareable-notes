package flag

// GetBytesHex is the reexport of pflag.FlagSet.GetBytesHex()
//
// GetBytesHex return the []byte value of a flag with the given name
func (f *FlagSet) GetBytesHex(name string) ([]byte, error) {
	return f.fs.GetBytesHex(name)
}

// BytesHexVar is the reexport of pflag.FlagSet.BytesHexVar()
//
// BytesHexVar defines an []byte flag with specified name, default value, and usage string.
// The argument p points to an []byte variable in which to store the value of the flag.
func (f *FlagSet) BytesHexVar(p *[]byte, name string, value []byte, usage string) {
	f.fs.BytesHexVar(p, name, value, usage)
}

// BytesHexVarP is the reexport of pflag.FlagSet.BytesHexVarP()
//
// BytesHexVarP is like BytesHexVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string) {
	f.fs.BytesHexVarP(p, name, shorthand, value, usage)
}

// BytesHex is the reexport of pflag.FlagSet.BytesHex()
//
// BytesHex defines an []byte flag with specified name, default value, and usage string.
// The return value is the address of an []byte variable that stores the value of the flag.
func (f *FlagSet) BytesHex(name string, value []byte, usage string) *[]byte {
	p := new([]byte)
	f.BytesHexVarP(p, name, "", value, usage)
	return p
}

// BytesHexP is the reexport of pflag.FlagSet.BytesHexP()
//
// BytesHexP is like BytesHex, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BytesHexP(name, shorthand string, value []byte, usage string) *[]byte {
	p := new([]byte)
	f.BytesHexVarP(p, name, shorthand, value, usage)
	return p
}

// GetBytesBase64 is the reexport of pflag.FlagSet.GetBytesBase64()
//
// GetBytesBase64 return the []byte value of a flag with the given name
func (f *FlagSet) GetBytesBase64(name string) ([]byte, error) {
	return f.fs.GetBytesBase64(name)
}

// BytesBase64Var is the reexport of pflag.FlagSet.BytesBase64Var()
//
// BytesBase64Var defines an []byte flag with specified name, default value, and usage string.
// The argument p points to an []byte variable in which to store the value of the flag.
func (f *FlagSet) BytesBase64Var(p *[]byte, name string, value []byte, usage string) {
	f.fs.BytesBase64Var(p, name, value, usage)
}

// BytesBase64VarP is the reexport of pflag.FlagSet.BytesBase64VarP()
//
// BytesBase64VarP is like BytesBase64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BytesBase64VarP(p *[]byte, name, shorthand string, value []byte, usage string) {
	f.fs.BytesBase64VarP(p, name, shorthand, value, usage)
}
