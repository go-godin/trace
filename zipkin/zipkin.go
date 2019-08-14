package zipkin

import (
	"github.com/go-godin/log"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	"github.com/openzipkin/zipkin-go/reporter/http"
)

func New(serviceName, backendAddress, hostPort string, logger log.Logger) (*zipkin.Tracer, error) {
	var rep reporter.Reporter
	if backendAddress == "" {
		rep = reporter.NewNoopReporter()
		logger.Warning("missing zipkin reporter address: TRACING WILL BE DISABLED")
	} else {
		rep = http.NewReporter(backendAddress)
	}
	localEndpoint, err := zipkin.NewEndpoint(serviceName, hostPort)
	if err != nil {
		return nil, err
	}
	tracer, err := zipkin.NewTracer(rep, zipkin.WithLocalEndpoint(localEndpoint))
	if err != nil {
		return nil, err
	}

	return tracer, nil
}
