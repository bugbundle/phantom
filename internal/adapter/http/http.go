package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bugbundle/phantom/internal/adapter/http/controls"
	"github.com/bugbundle/phantom/internal/adapter/http/front"
	"github.com/bugbundle/phantom/internal/adapter/logger"
)

type service struct {
	config config
}

type config struct {
	addr string
}

func panicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Print(fmt.Errorf("an error occured:\n%s", err))

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"error\": \"Internal Server Error\"}"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func NewService(addr string) *service {
	return &service{
		config: config{
			addr: addr,
		},
	}
}

func (svc *service) Run() error {
	router := http.NewServeMux()
	controls.RegisterRoutes(router)
	front.RegisterRoutes(router)

	serverConfig := &http.Server{
		Addr: svc.config.addr,
		Handler: logger.LoggingHandler(
			panicHandler(router),
		),
	}
	if err := serverConfig.ListenAndServe(); err != nil {
		return fmt.Errorf("an error occured while starting the server.\n%w", err)
	}
	return nil
}
