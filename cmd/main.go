package main

import (
	"log/slog"
	"os"

	"github.com/kevinharv/idp/internal/server"
)

/*
	- read in configuration
		- web
		- ad
		- idp
	- create OAuth and OIDC endpoints
	- create SAML endpoints
	- register middleware
*/

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)
	slog.Info("Starting IdP")
	server.Serve()
}
