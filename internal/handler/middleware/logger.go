package middleware

import (
	"context"
	"net/http"
	"time"

	"folderstructure/platform/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextkey string

const (
	StartTimeKey contextkey = "x-start-time"
	UserAgentKey contextkey = "user-agent"
	IPKey        contextkey = "ip"
)

func Logger(log logger.Logger) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ctx := context.WithValue(r.Context(), StartTimeKey, start)
			ctx = context.WithValue(ctx, UserAgentKey, r.UserAgent())
			ctx = context.WithValue(ctx, IPKey, r.RemoteAddr)

			r = r.WithContext(ctx)
			ww := NewWrapResponseWriter(w)
			next.ServeHTTP(ww, r)
			latency := time.Since(start)
			fields := []zapcore.Field{
				zap.Int("status", ww.Status()),
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("ip", r.RemoteAddr),
				zap.Any("query", r.URL.Query()),
				zap.String("user-agent", r.UserAgent()),
				zap.Float64("latency_ms", float64(latency.Milliseconds())),
			}
			log.Info(context.Background(), "Gin Request", fields...)
		})
	}
}
