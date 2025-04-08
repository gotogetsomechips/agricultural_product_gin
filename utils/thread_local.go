package utils

import (
	"sync"
)

// ThreadLocal 线程本地存储
type ThreadLocal struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

var userLocal = ThreadLocal{
	store: make(map[string]interface{}),
}

// Set 存储值
func (t *ThreadLocal) Set(key string, value interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.store[key] = value
}

// Get 获取值
func (t *ThreadLocal) Get(key string) interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.store[key]
}

// Remove 移除值
func (t *ThreadLocal) Remove(key string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.store, key)
}

// GetUserLocal 获取用户线程本地存储
func GetUserLocal() *ThreadLocal {
	return &userLocal
}
