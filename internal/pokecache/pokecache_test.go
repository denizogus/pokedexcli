package pokecache

import (
	"bytes"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewCache()
	expected := Cache{
		Data: map[string][]byte{
			"test": []byte("test"),
		},
	}

	cache.Add("test", []byte("test"))

	data, ok := cache.Get("test")
	if !ok {
		t.Fatal("couldnt get from the cache")
	}
	if !bytes.Equal(data, expected.Data["test"]) {
		t.Fatalf("expected: %v, got: %v", expected.Data["test"], data)
	}
}
