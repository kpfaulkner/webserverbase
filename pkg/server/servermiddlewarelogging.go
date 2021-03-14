package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
	return
}

func WithLogging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := wrapResponseWriter(w)
			next.ServeHTTP(rw, r)
			end := time.Now()

			log.Info(fmt.Sprintf("%s took %d ms : resp code %d", r.URL.EscapedPath(), end.Sub(start).Milliseconds(), rw.status))
		}
		return http.HandlerFunc(fn)
	}
}
