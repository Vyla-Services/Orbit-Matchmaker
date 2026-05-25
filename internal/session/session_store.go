package session

import (
	"pkg/models"
	"sync"
)

type Store struct {
	mu       sync.RWMutex
	sessions map[string]models.Session
}

func NewStore() *Store {
	return &Store{
		sessions: make(map[string]models.Session),
	}
}

func (s *Store) Save(sess models.Session) {
	s.mu.Lock()
	s.sessions[sess.ID] = sess
	s.mu.Unlock()
}

func (s *Store) FindByPlayer(accountID string) *models.Session {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, sess := range s.sessions {
		for _, p := range sess.Players {
			if p == accountID {
				cp := sess
				return &cp
			}
		}
	}
	return nil
}
