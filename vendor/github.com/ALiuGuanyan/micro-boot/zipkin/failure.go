package zipkin

import (
	"github.com/openzipkin/zipkin-go"
)


var defaultTracerExtractFailurePolicy = zipkin.ExtractFailurePolicyRestart

func SetDefaultTracerExtractFailurePolicy(policy zipkin.ExtractFailurePolicy)  {
	defaultTracerExtractFailurePolicy = policy
}
