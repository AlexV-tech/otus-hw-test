package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	RemoveAll()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (ll *list) Len() int {
	return ll.len
}

func (ll *list) Front() *ListItem {
	return ll.front
}

func (ll *list) Back() *ListItem {
	return ll.back
}

func (ll *list) PushFront(v interface{}) *ListItem {
	nItem := &ListItem{Value: v, Next: nil, Prev: nil}
	if ll.front == nil {
		ll.front = nItem
		ll.back = nItem
		ll.len = 1
	} else {
		nItem.Next = ll.front
		ll.front.Prev = nItem
		ll.front = nItem
		ll.len++
	}
	return nItem
}

func (ll *list) PushBack(v interface{}) *ListItem {
	nItem := &ListItem{Value: v, Next: nil, Prev: nil}
	if ll.back == nil {
		ll.back = nItem
		ll.front = nItem
		ll.len = 1
	} else {
		nItem.Prev = ll.back
		ll.back.Next = nItem
		ll.back = nItem
		ll.len++
	}
	return nItem
}

func (ll *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	if (i == ll.front) && (ll.front == ll.back) {
		ll.front = nil
		ll.back = nil
		ll.len--
		return
	}
	if i == ll.front {
		ll.front = i.Next
	}
	if i == ll.back {
		ll.back = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	ll.len--
}

func (ll *list) MoveToFront(i *ListItem) {
	if i == nil || i == ll.front {
		return
	}
	ll.Remove(i)
	ll.PushFront(i.Value)
}

func (ll *list) RemoveAll() {
	for i := ll.Front(); i != nil; i = i.Next {
		ll.Remove(i)
	}
}
