package main

import (
	"container/list"
	"fmt"
)

type Value interface {
	Len() int
}
type LRUEntry struct {
	key   string
	value Value
}
type Cache struct {
	//1.内存最大值
	maxBytes int
	//2.当前内存所在值
	nowBytes int
	//3.双向链表
	ll *list.List

	cache map[string]*list.Element
}

func (c *Cache) get(key string) (value Value, ok bool) {
	if e, ok := c.cache[key]; ok {
		c.ll.MoveToFront(e)
		kv := e.Value.(*LRUEntry)
		return kv.value, true
	}
	return
}

func (c *Cache) add(key string, value Value) {
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		kv := element.Value.(*LRUEntry)
		c.nowBytes += int(value.Len()) - int(kv.value.Len())
		kv.value = value
	} else {
		element := c.ll.PushFront(&LRUEntry{key, value})
		c.cache[key] = element
		c.nowBytes += int(len(key)) + int(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nowBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) RemoveOldest() {
	element := c.ll.Back()
	if element != nil {
		c.ll.Remove(element)
		kv := element.Value.(*LRUEntry)
		delete(c.cache, kv.key)
		c.nowBytes -= int(len(kv.key)) + int(kv.value.Len())
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}

func New(maxBytes int) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func main() {
	fmt.Println("Hello, World!")
}
