package main

import (
	"github.com/bugbundle/phantom/internal/adapter/http"
)

func main() {
	svc := http.NewService("0.0.0.0:8080")
	svc.Run()
}
