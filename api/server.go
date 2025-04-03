package api

import (
	"net/http"
	"os"

	"github.com/bugbundle/phantom/api/middlewares"
	"github.com/bugbundle/phantom/api/routes"
	"github.com/charmbracelet/log"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Server(addr string) {
	logger := log.NewWithOptions(os.Stderr, log.Options{Prefix: "phantom/http"})
	stdlog := logger.StandardLog(log.StandardLogOptions{
		ForceLevel: log.ErrorLevel,
	})

	server_config := &http.Server{
		Addr:     addr,
		Handler:  otelhttp.NewHandler(middlewares.ServerLogHandler(http.DefaultServeMux), ""),
		ErrorLog: stdlog,
	}

	http.HandleFunc("/", routes.Homepage)
	http.HandleFunc("POST /cameras", routes.CreateCamera)
	http.HandleFunc("DELETE /cameras", routes.DeleteCamera)
	http.HandleFunc("GET /cameras", routes.StreamVideo)

	log.Info("Starting server...", "interface", addr)
	if err := server_config.ListenAndServe(); err != nil {
		log.Fatal("An error occured !", "interface", addr, "error", err)
	}
}
