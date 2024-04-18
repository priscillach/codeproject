package utils

import (
	"sync"
)

var GlobalLock = sync.Mutex{}

type SyncLockByKey struct {
	locks map[string]*sync.Mutex
	mutex sync.Mutex
}

func NewSyncLockByeKey() *SyncLockByKey {
	return &SyncLockByKey{
		locks: make(map[string]*sync.Mutex),
	}
}

func (s *SyncLockByKey) AcquireSyncLock(key string) {
	s.mutex.Lock()

	if _, ok := s.locks[key]; !ok {
		s.locks[key] = &sync.Mutex{}
	}

	s.mutex.Unlock()

	lock := s.locks[key]
	lock.Lock()
}

func (s *SyncLockByKey) ReleaseSyncLock(key string) {
	lock := s.locks[key]
	if lock == nil {
		return
	}
	lock.Unlock()
}
