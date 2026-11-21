package main

import (
	"log/slog"
	"os"
)

func main() {
	server := newHTTPServer()

	slog.Info("starting HTTP server", "addr", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("HTTP server stopped", "error", err)
		os.Exit(1)
	}
}
