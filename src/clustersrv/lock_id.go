package clustersrv

import (
	"sync"

	"github.com/pborman/uuid"
)

// LockID is a concurrency safe unique identifier for locks
type LockID struct {
	id  string
	lck *sync.RWMutex
}

// NewLockID creates a new lock ID and returns it
func NewLockID() *LockID {
	return &LockID{id: uuid.New(), lck: &sync.RWMutex{}}
}

// Equals determines whether str is equal to the internal lock ID. This func is concurrency safe
func (l *LockID) Equals(str string) bool {
	l.lck.RLock()
	defer l.lck.RUnlock()
	return l.id == str
}

// Generate generates a new lock ID, overwrites the old, and returns the new. This func is concurrency safe
func (l *LockID) Generate() string {
	l.lck.Lock()
	defer l.lck.Unlock()
	l.id = uuid.New()
	return l.id
}
