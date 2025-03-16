package tests

import (
	"pokedexcli/internal/pokecache"
	"testing"
	"time"
)

func CachingTests(t *testing.T) {
	cache := pokecache.NewCache(nil, time.Duration(time.Duration.Seconds(10)))

	cache.Add("test", []byte("test"))
	cache.Add("test2", []byte("test2"))

	if len(cache.Entries) != 2 {
		t.Error("Incorrect Implementation of Add Function")
		t.Fail()
	}

	val, ok := cache.Get("test2")
	if !ok || string(val) != "test2" {
		t.Error("Incorrect Implementation of Get Function")
		t.Fail()
	}
}
