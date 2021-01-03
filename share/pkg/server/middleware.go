package server

import (
	"github.com/gorilla/mux"
	stdlog "log"
	"net"
	"net/http"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
)

func ListenAndServe(r *mux.Router, ln net.Listener, endpoint string)  {
	mw := nethttp.Middleware(
		opentracing.GlobalTracer(),
		r,
		nethttp.OperationNameFunc(func(r *http.Request) string {
			return "HTTP " + r.Method + ":" + endpoint
		}),
	)

	stdlog.Fatal(http.Serve(ln, mw))
}
