package main

import "sync"

type store struct {
	mutex   sync.Mutex
	records map[string][]byte
}

func newStore() *store {
	return &store{records: make(map[string][]byte)}
}

// chose to use global variable for simplicity's sake.
// in real life this would be a service reference to a resource that
// handles the transactional semantics
var kvStore *store = newStore()

func (st *store) add(key string, value []byte) bool {
	st.mutex.Lock()
	defer st.mutex.Unlock()
	if _, ok := st.records[key]; ok {
		return false
	}

	st.records[key] = value

	return true
}

func (st *store) get(key string) ([]byte, bool) {
	st.mutex.Lock()
	defer st.mutex.Unlock()
	value, ok := st.records[key]
	if !ok {
		return nil, false
	}
	return value, ok
}
