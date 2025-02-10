package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

var sugar zap.SugaredLogger

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar = *(logger.Sugar())

	// http.Handle("/ping", http.HandlerFunc(pingHandler1))
	http.Handle("/ping", pingHandler())

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
