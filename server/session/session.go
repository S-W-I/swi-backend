package session

import (
	sessionproto "swi/protobuf/session"
	"swi/server/db"
	"time"

	"github.com/google/uuid"
)

const (
	DefaultSessionTTL uint64 = 1 * 60 * 60 * 24 * 7 // unix timestamp
)

type SessionManager struct {
	// CurrentSessions map[string]sessionproto.Session
	StateDelegate   *db.StateManager
}

func NewSessionManager() *SessionManager {
	return &SessionManager{}
}


func (manager *SessionManager) NewSession() *sessionproto.Session {
	return &sessionproto.Session{
		Ttl: DefaultSessionTTL,
		SessionId: uuid.NewString(),
		CreatedAt: uint64(time.Now().Unix()),
	}
}


// type Session struct {

// }

type SessionRouter struct {
	
}