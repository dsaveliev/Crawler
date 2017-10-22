package pool

import (
	"crawler/models"
	"sync"
	"sync/atomic"
)

type StorageMap map[uint64]*models.Task
type StorageType struct {
	mx    sync.RWMutex
	m     StorageMap
	index uint64
}

func NewStorage() *StorageType {
	return &StorageType{
		m:     StorageMap{},
		index: 0,
	}
}

func (s *StorageType) NextIndex() uint64 {
	return atomic.AddUint64(&(s.index), 1)
}

func (s *StorageType) Get(id uint64) (task *models.Task, ok bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	task, ok = s.m[id]
	return
}

func (s *StorageType) Set(id uint64, task *models.Task) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.m[id] = task
}

func (s *StorageType) Delete(id uint64) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.m, id)
}

func (s *StorageType) Len() int {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return len(s.m)
}
