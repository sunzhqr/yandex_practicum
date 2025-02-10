package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("cannot initialize zap")
	}
	defer logger.Sync()

	const url = "http://example.com"

	// SugaredLogger is a structured logger
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL", "url", url)
	sugar.Infof("Failed to fetch URL: %s", url)
	sugar.Errorf("Failed to fetch URL: %s", url)

	// Converts Sugarred Logger to plain Logger
	plain := sugar.Desugar()
	plain.Info("Hello, from plain logger")
	plain.Warn("Simple warning")
	plain.Error("Failed to fetdch URL", zap.String("url", url))
}
