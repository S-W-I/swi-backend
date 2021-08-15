package fs

import "fmt"

const (
	NoSuchFileError = "no such file"
)

type FileSystemRepr struct {
	// mappings of the dir
	fsmappings map[string][]byte
}

func (fsr *FileSystemRepr) entityExists(path string) bool {
	return len(fsr.fsmappings[path]) > 0
}

func (fsr *FileSystemRepr) UpdateFile(path string, data []byte) error {
	if !fsr.entityExists(path) {
		return fmt.Errorf(NoSuchFileError)
	}

	fsr.fsmappings[path] = data
	return nil
}

