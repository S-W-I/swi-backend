package session

import (
	"encoding/base64"

	"github.com/google/uuid"
)


type SessionID string
func (session SessionID) Encoded() string {
	encoded := base64.StdEncoding.EncodeToString([]byte(session))
	return encoded
}
func (session SessionID) Decoded() string {
	return string(session)
}

func NewSessionID() SessionID {
	return SessionID(uuid.NewString())
}
type SessionManagerMeta struct {
	WorkdirPath, TemplatePath string
}
