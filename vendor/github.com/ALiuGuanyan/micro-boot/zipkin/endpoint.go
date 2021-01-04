package zipkin

import (
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
)

// LocalEndpoint is the settings for creating a standard Zipkin local endpoint
type LocalEndpoint struct {
	// Name is the service name which used to create a local endpoint
	Name string `json:"name" yaml:"name"`

	// HostPort is the host port which used to create a local endpoint
	HostPort string `json:"address" yaml:"address"`
}

func (le LocalEndpoint) Standardize() (*model.Endpoint, error) {
	return zipkin.NewEndpoint(le.Name, le.HostPort)
}
