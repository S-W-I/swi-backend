package session

import (
	"encoding/base64"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	sessionproto "swi/protobuf/session"

	// sessionproto "swi/protobuf/session"
	"swi/server/db"
	"time"
)

const (
	DefaultSessionTTL uint64 = 1 * 60 * 60 * 24 * 7 // unix timestamp
)

type SessionManager struct {
	// CurrentSessions map[string]sessionproto.Session
	StateDelegate   *db.StateManager

	workdir string
	templatePath string

	sessionWorkdirs []string
	sessions map[string]*sessionproto.Session
}

func (manager *SessionManager) SetMeta(meta SessionManagerMeta) *SessionManager {
	if meta.TemplatePath != "" {
		manager.templatePath = meta.TemplatePath
	}
	if meta.WorkdirPath != "" {
		manager.workdir = meta.WorkdirPath
	}

	return manager
}

func (manager *SessionManager) RetrieveSession(sessionId string) (*sessionproto.Session, error) {
	if !manager.SessionExists(sessionId) {
		return nil, fmt.Errorf("no such session")
	}

	return manager.sessions[sessionId], nil
}

func (manager *SessionManager) DownloadSolanaCompiledProject(sessionId string) ([]byte, error) {
	if !manager.SessionExists(sessionId) {
		return nil, fmt.Errorf("no such session")
	}

	projectDirectory := manager.workdir + "/" + sessionId
	compiledFilePath := projectDirectory + "/target/deploy/helloworld.so"

	compiledFileData, err := os.ReadFile(compiledFilePath)
	if err != nil {
		return make([]byte, 0), err
	}

	return compiledFileData, nil
}

func (manager *SessionManager) CompileSolanaProject(sessionId string) (*sessionproto.CompilationInfo, error) {
	if !manager.SessionExists(sessionId) {
		return nil, fmt.Errorf("no such session")
	}

	projectDirectory := manager.workdir + "/" + sessionId
	// cargoTomlCfgPath := projectDirectory + "/Cargo.toml"
	compileLogPath := projectDirectory + "/output.log"

	// cmd := exec.Run("cargo", "build-bpf", "--manifest-path", cargoTomlCfgPath, "&> ./output.log")
	// cmd := exec.Command("cargo", "build-bpf", "--manifest-path", cargoTomlCfgPath, "&>", "./result.log")
	cmd := exec.Command("/bin/bash", "wrapped-compile.sh", projectDirectory)

	result := new(sessionproto.CompilationInfo)
	result.Version = "1.6.9"

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	result.Message = string(out)

	logBytes, err := ioutil.ReadFile(compileLogPath)
	if err != nil {
		return nil, err
	}
	logs := string(logBytes)

	result.Message = logs

	if !strings.Contains(logs, "To deploy this program:") {
		result.CompileError = true
	}

	return result, nil
}

