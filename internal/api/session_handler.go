package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"orbit/internal/session"
)

func sessionHandler(s *session.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			w.WriteHeader(400)
			return
		}

		id := parts[2]
		sess := s.FindByPlayer(id)
		if sess == nil {
			w.WriteHeader(404)
			return
		}

		json.NewEncoder(w).Encode(sess)
	})
}
