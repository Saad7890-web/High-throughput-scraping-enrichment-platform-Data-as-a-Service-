package server

import (
	"net/http"

	"github.com/Saad7890-web/scrapper-platform/internal/health"
	"github.com/Saad7890-web/scrapper-platform/internal/job"
)

func New(jobHandler *job.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", health.Handler)

	// Job routes

	mux.HandleFunc("/jobs", jobHandler.Create) 

	return mux
}
