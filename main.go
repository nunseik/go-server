package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}



func main() {
	const port = "8080"
	apiCfg := apiConfig{}
	handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))

	
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	
	srv := &http.Server{Handler: mux, Addr: ":" + port}

	log.Printf("Serving files on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}