package pokecache

import (
	"testing"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

func TestCacheAdd(t *testing.T) {
	val := []byte("mock pokemon location data")
	newCache := NewCache(time.Duration(20 * time.Second))
	newCache.Add(baseUrl, val)
	if len(newCache.entries) == 0 {
		t.Error("Cache is empty")
	}
	if newCache.entries[baseUrl].createdAt.IsZero() {
		t.Error("CreatedAt is Zero, was not initialized")
	}
	if newCache.entries[baseUrl].val == nil {
		t.Error("Val is nil")
	}
}

func TestCacheGet(t *testing.T) {
	val := []byte("mock pokemon location data")
	newCache := NewCache(time.Duration(20 * time.Second))
	newCache.Add(baseUrl, val)

	val, exists := newCache.Get(baseUrl)

	if !exists && val != nil {
		t.Error("Val should be nil if entry does not exists")
	}
	if exists && val == nil {
		t.Error("Val should not be nil if entry exists")
	}
}

func TestCreateCache(t *testing.T) {
	newCache := NewCache(time.Duration(20 * time.Second))

	if newCache.interval == 0 {
		t.Errorf("ERR: Interval -> %v cannot be 0", newCache.interval)
	}

	if newCache.entries == nil {
		t.Errorf("Entries cannot be nil")
	}
}
