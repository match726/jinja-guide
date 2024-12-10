package logger

import (
	"context"
	"fmt"
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
		fmt.Println(sc)
		return logger
	}

	return logger.With(
		slog.String("traceID", sc.TraceID().String()),
		slog.String("spanID", sc.SpanID().String()),
	)

}

func WriteInfo(ctx context.Context, msg string, args ...any) {

	logger := newLogger(ctx)

	logger.Info(msg, args)

}
