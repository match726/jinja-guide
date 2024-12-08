package logger

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

func newLogger(ctx context.Context) *slog.Logger {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))

	return setTrace(ctx, logger)

}

func setTrace(ctx context.Context, logger *slog.Logger) *slog.Logger {

	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	if !sc.IsValid() {
		return logger
	}

	return logger.With(
		slog.String("traceID", span.SpanContext().TraceID().String()),
		slog.String("spanID", span.SpanContext().SpanID().String()),
	)

}

func WriteInfo(msg string, args ...any) {

	logger := newLogger(context.Background())

	logger.Info(msg, args)

}
