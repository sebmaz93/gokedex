package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "http://seb.com",
			val: []byte("example data"),
		},
		{
			key: "http://poken.com/pokedex",
			val: []byte("pokemon data"),
		},
	}

	for idx, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", idx), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReap(t *testing.T) {
	const initTime = 5 * time.Millisecond
	const waitTime = initTime + 5*time.Millisecond
	cache := NewCache(initTime)
	cache.Add("http://poke.com", []byte("pokkeemon"))

	_, ok := cache.Get("http://poke.com")
	if !ok {
		t.Errorf("expted to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("http://poke.com")
	if ok {
		t.Errorf("expted to not find key")
		return
	}
}
