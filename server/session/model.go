package session

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// CompilationResult

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


type SessionTree struct {
	// FilePaths map[string][]byte
	FilePaths map[string]string `json:"file_paths"`
}

type SessionManagerMeta struct {
	WorkdirPath, TemplatePath string
}


type CompilationResult struct {
	Version      string `json:"version"`
	CompileError bool `json:"compile_error"`
	Message      string `json:"message"`
}