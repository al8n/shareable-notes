package consul

import "time"

// HealthCheckDefinition is used to store the details about
// a health check's execution.
type HealthCheckDefinition struct {
	HTTP                                   string
	Header                                 map[string][]string
	Method                                 string
	Body                                   string
	TLSSkipVerify                          bool
	TCP                                    string
	IntervalDuration                       time.Duration `json:"-"`
	TimeoutDuration                        time.Duration `json:"-"`
	DeregisterCriticalServiceAfterDuration time.Duration `json:"-"`
}
