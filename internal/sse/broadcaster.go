package sse

import (
	"sync"
)

// Event is a real-time SSE event.
type Event struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

// Broadcaster manages SSE subscribers for a given project.
type Broadcaster struct {
	mu          sync.RWMutex
	subscribers map[string][]chan Event
}

// NewBroadcaster creates a new Broadcaster.
func NewBroadcaster() *Broadcaster {
	return &Broadcaster{subscribers: make(map[string][]chan Event)}
}

// Subscribe registers a channel for events on a project.
func (b *Broadcaster) Subscribe(projectID string) chan Event {
	ch := make(chan Event, 10)
	b.mu.Lock()
	b.subscribers[projectID] = append(b.subscribers[projectID], ch)
	b.mu.Unlock()
	return ch
}

// Publish sends an event to all subscribers of a project.
func (b *Broadcaster) Publish(projectID string, event Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, ch := range b.subscribers[projectID] {
		select {
		case ch <- event:
		default:
		}
	}
}
