package session

import (
	"errors"
	"sync"
)

// MemorySession `MemorySession` is a struct with a `sessionId` field of type `string`, a `data` field of type `map[string]interface{}`,
// and a `rwLock` field of type `sync.RWMutex`.
// @property {string} sessionId - session id
// @property data - The data that is stored in the session.
// @property rwLock - Read-write lock, used to prevent concurrent access to data
type MemorySession struct {
	sessionId string
	// 存储 kv
	data map[string]interface{}

	rwLock sync.RWMutex
}

// NewMemorySession 构造函数
// It creates a new session and returns a pointer to it
func NewMemorySession(sid string) *MemorySession {
	s := &MemorySession{
		sessionId: sid,
		data:      make(map[string]interface{}, 100),
	}
	return s
}

//Set  A method of the struct `MemorySession`.
func (m *MemorySession) Set(key string, value interface{}) (err error) {
	// 加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	// 设置值
	m.data[key] = value
	return
}

// Get A method of the struct `MemorySession`.
func (m *MemorySession) Get(key string) (value interface{}, err error) {
	// 加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	// 设置值
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not found")
		return
	}
	return
}

// Del Deleting the key from the map.
func (m *MemorySession) Del(key string) (err error) {
	// 加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.data, key)
	return
}

// Save A method of the struct `MemorySession`
func (m *MemorySession) Save() (err error) {
	return
}
