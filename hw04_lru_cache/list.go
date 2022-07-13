package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	// Place your code here.
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.Len() == 0 {
		l.front = item
		l.back = item
	} else {
		item.Next = l.front
		item.Next.Prev = item
		l.front = item
	}
	l.len++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.Len() == 0 {
		l.front = item
		l.back = item
	} else {
		item.Prev = l.back
		item.Prev.Next = item
		l.back = item
	}
	l.len++
	return item
}

func (l *list) Remove(i *ListItem) {
	l.len--
	if l.Len() == 0 {
		l.front = nil
		l.back = nil
		return
	}
	if i == l.front {
		l.front = i.Next
		l.front.Prev = nil
		return
	}
	if i == l.back {
		l.back = i.Prev
		l.back.Next = nil
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}
	if i == l.back {
		l.back = i.Prev
	}
	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	i.Next = l.front
	l.front.Prev = i
	l.front = i
	i.Prev = nil
}
