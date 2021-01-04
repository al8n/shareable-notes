package mongo

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
)

var (
	defaultReadConcernFlagsPrefix = defaultMongoFlagsPrefix + "-read-concern"
	defaultReadConcernLevel = ""
)

const (
	defaultReadConcernLevelUsage = "specifies the read concern to use for read operations"
)

func SetDefaultReadConcernFlagsPrefix(val string)  {
	defaultReadConcernFlagsPrefix = val
}

func SetDefaultReadConcernLevel(val string)  {
	defaultReadConcernLevel = val
}


// ReadConcern for replica sets and replica set shards determines which data to return from a query.
type ReadConcern struct {
	Level  string  `json:"level" yaml:"level"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (rc *ReadConcern) BindFlags(fs *bootflag.FlagSet)  {
	if rc.CustomBindFlagsFunc != nil {
		rc.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(&rc.Level, utils.BuildFlagName(defaultReadConcernFlagsPrefix, "level"), defaultReadConcernLevel, defaultReadConcernLevelUsage)
}

func (rc *ReadConcern) Parse() (err error) {
	if rc.CustomParseFunc != nil {
		return rc.CustomParseFunc()
	}
	return nil
}

