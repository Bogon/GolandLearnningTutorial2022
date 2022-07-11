package session

import "fmt"

var (
	sessionMgr SessionMgr
)

// Init 中间件让用户选择使用哪个版本
func Init(provider, addr string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()

	case "redis":
		sessionMgr = NewRedisSessionMgr()

	default:
		fmt.Errorf("不支持")
		return
	}
	return
}
