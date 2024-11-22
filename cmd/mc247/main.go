package main

import (
	"log/slog"
	"os"

	"github.com/interrrp/mc247/internal/mc247"
	"github.com/lmittmann/tint"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))

	mc247, err := mc247.New(logger)
	if err != nil {
		logger.Error("failed creating mc247", "err", err)
		os.Exit(1)
	}

	if err := mc247.Run(); err != nil {
		logger.Error("fatal error running mc247", "err", err)
		os.Exit(1)
	}
}
