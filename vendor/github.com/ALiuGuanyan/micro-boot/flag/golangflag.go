package flag

import (
	goflag "flag"
)

// AddGoFlag is the reexport of pflag.FlagSet.AddGoFlag()
//
// AddGoFlag will add the given *flag.Flag to the pflag.FlagSet
func (f *FlagSet) AddGoFlag(goflag *goflag.Flag) {
	f.fs.AddGoFlag(goflag)
}

// AddGoFlagSet is the reexport of pflag.FlagSet.AddGoFlagSet()
//
// AddGoFlagSet will add the given *flag.FlagSet to the pflag.FlagSet
func (f *FlagSet) AddGoFlagSet(newSet *goflag.FlagSet) {
	f.fs.AddGoFlagSet(newSet)
}