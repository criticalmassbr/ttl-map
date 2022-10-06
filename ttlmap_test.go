package ttlmap

import (
	"testing"
	"time"
)

func TestTTLMap(t *testing.T) {
	t.Run("can add values over the initial len", func(t *testing.T) {
		ttl := New[string](0, 10)
		for i := 0; i < 100; i++ {
			ttl.Put(i, "meh")
		}
	})

	t.Run("returns the correct len", func(t *testing.T) {
		ttl := New[string](time.Second, time.Second)

		if ttl.Len() != 0 {
			t.Fail()
		}

		for i := 0; i < 100; i++ {
			ttl.Put(i, "meh")
		}

		if ttl.Len() != 100 {
			t.Fail()
		}
	})

	t.Run("overides the values on put with same key", func(t *testing.T) {
		ttl := New[string](time.Second, time.Second)

		ttl.Put(1, "oldval")
		ttl.Put(1, "newval")

		if ttl.Get(1) != "newval" {
			t.Fail()
		}
	})

	t.Run("clears values with expired ttl on every tick", func(t *testing.T) {
		m := New[string](time.Second, time.Second)

		m.Put(1, "meh")

		// we need to sleep 2 seconds to wait for the second tick to trigger
		time.Sleep(time.Second * 3)

		if m.Get(1) != "" {
			t.Fail()
		}
	})
}
