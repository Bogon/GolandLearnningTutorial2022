package session

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"sync"
)

type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	// 设置session，可以先放置到内存中
	// 批量导入到redis，提升性能
	sessionMap map[string]interface{}
	// 读写锁
	rwLock *sync.RWMutex

	// 记录内存中map是否被操作
	flag int
}

// 用常量定义状态
const (
	// 数据在内存中没有变化
	SessionFlagNone = iota
	// 有变化
	SessionFlagModify
)

// NewRedisSession Create a new Redis session with the given ID and Redis pool.
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{}, 1024),
		flag:       SessionFlagNone,
	}
	return s
}

// Set The above code is locking the sessionMap, and then setting the value of the key to the value of the value.
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.sessionMap[key] = value
	// 标记记录
	r.flag = SessionFlagModify
	return
}

// Save
// 1. Lock the session
// 2. Check if the session has been modified
// 3. If not, return
// 4. Marshal the session data
// 5. Get a redis connection
// 6. Save the session data to redis
// 7. Unlock the session
// 8. Set the session flag to none
func (r *RedisSession) Save() (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	// 如果数据为发生变化不需要存储
	if r.flag == SessionFlagNone {
		return
	}
	// 将内存中的数据进行序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	// 获取redis 链接
	conn := r.pool.Get()
	// 保存数据到redis
	_, err = conn.Do("SET", r.sessionId, string(data))
	if err != nil {
		return
	}
	r.flag = SessionFlagNone
	return
}

func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key not exist！")
	}
	return
}

// 从redis中再次加载
func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	s, err := redis.String(reply, err)
	if err != nil {
		return
	}
	// json 反序列化
	err = json.Unmarshal([]byte(s), &r.sessionMap)
	if err != nil {
		return
	}
	return
}

func (r *RedisSession) Del(sessionId string) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, sessionId)
	return
}
