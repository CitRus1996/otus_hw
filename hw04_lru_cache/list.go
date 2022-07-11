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
	first *ListItem
	last  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.Len() == 0 {
		l.first = item
		l.last = item
	} else {
		item.Next = l.first
		l.first.Prev = item
		l.first = item
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
		l.first = item
		l.last = item
	} else {
		item.Prev = l.last
		l.last.Next = item
		l.last = item
	}
	l.len++
	return item
}

func (l *list) Remove(i *ListItem) {
	if i == l.first {
		l.first = i.Next
		l.first.Prev = nil
		return
	}
	if i == l.last {
		l.last = i.Prev
		l.last.Next = nil
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	i.Prev.Next = i.Next
	i.Next = l.first
	i.Prev = nil
	l.first.Prev = i
	l.first = i
}
