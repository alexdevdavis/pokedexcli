package pokecache

import (
	"testing"
	"time"
)

type cacheTestCase struct {
	name        string
	key         string
	value       []byte
	expireAfter time.Duration
	sleepBefore time.Duration
	expectFound bool
	expectValue []byte
}

func TestCache_Behavior(t *testing.T) {
	testCases := []cacheTestCase{
		{
			name:        "Add and Get - Valid Entry",
			key:         "pikachu",
			value:       []byte("electric"),
			expireAfter: 5 * time.Second,
			sleepBefore: 0,
			expectFound: true,
			expectValue: []byte("electric"),
		},
		{
			name:        "Expiration - Entry Expired",
			key:         "bulbasaur",
			value:       []byte("grass"),
			expireAfter: 1 * time.Second,
			sleepBefore: 2 * time.Second,
			expectFound: false,
		},
		{
			name:        "ReapLoop - Entry Reaped",
			key:         "charmander",
			value:       []byte("fire"),
			expireAfter: 500 * time.Millisecond,
			sleepBefore: 1 * time.Second,
			expectFound: false,
		},
		{
			name:        "Non-Existent Key",
			key:         "missingno",
			value:       nil,
			expireAfter: 5 * time.Second,
			sleepBefore: 0,
			expectFound: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache := NewCache(tc.expireAfter)
			if tc.name == "ReapLoop - Entry Reaped" {
				go cache.reapLoop()
			}

			if tc.value != nil {
				err := cache.Add(tc.key, tc.value)
				if err != nil {
					t.Fatalf("unexpected error adding to cache: %v", err)
				}
			}

			time.Sleep(tc.sleepBefore)

			retrievedValue, found := cache.Get(tc.key)
			if found != tc.expectFound {
				t.Errorf("expected key %s found status to be %v, got %v", tc.key, tc.expectFound, found)
			}

			if tc.expectFound && string(retrievedValue) != string(tc.expectValue) {
				t.Errorf("expected value %s, got %s", tc.expectValue, retrievedValue)
			}
		})
	}
}
