package api

import (
	"net/http"

	"orbit/internal/queue"
	"orbit/internal/server"
	"orbit/internal/session"
)

func RegisterRoutes(mux *http.ServeMux, q *queue.Queue, s *session.Store, pool *server.Pool) {
	mux.Handle("/queue", queueHandler(q))
	mux.Handle("/session/", sessionHandler(s))
	mux.Handle("/server/register", serverRegisterHandler(pool))
	mux.Handle("/server/heartbeat", serverHeartbeatHandler(pool))
}
