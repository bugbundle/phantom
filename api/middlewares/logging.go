package middlewares

import (
	"net/http"

	"github.com/charmbracelet/log"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func ServerLogHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		h.ServeHTTP(w, r)
		statusCode := lrw.statusCode
		log.Info(r.Method, "status", statusCode, "uri", r.URL.String())
	}
	return http.HandlerFunc(fn)
}
