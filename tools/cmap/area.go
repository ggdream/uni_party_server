package cmap

import "sync"

// cmapSubarea 单块分区
type cmapSubarea struct {
	// 读写锁
	lock sync.RWMutex
	// 真正的数据
	data map[string]interface{}
}

// Set 添加新值
func (c *cmapSubarea) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[key] = value
}

// Get 添加新值
func (c *cmapSubarea) Get(key string) (value interface{}, ok bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	value, ok = c.data[key]
	return
}

// Del 删除值
func (c *cmapSubarea) Del(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.data, key)
}
