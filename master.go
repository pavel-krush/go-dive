package dive

import (
	"sync"
)

type Master struct {
	mu         sync.RWMutex
	activities map[string]*Activity
}

func NewMaster() *Master {
	return &Master{
		activities: make(map[string]*Activity),
	}
}

func (m *Master) AddActivity(id string, activity *Activity) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.activities[id] = activity
}

func (m *Master) RemoveActivity(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.activities, id)
}

func (m *Master) GetActivity(id string) *Activity {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.activities[id]
}

func (m *Master) WalkActivities(cb func(activity *Activity)) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, activity := range m.activities {
		cb(activity)
	}
}
