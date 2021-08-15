package cache


type CacheManager struct {}

func (cm *CacheManager) Persist() {}

type SessionManager struct {
	cache *CacheManager
}


type CodeSession struct {
	Node     FileSystemNode
	// New bool
}

func (sm *SessionManager) fetch(path string) *CodeSession  {
	return nil
}


func (sm *SessionManager) Fetch(inputBody []byte) (*CodeSession, error) {
	
	// return &CodeSession{}
	return nil, nil
}

func NewCache() *CacheManager {
	return &CacheManager{}
}

func NewSession() *SessionManager {
	return &SessionManager{ cache: &CacheManager{} }
}

type FileSystemNode struct {
	Name string 
	IsFile bool 

	Children *[]FileSystemNode 
}