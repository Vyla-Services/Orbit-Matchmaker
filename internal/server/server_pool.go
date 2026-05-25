package server

import (
	"pkg/models"
	"sync"
	"time"
)

type Pool struct {
	mu      sync.Mutex
	servers map[string]models.Server
}

func NewPool() *Pool {
	return &Pool{
		servers: make(map[string]models.Server),
	}
}

func (p *Pool) Register(s models.Server) {
	p.mu.Lock()
	s.LastHeartbeat = time.Now().Unix()
	p.servers[s.ID] = s
	p.mu.Unlock()
}

func (p *Pool) Heartbeat(id string) {
	p.mu.Lock()
	s, ok := p.servers[id]
	if ok {
		s.LastHeartbeat = time.Now().Unix()
		p.servers[id] = s
	}
	p.mu.Unlock()
}

func (p *Pool) Pick(region string) *models.Server {
	p.mu.Lock()
	defer p.mu.Unlock()

	var best *models.Server
	for _, s := range p.servers {
		if s.Region != region {
			continue
		}
		if best == nil || s.CurrentLoad < best.CurrentLoad {
			cp := s
			best = &cp
		}
	}
	return best
}

func (p *Pool) IncrementLoad(id string, delta int) {
	p.mu.Lock()
	s, ok := p.servers[id]
	if ok {
		s.CurrentLoad += delta
		p.servers[id] = s
	}
	p.mu.Unlock()
}
