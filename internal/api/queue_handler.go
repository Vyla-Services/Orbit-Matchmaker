package api

import (
	"encoding/json"
	"net/http"

	"internal/queue"
	"pkg/models"
)

func queueHandler(q *queue.Queue) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}

		var p models.Player
		if json.NewDecoder(r.Body).Decode(&p) != nil {
			w.WriteHeader(400)
			return
		}

		q.Add(p)
		w.WriteHeader(200)
	})
}
