package logger

import (
	"fmt"
	"net/http"
	"time"
)

// Needs a goroutine to process the log entries
var LogChannel = make(chan string, 100)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		duration := float32(time.Since(start).Nanoseconds()/1000) / 1e3
		logEntry := fmt.Sprintf("%s %s %d %.3fms\n", r.Method, r.URL.Path, rw.statusCode, duration)
		LogChannel <- logEntry
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
