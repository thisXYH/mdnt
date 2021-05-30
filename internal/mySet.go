package internal

import "sync"

type any interface{}

// Set 接口
type Set interface {
	Add(any) bool
	Del(any)
	IsExist(any) bool
	Len() int
	Iterator() []any
}

// MySet 自定义的一个线程安全的set
type MySet struct {
	inner map[any]bool
	mu    *sync.RWMutex
}

func NewMySet() *MySet {
	return &MySet{inner: make(map[any]bool), mu: &sync.RWMutex{}}
}

func (set *MySet) Add(item any) bool {
	set.mu.Lock()
	defer set.mu.Unlock()

	if set.IsExist(item) {
		return false
	}

	set.inner[item] = true
	return true
}

func (set *MySet) Del(item any) {
	set.mu.Lock()
	defer set.mu.Unlock()
	delete(set.inner, item)
}

func (set *MySet) IsExist(item any) bool {
	_, ok := set.inner[item]
	return ok
}

func (set *MySet) Len() int {
	return len(set.inner)
}

func (set *MySet) Iterator() []any {
	items := make([]any, 0, set.Len())

	for k := range set.inner {
		items = append(items, k)
	}
	return items
}
