package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    *sync.RWMutex
}

//  type cacheItem struct {
//		key   Key
//		value interface{}
//	}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    new(sync.RWMutex),
	}
}

func (l lruCache) Set(key Key, value interface{}) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if _, ok := l.items[key]; !ok {
		l.checkLen()
		item := l.queue.PushFront(value)
		l.items[key] = item
		return false
	}
	l.items[key].Value = value
	l.queue.MoveToFront(l.items[key])
	return true
}

func (l lruCache) Get(key Key) (interface{}, bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if _, ok := l.items[key]; !ok {
		return nil, false
	}
	l.queue.MoveToFront(l.items[key])
	item := l.items[key].Value
	return item, true
}

func (l *lruCache) Clear() {
	l.items = make(map[Key]*ListItem)
	l.queue = NewList()
}

func (l lruCache) checkLen() {
	if l.queue.Len() >= l.capacity {
		last := l.queue.Back()
		for k := range l.items {
			if l.items[k] == last {
				delete(l.items, k)
			}
		}
		l.queue.Remove(last)
	}
}
