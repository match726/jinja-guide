package trace

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracerProvider() (func(context.Context) error, error) {

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tp.Shutdown, nil

}

func GetContextWithTraceID(ctx context.Context, spanName string) context.Context {

	_, span := otel.Tracer("shrine-guide").Start(ctx, spanName)
	defer span.End()

	fmt.Printf("traceID: %s\n", span.SpanContext().TraceID().String())

	return context.WithValue(ctx, "traceID", span.SpanContext().TraceID().String())

}
