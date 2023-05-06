package service

import "sync"

type Storage struct {
	mux                  sync.Mutex
	totalRequestCount    int
	successRequestsCount int
	errorRequestsCount   int
}

func (s *Storage) IncTotalRequestCount(delta int) {
	if delta == 0 {
		delta = 1
	}
	s.mux.Lock()
	defer s.mux.Unlock()

	s.totalRequestCount += delta
}

func (s *Storage) GetTotalRequestCount() int {
	return s.totalRequestCount
}

func (s *Storage) IncSuccessRequestCount() {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.successRequestsCount++
}

func (s *Storage) GetSuccessRequestCount() int {
	return s.successRequestsCount
}

func (s *Storage) IncErrorRequestCount() {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.errorRequestsCount++
}

func (s *Storage) GetErrorRequestCount() int {
	return s.errorRequestsCount
}
