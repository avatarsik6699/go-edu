package condv1

import "sync"

type Semaphore struct {
	count int
	limit int
	cond  *sync.Cond
}

func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{
		count: 0,
		limit: limit,
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Semaphore) Acquire() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	for s.count == s.limit {
		s.cond.Wait()
	}

	s.count++
}

func (s *Semaphore) Release() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	s.count--
	s.cond.Signal()
}

func (s *Semaphore) Available() int {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	return s.count
}
