package storage

import (
	"path/filepath"
)

type factoryDisk struct {
	dirPath    string
	isPlayBack bool
}

// NewFactoryDisk allocates a disk-backed factory.
func NewFactoryDisk(dirPath string, isPlayBack bool) Factory {
	return &factoryDisk{
		dirPath:    dirPath,
		isPlayBack: isPlayBack,
	}
}

// NewFile implements Factory.
func (s *factoryDisk) NewFile(fileName string) (File, error) {
	return newFileDisk(filepath.Join(s.dirPath, fileName), s.isPlayBack)
}
