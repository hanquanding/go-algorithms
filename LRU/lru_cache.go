package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	Size  int                      // 缓存的最大元素个数
	list  *list.List               // 存储数据的双向链表
	cache map[string]*list.Element // 链表的map索引
}
type Entry struct {
	Key   string
	Value interface{}
}

// 初始化 cache := NewLRUCache(100)就是可以初始化100个缓存的元素列表
func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		Size:  size,
		list:  list.New(),
		cache: make(map[string]*list.Element),
	}
}

// 获取缓存长度
func (c *LRUCache) Len() int {
	return c.list.Len()
}

// 添加数据到缓存中
func (c *LRUCache) Add(key string, value interface{}) {
	// 初始化
	if c.cache == nil {
		c.cache = make(map[string]*list.Element)
		c.list = list.New()
	}
	// 获取到缓存后把它移到双向链表的最前，表示最热数据
	// 更新缓存项并移动都链表的头部
	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		e.Value.(*Entry).Value = value
		return
	}
	entry := &Entry{Key: key, Value: value}
	e := c.list.PushFront(entry)
	c.cache[key] = e
	// 缓存有限，把不是热数据移走
	if c.Size != 0 && c.list.Len() > c.Size {
		e := c.list.Back()
		if e != nil {
			// 移除list数据，同时删除map数据
			c.list.Remove(e)
			delete(c.cache, e.Value.(*Entry).Key)
		}
	}
}

// 获取缓存
func (c *LRUCache) Get(key string) interface{} {
	if c.cache == nil {
		return nil
	}
	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		return e.Value.(*Entry).Value
	}
	return nil
}

// 从缓存中移除指定数据
func (c *LRUCache) Remove(key string) {
	if c.cache == nil {
		return
	}
	if e, ok := c.cache[key]; ok {
		c.list.Remove(e)
		delete(c.cache, e.Value.(*Entry).Key)
	}
}
func main() {
	cache := NewLRUCache(2)
	cache.Add("aa", 1)
	cache.Add("bb", 2)
	cache.Add("cc", 3)
	fmt.Println(cache.Get("aa"))
	fmt.Println(cache.Get("bb"))
	fmt.Println(cache.Get("cc"))
}
