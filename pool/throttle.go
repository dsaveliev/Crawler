package pool

import (
	"sync"
	"time"
)

var ThrottleRate = time.Second

type ChannelThrottle <-chan time.Time
type ThrottleMap map[string]ChannelThrottle
type ThrottleType struct {
	mx sync.RWMutex
	m  ThrottleMap
}

func NewThrottle() *ThrottleType {
	return &ThrottleType{
		m: ThrottleMap{},
	}
}

func (t *ThrottleType) Get(hostname string) ChannelThrottle {
	t.mx.Lock()
	defer t.mx.Unlock()
	if _, ok := t.m[hostname]; !ok {
		t.m[hostname] = time.Tick(ThrottleRate)
	}
	return t.m[hostname]
}

func (t *ThrottleType) Len() int {
	t.mx.RLock()
	defer t.mx.RUnlock()
	return len(t.m)
}
