package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}


type CacheItem struct {
	Value *ListItem
	Index *ListItem
}

type lruCache struct {
	capacity int
	queue    List
	idxQueue List
	items    map[Key]*CacheItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		idxQueue: NewList(),
		items:    make(map[Key]*CacheItem, capacity),
	}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	_, ok := lc.items[key]
	vl := lc.queue.PushFront(value)
	idx := lc.idxQueue.PushFront(key)
	lc.items[key] = &CacheItem{Value: vl, Index: idx}
	lc.checkCapacity()
	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	itm, ok := lc.items[key]
	if !ok {
		return nil, false
	}
	lc.queue.MoveToFront(itm.Value)
	lc.idxQueue.MoveToFront(itm.Index)
	lc.checkCapacity()
	return itm.Value.Value, true
}

func (lc *lruCache) Clear() {
	lc.queue.RemoveAll()
	lc.idxQueue.RemoveAll()
	lc.items = make(map[Key]*CacheItem, lc.capacity)
}

func (lc *lruCache) checkCapacity() {
	if lc.queue.Len() <= lc.capacity {
		return
	}
	lastIdxNode := lc.idxQueue.Back()
	lastIdxKey := lastIdxNode.Value
	lastItem := lc.items[lastIdxKey.(Key)]
	vl := lastItem.Value
	idx := lastItem.Index
	lc.queue.Remove(vl)
	lc.idxQueue.Remove(idx)
	delete(lc.items, lastIdxKey.(Key))
}
