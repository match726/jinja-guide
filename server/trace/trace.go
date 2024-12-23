package trace

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracer() error {

	tp := sdkTrace.NewTracerProvider(
		sdkTrace.WithSampler(sdkTrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return nil

}

func ShowTraceID(ctx context.Context) {

	tracer := otel.Tracer("backend-tracer")

	ctx, span := tracer.Start(ctx, "some operation")
	traceID := span.SpanContext().TraceID().String()
	fmt.Println("Trace ID is ", traceID)
}
