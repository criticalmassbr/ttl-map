package ttlmap

import (
	"sync"
	"time"
)

type item[T any] struct {
	value      T
	lastAccess int64
}

type TTLMap[T any] struct {
	m map[int]*item[T]
	l sync.Mutex
}

// ttl: max time in seconds every value can live
// tickEvery: interval in seconds all values with expired ttl should be deleted
//
// IMPORTANT: ttl and tickEvery must be at least one second
func New[T any](ttl, tickEvery time.Duration) *TTLMap[T] {
	m := &TTLMap[T]{m: map[int]*item[T]{}}

	go func() {
		for now := range time.Tick(tickEvery) {
			m.l.Lock()

			for k, v := range m.m {
				if now.Unix()-v.lastAccess > int64(ttl.Seconds()) {
					delete(m.m, k)
				}
			}

			m.l.Unlock()
		}
	}()

	return m
}

func (m *TTLMap[T]) Len() int {
	return len(m.m)
}

func (m *TTLMap[T]) Put(k int, v T) {
	m.l.Lock()

	it, ok := m.m[k]

	if !ok {
		it = &item[T]{value: v}
		m.m[k] = it
	}

	it.value = v
	it.lastAccess = time.Now().Unix()

	m.l.Unlock()
}

func (m *TTLMap[T]) Get(k int) (v T) {
	m.l.Lock()

	if it, ok := m.m[k]; ok {
		v = it.value
		it.lastAccess = time.Now().Unix()
	}

	m.l.Unlock()

	return
}
