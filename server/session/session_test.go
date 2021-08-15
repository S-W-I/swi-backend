package session

import (
	"fmt"
	"reflect"
	"swi/utility"
	"testing"
)


func TestSessionManagerBehaviour_CreateNDrop(t *testing.T) {
	// workDir := "/usr/local/var/www/swi-backend/sessions"
	// templateDir := "/usr/local/var/www/swi-backend/solana-template"
	workDir := "../../sessions"
	templateDir := "../../solana-template"

	sessionManager, err := NewSessionManagerAt(workDir, templateDir)
	utility.ValidateError(t, err)

	newSession, err := sessionManager.NewSession()

	utility.Assert(t, err == nil)
	utility.Assert(t, sessionManager.SessionExists(newSession.SessionId))

	// drop session
	err = sessionManager.DropSession(newSession.SessionId)

	utility.Assert(t, err == nil)
	utility.Assert(t, !sessionManager.SessionExists(newSession.SessionId))
}

func TestSessionManagerBehaviour_CreateRead_BuildTree_Drop(t *testing.T) {
	workDir := "../../sessions"
	templateDir := "../../solana-template"

	sessionManager, err := NewSessionManagerAt(workDir, templateDir)
	utility.ValidateError(t, err)

	newSession, err := sessionManager.NewSession()

	utility.ValidateError(t, err)
	utility.Assert(t, sessionManager.SessionExists(newSession.SessionId))
	
	sessionTree, err := sessionManager.BuildSessionTreeFor(newSession.SessionId)
	utility.ValidateError(t, err)
	fmt.Printf("Session Tree: %+v \n", sessionTree)

	err = sessionManager.DropSession(newSession.SessionId)

	utility.Assert(t, err == nil)
	utility.Assert(t, !sessionManager.SessionExists(newSession.SessionId))
}

func TestSessionManagerBehaviour_CreateRead_Write_BuildTree_Drop(t *testing.T) {
	workDir := "../../sessions"
	templateDir := "../../solana-template"

	sessionManager, err := NewSessionManagerAt(workDir, templateDir)
	utility.ValidateError(t, err)

	newSession, err := sessionManager.NewSession()

	utility.ValidateError(t, err)
	utility.Assert(t, sessionManager.SessionExists(newSession.SessionId))
	
	// retrieve tree
	sessionTreeBefore, err := sessionManager.BuildSessionTreeFor(newSession.SessionId)
	utility.ValidateError(t, err)
	

	inputBody := sessionTreeBefore.FilePaths
	newInputBody := make(map[string]string)

	for key, value := range inputBody {
		rawValue, err := DecodeContent(value)
		utility.ValidateError(t, err)

		rawValue = append(rawValue, []byte("\n random stuff")...)

		newInputBody[key] = EncodeContent(rawValue)
	}

	// substitute changes
	sessionTreeRightAfterEdit, err := sessionManager.UpdateSessionData(newSession.SessionId, newInputBody)
	utility.ValidateError(t, err)

	utility.Assert(t, !reflect.DeepEqual(sessionTreeRightAfterEdit.FilePaths, sessionTreeBefore.FilePaths))


	sessionTreeAfter, err := sessionManager.BuildSessionTreeFor(newSession.SessionId)
	utility.ValidateError(t, err)

	utility.Assert(t, reflect.DeepEqual(sessionTreeRightAfterEdit.FilePaths, sessionTreeAfter.FilePaths))

	err = sessionManager.DropSession(newSession.SessionId)

	utility.ValidateError(t, err)
	utility.Assert(t, !sessionManager.SessionExists(newSession.SessionId))
}