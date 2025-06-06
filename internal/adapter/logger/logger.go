package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusResponseWriter := &statusResponseWriter{
			w,
			http.StatusOK,
		}
		h.ServeHTTP(statusResponseWriter, r)
		slog.LogAttrs(
			context.Background(),
			slog.LevelInfo,
			"Responded to request",
			slog.Int("status", statusResponseWriter.statusCode),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
		)
	})
}
