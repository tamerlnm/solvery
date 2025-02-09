package model

import "container/list"

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
