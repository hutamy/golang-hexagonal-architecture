package util

import (
	"context"

	"github.com/hutamy/golang-hexagonal-architecture/config"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Trace(ctx context.Context, operationName, resourceName string) (ddtrace.Span, context.Context) {
	configuration := config.GetConfig()
	span, returnedCtx := tracer.StartSpanFromContext(
		ctx, operationName, tracer.ServiceName(configuration.DDService), tracer.ResourceName(resourceName),
	)

	return span, returnedCtx
}
