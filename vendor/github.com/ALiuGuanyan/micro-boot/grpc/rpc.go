package grpc

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"strconv"
)

var (
	defaultGRPCFlagsPrefix = "rpc"
	defaultGRPCRunnable = true
	defaultGRPCPort = "")

func SetDefaultGRPCFlagsPrefix(val string)  {
	defaultGRPCFlagsPrefix = val
}

func SetDefaultGRPCRunnable(val bool)  {
	defaultGRPCRunnable = val
}

func SetDefaultGRPCPort(val string)  {
	defaultGRPCPort = val
}

type GRPC struct {
	Name string `json:"name" yaml:"name"`
	Runnable          bool          `json:"runnable" yaml:"runnable"`
	Port              string        `json:"port" yaml:"port"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
	//ConnectionTimeout     time.Duration `json:"connection-timeout" yaml:"connection-timeout" default:"120s"`
	//MaxSendMessageSize    int           `json:"max-send-message-size" json:"max-send-message-size" default:"1073741824"`
	//MaxReceiveMessageSize int           `json:"max-receive-message-size" json:"max-receive-message-size" default:"4194304"`
	//WriteBufferSize       int           `json:"write-buffer-size" yaml:"write-buffer-size" default:"32768"`
	//ReadBufferSize        int           `json:"read-buffer-size" yaml:"read-buffer-size" default:"32768"`
	//
	//InitialWindowSize     int32 `json:"initial-window-size" yaml:"initial-window-size" default:"64"`
	//InitialConnWindowSize int32 `json:"initial-conn-window-size" yaml:"initial-conn-window-size"`
	//
	//KeepAlive KeepAlive `json:"keep-alive" yaml:"keep-alive"`
	//
	//MaxHeaderListSize     uint32 `json:"max-header-list-size" yaml:"max-header-list-size"`
	//HeaderTableSize       uint32 `json:"header-table-size" yaml:"header-table-size"`
	//NumServerWorkers      uint32 `json:"num-server-workers" yaml:"num-server-workers"`
}

func (r *GRPC) GetIntPort() (port int) {
	var port64 int64
	port64, _ = strconv.ParseInt(r.Port, 10, 64)
	return int(port64)
}

func (r *GRPC) BindFlags(fs *bootflag.FlagSet)  {
	if r.CustomBindFlagsFunc != nil {
		r.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&r.Name,
		utils.BuildFlagName(defaultGRPCFlagsPrefix, "name"),
		"",
		"set the grpc service name")

	fs.BoolVar(
		&r.Runnable,
		utils.BuildFlagName(defaultGRPCFlagsPrefix, "runnable"),
		defaultGRPCRunnable,
		"run GRPC server or not (default is true)")


	fs.IPv4PortStringVar(
		&r.Port,
		utils.BuildFlagName(defaultGRPCFlagsPrefix,"port"),
		defaultGRPCPort,
		"GRPC listen port")

}

func (r *GRPC) Parse() (err error) {
	if r.CustomParseFunc != nil {
		return r.CustomParseFunc()
	}
	return nil
}

//type ServerParameters struct {
//	MaxConnectionIdle time.Duration `json:"max-connection-idle" yaml:"max-connection-idle" default:"infinity"`
//
//	MaxConnectionAge time.Duration `json:"max-connection-age" yaml:"max-connection-age" default:"infinity"`
//
//	MaxConnectionAgeGrace time.Duration `json:"max-connection-age-grace" yaml:"max-connection-age-grace" default:"infinity"`
//
//	Time time.Duration `json:"time" yaml:"time" default:"2h"`
//
//	Timeout time.Duration `json:"timeout" yaml:"timeout" default:"20s"`
//}
//
//func (sp ServerParameters) ToStd() keepalive.ServerParameters {
//	return keepalive.ServerParameters{
//		MaxConnectionIdle:     sp.MaxConnectionIdle,
//		MaxConnectionAge:      sp.MaxConnectionAge,
//		MaxConnectionAgeGrace: sp.MaxConnectionAgeGrace,
//		Time:                  sp.Time,
//		Timeout:               sp.Timeout,
//	}
//}
//
//type EnforcementPolicy struct {
//	// The current default value is 5 minutes.
//	MinTime time.Duration `json:"min-time" yaml:"min-time" default:"5m"`
//	// false by default.
//	PermitWithoutStream bool `json:"permit-without-stream" yaml:"permit-without-stream" default:"false"`
//}
//
//func (ep EnforcementPolicy) ToStd() keepalive.EnforcementPolicy {
//	return keepalive.EnforcementPolicy{
//		MinTime:             ep.MinTime,
//		PermitWithoutStream: ep.PermitWithoutStream,
//	}
//}
//
//type KeepAlive struct {
//	ServerParameters  ServerParameters  `json:"server-parameters" yaml:"server-parameters"`
//	EnforcementPolicy EnforcementPolicy `json:"enforcement-policy" yaml:"enforcement-policy"`
//}
