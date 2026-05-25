package api

import (
	"encoding/json"
	"net/http"

	"internal/server"
	"pkg/models"
)

func serverRegisterHandler(pool *server.Pool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}

		var s models.Server
		if json.NewDecoder(r.Body).Decode(&s) != nil {
			w.WriteHeader(400)
			return
		}

		pool.Register(s)
		w.WriteHeader(200)
	})
}

func serverHeartbeatHandler(pool *server.Pool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}

		var hb models.Server
		if json.NewDecoder(r.Body).Decode(&hb) != nil {
			w.WriteHeader(400)
			return
		}

		pool.Heartbeat(hb.ID)
		w.WriteHeader(200)
	})
}
