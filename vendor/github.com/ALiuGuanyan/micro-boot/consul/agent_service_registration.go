package consul

import (
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"github.com/hashicorp/consul/api"
)

var (
	defaultAgentServiceRegistrationFlagsPrefix = defaultConsulFlagsPrefix + "-agent"

	defaultKind = ""
	defaultAgentID = ""
	defaultAgentName = ""
	defaultAgentAddress = ""
	defaultAgentPort = 0
	defaultAgentTags = []string{}

	defaultAgentMeta = map[string]string{}

	defaultAgentNamespace = ""
	defaultAgentEnableTagOverride = false
)

func SetDefaultAgentPort(val int)  {
	defaultAgentPort = val
}

func SetDefaultAgentKind(val string)  {
	defaultKind = val
}

func SetDefaultAgentAddress(val string)  {
	defaultAgentAddress = val
}

func SetDefaultAgentID(val string)  {
	defaultAgentID = val
}

func SetDefaultAgentName(val string)  {
	defaultAgentName = val
}

func SetDefaultAgentNamespace(val string)  {
	defaultAgentNamespace = val
}

func SetDefaultAgentTags(val []string)  {
	defaultAgentTags = val
}

func SetDefaultAgentMeta(val map[string]string)  {
	defaultAgentMeta = val
}

func SetDefaultAgentEnableTagOverride(val bool)  {
	defaultAgentEnableTagOverride = val
}

type AgentServiceRegistration struct {
	Kind              string               `json:"kind,omitempty" yaml:"kind,omitempty"`
	ID                string               `json:"id,omitempty" yaml:"id,omitempty"`
	Name              string               `json:"name,omitempty" yaml:"name,omitempty"`
	Tags              []string             `json:"tags,omitempty" yaml:"tags,omitempty"`
	Port              int                  `json:"port,omitempty" yaml:"port,omitempty"`
	Address           string               `json:"address,omitempty" yaml:"address,omitempty"`
	EnableTagOverride bool                 `json:"enable-tag-override,omitempty" yaml:"enable-tag-override.omitempty"`
	Meta              map[string]string    `json:"meta,omitempty" yaml:"meta,omitempty"`

	Namespace         string               `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (asr *AgentServiceRegistration) Parse() (err error) {
	if asr.CustomParseFunc != nil {
		return asr.CustomParseFunc()
	}

	return nil
}

func (asr *AgentServiceRegistration) BindFlags(fs *bootflag.FlagSet)  {
	if asr.CustomBindFlagsFunc != nil {
		asr.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&asr.Kind,
		utils.BuildFlagName(defaultAgentServiceRegistrationFlagsPrefix, "kind"),
		defaultKind,
		"the kind of service being registered")

	fs.StringVar(
		&asr.ID,
		utils.BuildFlagName(defaultAgentServiceRegistrationFlagsPrefix, "id"),
		defaultAgentID,
		"the id of service being registered")

	fs.StringVar(
		&asr.Name,
		utils.BuildFlagName(defaultAgentServiceRegistrationFlagsPrefix, "name"),
		defaultAgentName,
		"the name of service being registered")

	fs.StringSliceVar(
		&asr.Tags,
		utils.BuildFlagName(defaultAgentServiceRegistrationFlagsPrefix, "tags"),
		defaultAgentTags,
		"the tags of service being registered")

	fs.IPv4PortIntVar(
		&asr.Port,
		utils.BuildFlagName(defaultAgentServiceRegistrationFlagsPrefix, "port"),
		defaultAgentPort,
		"the port of service being registered")

	fs.StringVar(
		&asr.Address,
		utils.BuildFlagName(defaultAgentServiceRegistrationFlagsPrefix, "address"),
		defaultAgentAddress,
		"the address of service being registered")

	fs.BoolVar(
		&asr.EnableTagOverride,
		utils.BuildFlagName(defaultAgentServiceRegistrationFlagsPrefix, "enable-tag-override"),
		defaultAgentEnableTagOverride,
		"the enable tag override of service being registered")

	fs.StringToStringVar(
		&asr.Meta,
		utils.BuildFlagName(defaultAgentServiceRegistrationFlagsPrefix, "meta"),
		defaultAgentMeta,
		"the meta of service being registered")

	fs.StringVar(
		&asr.Namespace,
		utils.BuildFlagName(defaultAgentServiceRegistrationFlagsPrefix, "namespace"),
		defaultAgentNamespace,
		"the namespace of service being registered")
}

func (asr *AgentServiceRegistration) Standardize() *api.AgentServiceRegistration {
	return &api.AgentServiceRegistration{
		Kind:              api.ServiceKind(asr.Kind),
		ID:                asr.ID,
		Name:              asr.Name,
		Tags:              asr.Tags,
		Port:              asr.Port,
		Address:           asr.Address,
		EnableTagOverride: asr.EnableTagOverride,
		Meta:              asr.Meta,
		Namespace:         asr.Namespace,
	}
}
