package main

import (
	"net/http"
	"fmt"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		hitsMsg := fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())
		w.Write([]byte(hitsMsg))
	}

	
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}