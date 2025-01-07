package logger

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"time"
)

// ハンドラーラッパー
type Handler struct {
	handler slog.Handler
}

var _ slog.Handler = &Handler{}

type contextKey string

var (
	fields contextKey = "slog_fields"
)

func NewHandler() *Handler {

	opt := slog.HandlerOptions{
		AddSource: true,
	}

	return &Handler{
		handler: slog.NewJSONHandler(os.Stdout, &opt),
	}

}

func (h *Handler) Handle(ctx context.Context, record slog.Record) error {

	if v, ok := ctx.Value(fields).(*sync.Map); ok {
		v.Range(func(key, val any) bool {
			if keyString, ok := key.(string); ok {
				record.AddAttrs(slog.Any(keyString, val))
			}
			return true
		})
	}

	return h.handler.Handle(ctx, record)

}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{h.handler.WithAttrs(attrs)}

}
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{h.handler.WithGroup(name)}
}

func WithValue(parent context.Context, key string, val any) context.Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if v, ok := parent.Value(fields).(*sync.Map); ok {
		mapCopy := copySyncMap(v)
		mapCopy.Store(key, val)
		return context.WithValue(parent, fields, mapCopy)
	}
	v := &sync.Map{}
	v.Store(key, val)
	return context.WithValue(parent, fields, v)
}

func copySyncMap(m *sync.Map) *sync.Map {
	var cp sync.Map
	m.Range(func(k, v interface{}) bool {
		cp.Store(k, v)
		return true
	})
	return &cp
}

func log(ctx context.Context, level slog.Level, msg string, args ...any) {

	slog.SetDefault(slog.New(NewHandler()))
	logger := slog.Default()
	if !logger.Enabled(ctx, level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])

	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(args...)

	_ = logger.Handler().Handle(ctx, r)

}

func Info(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelInfo, msg, args...)
}
