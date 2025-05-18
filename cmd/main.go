package main

import (
	"github.com/bugbundle/phantom/internal/adapter/http"
)

func main() {
	svc := http.NewService("127.0.0.1:8080")
	svc.Run()
}