func (manager *SessionManager) UpdateSessionData(sessionId string, sessionTree map[string]string) (*sessionproto.SessionTree, error) {
	if !manager.SessionExists(sessionId) {
		return nil, fmt.Errorf("no such session")
	}

	readRoot := manager.workdir
	for filepath, filedata := range sessionTree {
		innerFilePath := readRoot + "/" + filepath
		content, err := ioutil.ReadFile(innerFilePath)
		if err != nil {
			return nil, err
		}

		if EncodeContent(content) == filedata {
			continue
		}

		// substitute changes
		decodedData, err := DecodeContent(filedata)
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(innerFilePath, decodedData, fs.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return manager.BuildSessionTreeFor(sessionId)
}

func (manager *SessionManager) BuildSessionLegacyTreeFor(sessionId string) ([]*sessionproto.SessionLegacyNode, error) {
	tree, err := manager.BuildSessionTreeFor(sessionId)

	if err != nil {
		return nil, err
	}

	// building legacy
	// legacyNode := new(sessionproto.SessionLegacyNode)
	// legacyNode.Name = sessionId
	// legacyNode.IsFile = false
	var result []*sessionproto.SessionLegacyNode

	for _, path := range tree.FilesOrder {
		legacyNode := new(sessionproto.SessionLegacyNode)
		buildLegacyNode(legacyNode, path)
		result = append(result, legacyNode.Children[0].Children...)
	}

	// return legacyNode.Children, nil
	
	return result, nil
}

func buildLegacyNode(parent *sessionproto.SessionLegacyNode, path string) {
	if strings.Contains(path, "/") {
		splittedPath := strings.Split(path, "/")
		name := splittedPath[0]
		nextPath := strings.Join(splittedPath[1:], "/")

		childNode := new(sessionproto.SessionLegacyNode)
		childNode.Name = name
		childNode.IsFile = false
		parent.Children = append(parent.Children, childNode)

		buildLegacyNode(childNode, nextPath)
	} else {
		childNode := new(sessionproto.SessionLegacyNode)
		childNode.IsFile = true
		childNode.Name = path
		parent.Children = append(parent.Children, childNode)
	}
}

func (manager *SessionManager) BuildSessionTreeFor(sessionId string) (*sessionproto.SessionTree, error) {
	if !manager.SessionExists(sessionId) {
		return nil, fmt.Errorf("no such session")
	}

	sessionTree := new(sessionproto.SessionTree)
	sessionTree.FilePaths = make(map[string]string)

	entryPath := manager.workdir + "/" + sessionId
	workdirs, err := ioutil.ReadDir(entryPath)
	if err != nil {
		return nil, err
	}
	
	err = persistEntity(entryPath, entryPath, workdirs, sessionTree.FilePaths, &sessionTree.FilesOrder)

	return sessionTree, err
}

func EncodeContent(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}
func DecodeContent(content string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(content)
}

func persistEntity(rootDir string, currentPath string, files []fs.FileInfo, persistTo map[string]string, ordering *[]string) error {
	for _, file := range files {
		fileRelativePath := currentPath + "/" + file.Name()

		if file.Name() == "target" || file.Name() == "output.log" {
			continue
		}

		if !file.IsDir() {

			fileContents, err := ioutil.ReadFile(fileRelativePath)
			if err != nil {
				return err
			}

			noRootPath := strings.Split(fileRelativePath, "/")
			persistKey := strings.Join(noRootPath[1:], "/")
			persistTo[persistKey] = EncodeContent(fileContents)

			// *ordering = append(*ordering, persistKey)
			*ordering = append([]string { persistKey }, *ordering...)
			continue
		}

		nextCurrentPath := currentPath + "/" + file.Name()
		innerFiles, err := ioutil.ReadDir(nextCurrentPath)
		if err != nil {
			return err
		}
		
		err = persistEntity(rootDir, nextCurrentPath, innerFiles, persistTo, ordering)
		if err != nil {
			return err
		}
	}
	return nil
}


func (manager *SessionManager) ConsumeSessions() (*SessionManager, error) {
	workdirs, err := ioutil.ReadDir(manager.workdir)
	if err != nil {
		return nil, err
	}

	var sessionWorkdirs []string

	for _, workspace := range workdirs {
		if !workspace.IsDir() {
			continue
		}

		sessionWorkdirs = append(sessionWorkdirs, workspace.Name())

		manager.sessions[workspace.Name()] = &sessionproto.Session{
			Ttl: DefaultSessionTTL,
			SessionId: workspace.Name(),
			// CreatedAt: uint64(time.Now().Unix()),
		}
	}

	manager.sessionWorkdirs = sessionWorkdirs


	return manager, nil
}

func (manager *SessionManager) newSessionId() string {
	sessionId := NewSessionID()

	if manager.sessions[sessionId.Encoded()] != nil {
		return manager.newSessionId()
	}
	return sessionId.Encoded()
}

func (manager *SessionManager) dropFSSession(sessionId string) error {
	destinationDir := manager.workdir + "/" + sessionId
	err := os.RemoveAll(destinationDir)

	return err
}

func (manager *SessionManager) createFSSession(sessionId string) error {
	destinationDir := manager.workdir + "/" + sessionId
	err := os.Mkdir(destinationDir, fs.ModePerm)

	if err != nil {
		return err
	}

	cmd := exec.Command("cp", "-r", manager.templatePath + "/", destinationDir)
	fmt.Printf("in: %v; out: %v \n", manager.templatePath + "/", destinationDir)

	output, err := cmd.Output()
	if err != nil {
		log.Printf("output: %v \n", string(output))
		log.Printf("Command finished with error: %v", err)
		return err
	}

	return nil
}

func (manager *SessionManager) DropSession(sessionId string) error {
	if manager.sessions[sessionId] == nil {
		return fmt.Errorf("no such session to drop")
	}

	manager.sessions[sessionId] = nil
	err := manager.dropFSSession(sessionId)

	return err
}

func (manager *SessionManager) SessionExists(sessionId string) bool {
	return manager.sessions[sessionId] != nil
}

func (manager *SessionManager) NewSession() (*sessionproto.Session, error) {
	// session
	sessionId := manager.newSessionId()
	session := &sessionproto.Session{
		Ttl: DefaultSessionTTL,
		SessionId: sessionId,
		CreatedAt: uint64(time.Now().Unix()),
	}

	manager.sessions[sessionId] = session
	err := manager.createFSSession(sessionId)

	return session, err
}

func NewSessionManager() *SessionManager {
	sessionManager := &SessionManager{}
	sessionManager.sessions = make(map[string]*sessionproto.Session)

	return sessionManager
}

func NewSessionManagerAt(dir string, templatePath string) (*SessionManager, error) {
	res, err := NewSessionManager().
		SetMeta(SessionManagerMeta{ WorkdirPath: dir, TemplatePath: templatePath }).
		ConsumeSessions()
	return res, err
}

type SessionRouter struct {
	
}