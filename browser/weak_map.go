package browser

import (
	"fmt"
	"runtime"
)

type NodeMap[T any] struct {
	entries map[int]T
}

func NewNodeMap[T any]() *NodeMap[T] {
	return &NodeMap[T]{
		entries: make(map[int]T),
	}
}

func (m *NodeMap[T]) Length() int { return len(m.entries) }

func (m *NodeMap[T]) Set(k Node, val T) {
	fmt.Println("Adding object:", k.ObjectId())
	objectId := k.ObjectId()
	id := objectId.Id
	runtime.SetFinalizer(objectId.Trigger, func(x *Dummy) {
		fmt.Println("Finalizer")
		delete(m.entries, id)
	})
	m.entries[id] = val
}

/*
type Cache[K any, V any] struct {
    f func(*K) V
    m atomic.Map[uintptr, func() V]
}

func NewCache[K comparable, V any](f func(*K)V) *Cache[K, V] {
    return &Cache[K, V]{f: f}
}

func (c *Cache[K, V]) Get(k *K) V {
    kw := uintptr(unsafe.Pointer((k))
    vf, ok := c.m.Load(kw)
    if ok {
        return vf()
    }
    vf = sync.OnceValue(func() V { return c.f(k) })
    vf, loaded := c.m.LoadOrStore(kw)
    if !loaded {
        // Stored kw→vf to c.m; add the cleanup.
        runtime.AddCleanup(k, c.cleanup, kw)
    }
    return vf()
}

func (c *Cache[K, V]) cleanup(kw uintptr) {
    c.m.Delete(kw)
}

*/
