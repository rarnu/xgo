package xgo

import (
	"errors"
	"reflect"
	"runtime"
	"sync"
	"time"
)

type Cache struct {
	defaultExpiration time.Duration
	items             map[string]Item
	mu                sync.RWMutex
	onEvicted         func(string, any)
	j                 *janitor
}

type Item struct {
	//data
	Data any
	//ttl time.UnixNano.Expiration of Item,if it is -1,it will be not Expired
	Ttl int64
}

func (item *Item) Expired() bool {
	if item.Ttl == -1 {
		//用户
		return false
	}
	return time.Now().UnixNano() > item.Ttl
}

func NewCache() *Cache {
	return NewWithExpiration(0)
}

func NewWithExpiration(expiration time.Duration) *Cache {
	return NewWithExpirationAndCleanupInterval(expiration, 0)
}

func NewWithCleanupInterval(cleanupInterval time.Duration) *Cache {
	return NewWithExpirationAndCleanupInterval(0, cleanupInterval)
}
func NewWithExpirationAndCleanupInterval(defaultExpiration, cleanupInterval time.Duration) *Cache {
	if defaultExpiration == 0 {
		defaultExpiration = -1
	}
	ch := make(chan bool)
	c := &Cache{
		defaultExpiration: defaultExpiration,
		items:             make(map[string]Item),
		j: &janitor{
			Interval: 1 * time.Second,
			stop:     ch,
		},
	}
	//启动清理协程
	go func() {
		c.runCleanup(cleanupInterval)
		runtime.SetFinalizer(c, stopJanitor)
	}()
	return c
}

type janitor struct {
	Interval time.Duration
	stop     chan bool
}

func (c *Cache) runCleanup(cleanupInterval time.Duration) {
	if cleanupInterval == 0 {
		cleanupInterval = 500 * time.Millisecond
	}
	ticker := time.NewTicker(cleanupInterval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-c.j.stop:
			ticker.Stop()
			return
		}
	}
}

func stopJanitor(c *Cache) {
	c.j.stop <- true
}

func (c *Cache) DeleteExpired() {
	l := len(c.items)
	if l < 1 {
		return
	}
	cloneMap := map[string]Item{}
	c.mu.Lock()
	for k, v := range c.items {
		cloneMap[k] = v
	}
	c.mu.Unlock()

	ch := make(chan int8, len(c.items))
	for key, item := range cloneMap {
		go func(k string, i Item) {
			if i.Expired() {
				c.mu.Lock()
				delete(c.items, k)
				c.mu.Unlock()
			}
			ch <- int8(1)
		}(key, item)
	}
}

func (c *Cache) getUnixNano() int64 {
	de := c.defaultExpiration
	var e int64
	if de != -1 {
		e = time.Now().Add(de).UnixNano()
	}
	return e
}

// Set Add an item to the cache,replacing any existing item.
// note key is primary key
func (c *Cache) Set(key string, value any) error {
	e := c.getUnixNano()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = Item{
		Data: value,
		Ttl:  e,
	}
	return nil
}

// Get an item from the cache.Returns the item or nil, and a bool indicating
// whether the key was found
func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, found := c.items[key]; !found {
		return nil, false
	} else {
		//check item has expired
		if item.Ttl > 0 && time.Now().UnixNano() > item.Ttl {
			return nil, false
		}
		return item.Data, true
	}
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, found := c.items[key]; found {
		delete(c.items, key)
		return
	}
}

func (c *Cache) Clean() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k := range c.items {
		delete(c.items, k)
	}
}

func (c *Cache) Cap() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	ci := c.items
	return len(ci)
}

func (c *Cache) SetHash(key, subKey string, value any) error {
	e := c.getUnixNano()
	c.mu.Lock()
	defer c.mu.Unlock()

	if vv, b := c.items[key]; b {
		vv.Data.(map[string]any)[subKey] = value
	} else {
		subKeyValue := make(map[string]any)
		subKeyValue[subKey] = value
		v := Item{
			Data: subKeyValue,
			Ttl:  e,
		}
		c.items[key] = v
	}
	return nil
}

func (c *Cache) RemoveHash(key, subKey string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, found := c.items[key]; found {
		data := item.Data
		if reflect.TypeOf(data).Kind() != reflect.Map {
			return errors.New("key的值不是hash")
		}
		subValue := data.(map[string]any)
		delete(subValue, subKey)
		if len(subValue) == 0 {
			delete(c.items, key)
		}
		return nil
	}
	return nil
}

// GetHash get a hash value from the cache.Returns the hashes or nil, and a bool indicating
// whether the key was found
func (c *Cache) GetHash(key, subKey string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, found := c.items[key]; !found {
		return nil, found
	} else {
		data := item.Data
		if reflect.TypeOf(data).Kind() != reflect.Map {
			return nil, false
		}
		if value, found := item.Data.(map[string]any)[subKey]; !found {
			return nil, found
		} else {
			return value, true
		}
	}
}

func (c *Cache) AddItem(key string, value ...any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, found := c.items[key]; found {
		data := item.Data
		if reflect.TypeOf(data).Kind() != reflect.Slice {
			return errors.New("key 对应的数据类型不是 slice")
		}
		item.Data = append(data.([]any), value...)
		c.items[key] = item
	} else {
		e := c.getUnixNano()
		data := Item{
			Data: value,
			Ttl:  e,
		}
		c.items[key] = data
	}
	return nil
}

// SetItem set or replace a value of items by index
func (c *Cache) SetItem(key string, idx int, value any) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, found := c.items[key]; found {
		data := item.Data
		if reflect.TypeOf(data).Kind() != reflect.Slice {
			return errors.New("key 对应的数据类型不是 slice")
		}
		items := data.([]any)
		if len(items) <= idx {
			return errors.New("数组下标越界")
		}

		items[idx] = value
		item.Data = items
		return nil
	}
	return errors.New("key不存在")
}

// GetItem return an array of points or nil
func (c *Cache) GetItem(key string) []any {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, found := c.items[key]; !found {
		return nil
	} else {
		data := item.Data
		if reflect.TypeOf(data).Kind() != reflect.Slice {
			return nil
		}
		return data.([]any)
	}
}

// GetItemByIndex return a value of Type is T or nil
func (c *Cache) GetItemByIndex(key string, idx int) any {
	if idx < 0 {
		return nil
	}
	if items := c.GetItem(key); items != nil {
		if len(items) <= idx {
			return nil
		}
		return items[idx]
	}
	return nil
}

// RemoveItem an item from the cache. Does nothing if the key is not in the cache.
func (c *Cache) RemoveItem(key string, idx int) error {
	if idx < 0 {
		c.Remove(key)
		return nil
	}
	item := c.GetItem(key)
	if item == nil {
		return errors.New("key不存在")
	}
	if len(item) <= idx {
		return nil
	}
	newItem := item[:idx]
	newItem = append(newItem, item[idx+1:]...)
	_ = c.Set(key, newItem)
	return nil
}
