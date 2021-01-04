module github.com/ALiuGuanyan/margin

go 1.15

replace github.com/ALiuGuanyan/micro-boot => ../micro-boot

require (
	github.com/ALiuGuanyan/micro-boot v0.0.0-00010101000000-000000000000
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/consul/api v1.8.1
	github.com/imdario/mergo v0.3.11
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.9.0
	github.com/sony/gobreaker v0.4.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d
	google.golang.org/grpc v1.34.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)
