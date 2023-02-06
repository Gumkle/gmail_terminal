package filesystem

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type PathString string

type FileSystem struct {
	root string
}

func NewFileSystem(appDataRoot PathString) (*FileSystem, error) {
	normalizedRoot := strings.Trim(strings.TrimPrefix(string(appDataRoot), "/tmp/"), "/")
	normalizedRoot = fmt.Sprintf("/tmp/%s", normalizedRoot)
	err := os.MkdirAll(normalizedRoot, fs.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create app data directory: %w", err)
	}
	return &FileSystem{normalizedRoot}, nil
}

func (f *FileSystem) Save(destination, contents string) error {
	fullPath := fmt.Sprintf("%s/%s", f.root, destination)
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to open file for save: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(contents)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

func (f *FileSystem) Remove(path string) error {
	fullPath := fmt.Sprintf("%s/%s", f.root, path)
	err := os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("failed to remove file: %w", err)
	}
	return nil
}
