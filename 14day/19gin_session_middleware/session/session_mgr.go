package session

// SessionMgr is an interface that has three methods: Init, CreateSession, and Get.
// @property Init - Initialize the session manager.
// @property CreateSession - Create a new session and return the session object.
// @property Get - Get the session object by sessionId
type SessionMgr interface {
	Init(addr string, options ...string)
	CreateSession() (session Session, err error)
	Get(sessionId string) (session Session, err error)
}
