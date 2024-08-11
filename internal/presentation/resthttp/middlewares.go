package resthttp

import (
	"net/http"
	"time"

	"github.com/golanglowell/quick-link/pkg/logger"
)

func loggingMiddleware(logger *logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		logger.Info("Request processed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", duration,
			"remote_addr", r.RemoteAddr,
		)
	})
}
