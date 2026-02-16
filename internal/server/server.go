package server

import (
	"net/http"

	"github.com/Saad7890-web/scrapper-platform/internal/health"
)

func New() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.Handler)
	return mux
}