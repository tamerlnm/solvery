package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
}

type Item struct {
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRUCache) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.key)
	}
}

func (с *LRUCache) Get(key int) (int, bool) {
	if element, exist := с.items[key]; exist {
		с.queue.MoveToFront(element)
		return element.Value.(*Item).value, true
	}
	return 0, false
}

func (c *LRUCache) Set(key int, value int) {
	if element, exists := c.items[key]; exists {
		element.Value.(*Item).value = value
		c.queue.MoveToFront(element)
		return
	}

	if c.queue.Len() >= c.capacity {
		c.purge()
	}

	item := &Item{
		key:   key,
		value: value,
	}

	el := c.queue.PushFront(item)
	c.items[item.key] = el
}

func main() {
	cache := NewLRUCache(2)

	cache.Set(1, 10)
	cache.Set(2, 20)

	if val, ok := cache.Get(1); ok {
		fmt.Println("Key 1:", val)
	}

	cache.Set(3, 30)

	if _, ok := cache.Get(2); !ok {
		fmt.Println("Key 2 not found")
	}

	cache.Set(4, 40)

	if _, ok := cache.Get(1); !ok {
		fmt.Println("Key 1 not found")
	}

	if val, ok := cache.Get(3); ok {
		fmt.Println("Key 3:", val)
	}

	if val, ok := cache.Get(4); ok {
		fmt.Println("Key 4:", val)
	}
}
