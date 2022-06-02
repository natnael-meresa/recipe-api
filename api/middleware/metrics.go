package middleware

import (
	"bufio"
	"errors"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/natnael-meresa/recipe/pkg/metric"
)

type MetricsMiddleware struct {
	metricService metric.Service
}

func NewMetricsMiddleware() *MetricsMiddleware {
	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}

	return &MetricsMiddleware{
		metricService: metricService,
	}
}

func (mt MetricsMiddleware) Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appMetric := metric.NewHTTP(r.URL.Path, r.Method)
		appMetric.Started()
		rwi := &responseWriterInterceptor{
			statusCode:     http.StatusOK,
			ResponseWriter: w,
		}

		next.ServeHTTP(rwi, r)

		appMetric.Finished()
		appMetric.StatusCode = strconv.Itoa(rwi.statusCode)
		mt.metricService.SaveHTTP(appMetric)

	})
}

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterInterceptor) Write(p []byte) (int, error) {
	return w.ResponseWriter.Write(p)
}

func (w *responseWriterInterceptor) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("type assertion failed http.ResponseWriter not a http.Hijacker")
	}
	return h.Hijack()
}

func (w *responseWriterInterceptor) Flush() {
	f, ok := w.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}

	f.Flush()
}
