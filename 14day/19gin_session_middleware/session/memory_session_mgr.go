package session

import (
	"errors"
	"github.com/gofrs/uuid"
	"sync"
)

// MemorySessionMgr 设计思路
// 定义 MemorySessionMgr 对象(字段：存放所有 session 的map， 读写锁)
// 构造函数
// Init()
// CreateSession()
// GetSession()

// MemorySessionMgr > A MemorySessionMgr is a struct that contains a map of string to Session and a read-write lock.
// @property sessionMap - a map that stores the session information. The key is the session ID, and the value is the
// session object.
// @property rwLock - This is a read-write lock. It's used to synchronize access to the sessionMap.
type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

func (m *MemorySessionMgr) Init(addr string, options ...string) {
	return
}

func (m *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	session, ok := m.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exist!!!")
		return
	}
	return
}

// NewMemorySessionMgr > The function `NewMemorySessionMgr` returns a pointer to a `MemorySessionMgr` struct
func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{sessionMap: make(map[string]Session, 64)}
	return sr
}

// CreateSession It's a function that takes in no parameters and returns a session and an error.
func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	// go get github.com/satori/go.uuid
	// 用UUID作为sessionId
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	// 转字符串
	sessionId := id.String()
	// 创建 session
	memorySession := NewMemorySession(sessionId)
	m.sessionMap[sessionId] = memorySession
	return

}

// GetSession It's a function that takes in a string and returns a session and an error.
func (m *MemorySessionMgr) GetSession(sessionId string) (session Session, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	session, ok := m.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exist!!!")
		return
	}
	return
}
