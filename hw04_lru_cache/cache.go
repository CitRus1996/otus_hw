package hw04lrucache

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
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l lruCache) Set(key Key, value interface{}) bool {
	if _, ok := l.items[key]; !ok {
		l.checkLen()
		item := l.queue.PushFront(value)
		l.items[key] = item
		return false
	} else {
		l.items[key].Value = value
		l.queue.MoveToFront(l.items[key])
		return true
	}
}

func (l lruCache) Get(key Key) (interface{}, bool) {
	if _, ok := l.items[key]; !ok {
		return nil, false
	} else {
		l.queue.MoveToFront(l.items[key])
		return l.items[key].Value, true
	}
}

func (l lruCache) Clear() {
	l.items = make(map[Key]*ListItem)
	l.queue = NewList()
}

func (l lruCache) checkLen() {
	if l.queue.Len() > l.capacity {
		last := l.queue.Back()
		for k := range l.items {
			if l.items[k] == last {
				delete(l.items, k)
			}
		}
		l.queue.Remove(last)
	}
}
