package queue

import (
	"pkg/models"
	"sync"
)

type Queue struct {
	mu      sync.Mutex
	players []models.Player
}

func New() *Queue {
	return &Queue{
		players: make([]models.Player, 0),
	}
}

func (q *Queue) Add(p models.Player) {
	q.mu.Lock()
	q.players = append(q.players, p)
	q.mu.Unlock()
}

func (q *Queue) PopBatch(playlist, region string, size int) []models.Player {
	q.mu.Lock()
	defer q.mu.Unlock()

	batch := make([]models.Player, 0, size)
	remaining := make([]models.Player, 0, len(q.players))

	for _, p := range q.players {
		if p.Playlist == playlist && p.Region == region && len(batch) < size {
			batch = append(batch, p)
		} else {
			remaining = append(remaining, p)
		}
	}

	q.players = remaining
	return batch
}
