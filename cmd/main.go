package main

import (
	"github.com/bugbundle/phantom/internal/adapter/http"
)

func main() {
	http.Server("0.0.0.0:8080")
}
