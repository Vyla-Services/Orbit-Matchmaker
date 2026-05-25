package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"internal/api"
	"internal/matchloop"
	"internal/queue"
	"internal/server"
	"internal/session"
	"orbit/config"
)

func main() {
	cfg := config.Load()

	q := queue.New()
	sessions := session.NewStore()
	servers := server.NewPool()
	loop := matchloop.New(q, sessions, servers)

	mux := http.NewServeMux()
	api.RegisterRoutes(mux, q, sessions, servers)

	go loop.Run(3 * time.Second)

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}

	log.Println("orbit listening on", cfg.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
		os.Exit(1)
	}
}
