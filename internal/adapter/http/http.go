package http

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/bugbundle/phantom/internal/adapter/logger"
	httpRoutes "github.com/bugbundle/phantom/internal/app/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				slog.Error(fmt.Sprint(err))

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"error\": \"Internal Server Error\"}"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func Server(addr string) {
	fmt.Println("here")
	router := http.NewServeMux()

	router.HandleFunc("/", httpRoutes.Homepage)
	router.HandleFunc("POST /cameras", httpRoutes.CreateCamera)
	router.HandleFunc("GET /cameras/status", httpRoutes.StreamStatus)
	router.HandleFunc("DELETE /cameras", httpRoutes.DeleteCamera)
	router.HandleFunc("GET /cameras", httpRoutes.StreamVideo)

	fmt.Println("hore")
	server_config := &http.Server{
		Addr: addr,
		Handler: logger.LoggingHandler(
			Recovery(router),
		),
	}

	slog.Info("Starting server...", "interface", addr)
	if err := server_config.ListenAndServe(); err != nil {
		slog.Error("An error occured !", "interface", addr, "error", err)
	}
}
