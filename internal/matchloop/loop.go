package matchloop

import (
	"internal/playlist"
	"internal/queue"
	"internal/server"
	"internal/session"
	"pkg/models"
	"time"
)

type Loop struct {
    queue   *queue.Queue
    store   *session.Store
    servers *server.Pool
}

func New(q *queue.Queue, s *session.Store, pool *server.Pool) *Loop {
    return &Loop{
        queue:   q,
        store:   s,
        servers: pool,
    }
}

func (l *Loop) Run(interval time.Duration) {
    for {
        l.tick()
        time.Sleep(interval)
    }
}

func (l *Loop) tick() {
    for _, pl := range playlist.Playlists() {
        for _, region := range playlist.Regions() {
            l.match(pl, region)
        }
    }
}

func (l *Loop) match(pl models.Playlist, region string) {
    batch := l.queue.PopBatch(pl.Name, region, pl.MaxPlayers)
    if len(batch) == 0 {
        return
    }

    srv := l.servers.Pick(region)
    if srv == nil {
        return
    }

    id := time.Now().Format("20060102150405")
    players := make([]string, 0, len(batch))
    for _, p := range batch {
        players = append(players, p.AccountID)
    }

    sess := models.Session{
        ID:       id,
        Playlist: pl.Name,
        Region:   region,
        Players:  players,
        ServerIP: srv.IP,
    }

    l.store.Save(sess)
    l.servers.IncrementLoad(srv.ID, len(players))
}
