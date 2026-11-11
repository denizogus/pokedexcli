package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(5 * time.Second)

	if _, ok := cache.Get("test"); ok {
		t.Fatal("getting keys that doesnt exist")
	}

	cache.Add("test", []byte("test"))

	data, ok := cache.Get("test")
	if !ok {
		t.Fatal("couldnt get from the cache")
	}
	if !bytes.Equal(data, []byte("test")) {
		t.Fatalf("expected: %v, got: %v", []byte("test"), data)
	}

	time.Sleep(5 * time.Second)

	if _, ok := cache.Get("test"); ok {
		t.Fatal("shouldve been deleted but its not.")
	}
}
