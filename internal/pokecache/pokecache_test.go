package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(100 * time.Millisecond)

	cache.Add("key1", []byte("value1"))

	val, ok := cache.Get("key1")
	if !ok {
		t.Errorf("Expected to find key1 in cache")
	} else if string(val) != "value1" {
		t.Errorf("Expected value 'value1' but got '%s'", string(val))
	}

	time.Sleep(150 * time.Millisecond)

	_, ok = cache.Get("key1")
	if ok {
		t.Errorf("Expected key1 to be reaped from cache")
	}
}
