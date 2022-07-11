package session

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	addr       string
	password   string
	pool       *redis.Pool
	rwLock     sync.RWMutex
	sessionMap map[string]Session
}

func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{sessionMap: make(map[string]Session, 32)}
	return sr
}

func (r *RedisSessionMgr) Init(addr string, options ...string) {
	if len(options) > 0 {
		r.password = options[0]
	}
	// 初始化连接池
	r.pool = initRedisPool(addr, r.password)
	r.addr = addr

}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	// go get github.com/satori/go.uuid
	// 用UUID作为sessionId
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	// 转字符串
	sessionId := id.String()
	// 创建 session
	memorySession := NewRedisSession(sessionId, r.pool)
	r.sessionMap[sessionId] = memorySession
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exist!!!")
		return
	}
	return
}

// initRedisPool 创建连接池
func initRedisPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: time.Second * 10,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			// 若有密码，判断
			if _, err := conn.Do("AUTH", password); err != nil {
				err := conn.Close()
				if err != nil {
					return nil, err
				}
				return nil, err
			}
			return conn, nil
		},
		// 连接测试，开发时写
		// 上线时注释掉
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
