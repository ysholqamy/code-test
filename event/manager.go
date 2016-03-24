package event

import (
	"fmt"
	"sync"
)

// EventDecoder should be holding a stream of the event. Decode decodes the stream.
type EventDecoder interface {
	Decode(v interface{}) error
}

type Manager struct {
	sync.Mutex
	eventsStore map[string]*eventsData //map[sessionId]*eventsData
}

func NewManager() *Manager {
	return &Manager{eventsStore: make(map[string]*eventsData, 0)}
}

//findOrCreate should find the eventsData with the gived sid. a new eventsData will be created if not found.
func (m *Manager) findOrCreate(sid, url string) *eventsData {
	m.Lock()
	defer m.Unlock()

	ev, ok := m.eventsStore[sid]

	if !ok {
		ev = newEventsData(url)
		m.eventsStore[sid] = ev
	}

	return ev
}

// RegisterEvent takes an EventDecoder holding a stream describing the event and populates the corresponding eventsData
func (m *Manager) RegisterEvent(decoder EventDecoder, ch chan Result) {
	m.Lock()
	defer m.Unlock()

	var event map[string]interface{}

	err := decoder.Decode(&event)
	if err != nil {
		ch <- Result{"", err}
		return
	}

	sid, ok := event["sessionId"].(string)

	if !ok {
		err := fmt.Errorf("Missing sessionId in decoded event %+v", event)
		ch <- Result{"", err}
		return
	}

	url := event["websiteUrl"].(string)
	ev := m.findOrCreate(sid, url)

	ch <- ev.populate(event, sid)
}
