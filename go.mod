module github.com/ALiuGuanyan/margin

go 1.14

replace github.com/ALiuGuanyan/micro-boot => ../micro-boot

require (
	github.com/ALiuGuanyan/micro-boot v0.0.0-00010101000000-000000000000
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/consul/api v1.8.1
	github.com/imdario/mergo v0.3.11
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5 // indirect
	github.com/openzipkin/zipkin-go v0.2.5 // indirect
	github.com/prometheus/client_golang v1.8.0
	github.com/sony/gobreaker v0.4.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
	google.golang.org/genproto v0.0.0-20201210142538-e3217bee35cc
	google.golang.org/grpc v1.34.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)
