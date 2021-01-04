package consul

// HealthCheck is used to represent a single check
type HealthCheck struct {
	Node        string
	CheckID     string
	Name        string
	Status      string
	Notes       string
	Output      string
	ServiceID   string
	ServiceName string
	ServiceTags []string
	Type        string
	Namespace   string `json:",omitempty"`

	Definition HealthCheckDefinition

	CreateIndex uint64
	ModifyIndex uint64
}
