package service

import (
	"testing"
	"time"
)

func TestStorageIncTotalRequestCount(t *testing.T) {
	delta := 10
	s := Storage{}
	s.IncTotalRequestCount(delta)

	result := s.GetTotalRequestCount()
	if result != delta {
		t.Fatalf(`StorageIncTotalRequestCount count mismatch %d %d`, result, delta)
	}
}

func TestStorageIncTotalRequestAsyncCount(t *testing.T) {
	delta := 10
	s := Storage{}
	incCount := 1000
	for i := 0; i < incCount; i++ {
		go s.IncTotalRequestCount(delta)
	}

	time.Sleep(100 * time.Millisecond)
	result := s.GetTotalRequestCount()
	expected := delta * incCount
	if result != expected {
		t.Fatalf(`StorageIncTotalRequestCount count mismatch %d %d`, result, expected)
	}
}
