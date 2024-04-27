package service

type FileManager struct {
	path string
}

func NewFileManager(path string) *FileManager {
	return &FileManager{
		path: path,
	}
}

func (fm *FileManager) GetPath() string {
	return fm.path
}

func (fm *FileManager) Ls() []string {
	return []string{}
}
