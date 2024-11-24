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

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	_, ok := lc.items[key]
	lc.items[key] = lc.queue.PushFront(value)
	lc.checkCapacity()
	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	itm, ok := lc.items[key]
	if !ok {
		return nil, false
	}
	lc.queue.MoveToFront(itm)
	lc.checkCapacity()
	return itm.Value, true
}

func (lc *lruCache) Clear() {
	lc.queue.RemoveAll()
	lc.items = make(map[Key]*ListItem, lc.capacity)
}

func (lc *lruCache) checkCapacity() {
	if lc.queue.Len() <= lc.capacity {
		return
	}
	bck := lc.queue.Back()
	for k, v := range lc.items {
		if v == bck {
			delete(lc.items, k)
			break
		}
	}
	lc.queue.Remove(bck)
}
