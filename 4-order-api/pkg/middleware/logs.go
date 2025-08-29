package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)
		logrus.WithFields(logrus.Fields{
			"statusCode":    wrapper.StatusCode,
			"method":        r.Method,
			"path":          r.URL.Path,
			"executionTime": time.Since(start),
		}).Info("Basic info")
	})
}
