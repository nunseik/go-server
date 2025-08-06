package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
		if cfg.platform != "dev" {
			w.WriteHeader(403)
			w.Write([]byte("403 Forbidden"))
			return
		}
		err := cfg.dbQueries.DeleteAllUsers(r.Context())
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("User not deleted"))
			return
		}
		cfg.fileserverHits.Store(0)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hits reset to 0"))
	}

