package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nunseik/go-server/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries    *database.Queries
	platform     string
	secretKey    string
	polkaKey     string
}



func main() {
	const port = "8080"

	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("error opening sql db: %s", err)
	}

	platform := os.Getenv("PLATFORM")
	secretKey := os.Getenv("SECRET_KEY")
	polkaKey := os.Getenv("POLKA_KEY")
	dbQueries := database.New(db)


	apiCfg := apiConfig{
		dbQueries: dbQueries,
		platform:  platform,
		secretKey: secretKey,
		polkaKey:  polkaKey,
	}
	handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))

	
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /api/chirps", apiCfg.handlerChirpsCreate)
	mux.HandleFunc("POST /api/users", apiCfg.handlerUserCreation)
	mux.HandleFunc("PUT /api/users", apiCfg.handlerUserUpdate)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)
	mux.HandleFunc("GET /api/chirps", apiCfg.handlerChirpsRetrieve)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handlerChirpsGet)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.handlerChirpsDelete)
	mux.HandleFunc("POST /api/refresh", apiCfg.handlerTokenRefresh)
	mux.HandleFunc("POST /api/revoke", apiCfg.handlerTokenRevoke)
	mux.HandleFunc("POST /api/polka/webhooks", apiCfg.handlerPolkaWebhooks)

	srv := &http.Server{Handler: mux, Addr: ":" + port}

	log.Printf("Serving files on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}