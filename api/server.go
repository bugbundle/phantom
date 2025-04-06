package api

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/bugbundle/phantom/api/middlewares"
	"github.com/bugbundle/phantom/api/routes"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Server(addr string) {
	// Configure default logging to JSON
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	router := http.NewServeMux()

	router.HandleFunc("/", routes.Homepage)
	router.HandleFunc("POST /cameras", routes.CreateCamera)
	router.HandleFunc("GET /cameras/status", routes.StreamStatus)
	router.HandleFunc("DELETE /cameras", routes.DeleteCamera)
	router.HandleFunc("GET /cameras", routes.StreamVideo)

	server_config := &http.Server{
		Addr: addr,
		Handler: otelhttp.NewHandler(
			middlewares.LoggingHandler(
				middlewares.Recovery(router),
			),
			"",
		),
	}

	slog.Info("Starting server...", "interface", addr)
	if err := server_config.ListenAndServe(); err != nil {
		slog.Error("An error occured !", "interface", addr, "error", err)
	}
}
