package ressource

import (
	"io"
	"os"
)

type Adapter interface {
	Get() (io.Reader, error)
	Name() string
}

type FileAdapter struct {
	path string
	file *os.File
}

func (f *FileAdapter) openFile() error {
	if f.file == nil {
		file, err := os.Open(f.path)
		f.file = file
		return err
	}
	return nil
}

func (f *FileAdapter) Get() (reader io.Reader, err error) {
	err = f.openFile()
	reader = f.file
	return
}

func (f *FileAdapter) Name() string {
	return f.path
}

func NewFileAdapter(path string) *FileAdapter {
	return &FileAdapter{path: path}
}
