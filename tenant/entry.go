package tenant

import (
	"time"
	"unsafe"
)

type Entry[T any] struct {
	id      int        // unique identifier
	ttl     *time.Time // expiration time
	created time.Time  // creation time
	mod     time.Time  // last modification time
	offset  int        // offset in file
	length  int        // length in file
	value   []byte     // slice into file
}

func (e *Entry[T]) Id() int {
	return e.id
}

func (e *Entry[T]) Modified() time.Time {
	return e.mod
}

func (e *Entry[T]) Created() time.Time {
	return e.created
}

func (e *Entry[T]) Size() int {
	return e.length
}

func (e *Entry[T]) TTL() time.Duration {
	if e.ttl == nil {
		return 0
	}
	return e.ttl.Sub(time.Now())
}

func (e *Entry[T]) Expired() bool {
	if e.ttl == nil {
		return false
	}
	return time.Now().After(*e.ttl)
}

func (e *Entry[T]) Value() T {
	tmp := new(T)
	*tmp = *(*T)(unsafe.Pointer(&e.value[0]))
	return *tmp
}

func (e *Entry[T]) SetTTL(ttl time.Duration) {
	if ttl > 0 {
		n := time.Now().Add(ttl)
		e.ttl = &n
	} else {
		e.ttl = nil
	}
}
