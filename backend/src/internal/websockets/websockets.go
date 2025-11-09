package websockets

import (
	"sync"

	"github.com/CZnavody19/music-manager/src/graph/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Websockets struct {
	subscribers map[uuid.UUID]chan *model.Task
	mu          sync.Mutex
}

func NewWebsockets() (*Websockets, error) {
	return &Websockets{
		subscribers: make(map[uuid.UUID]chan *model.Task),
	}, nil
}

func (w *Websockets) AddSubscriber(id uuid.UUID) chan *model.Task {
	w.mu.Lock()
	defer w.mu.Unlock()

	channel := make(chan *model.Task)

	w.subscribers[id] = channel

	return channel
}

func (w *Websockets) RemoveSubscriber(id uuid.UUID) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if channel, exists := w.subscribers[id]; exists {
		close(channel)
		delete(w.subscribers, id)
	}
}

func (w *Websockets) SendTask(task *model.Task) {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, sub := range w.subscribers {
		select {
		case sub <- task:
			zap.S().Info("Task sent")
		default:
			zap.S().Info("Task channel is full, skipping send")
		}
	}
}
