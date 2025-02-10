package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

type (
	responseData struct {
		status int
		size   int
	}

	logginResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (l *logginResponseWriter) Write(b []byte) (int, error) {
	size, err := l.ResponseWriter.Write(b)
	l.responseData.size += size
	return size, err
}

func (l *logginResponseWriter) WriteHeader(statusCode int) {
	l.ResponseWriter.WriteHeader(statusCode)
	l.responseData.status = statusCode
}

var sugar zap.SugaredLogger

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar = *(logger.Sugar())

	// http.Handle("/ping", http.HandlerFunc(pingHandler1))
	http.Handle("/ping", WithLogging(pingHandler()))

	address := "127.0.0.1:8080"

	sugar.Infow(color.GreenString("Starting server"), "addr", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		sugar.Fatalw(color.RedString(err.Error()), "event", "starting server")
	}
}

// func pingHandler1(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("pong\n"))
// }

func pingHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "pong\n")
	}
	return http.HandlerFunc(fn)
}

func WithLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := logginResponseWriter{
			responseData:   responseData,
			ResponseWriter: w,
		}

		uri := r.RequestURI
		method := r.Method
		h.ServeHTTP(&lw, r)
		duration := time.Since(start)

		sugar.Infoln(
			"uri", uri,
			"method", method,
			"status", responseData.status,
			"duration", duration,
			"size", responseData.size,
		)
	}
	return http.HandlerFunc(logFn)
}
