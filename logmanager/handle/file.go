package handle

import (
	"fmt"
	"os"
	"sync"

	"github.com/StrataLinks/kaizen-Utils/internal/logger"
)

type FileHandle struct {
	file  *os.File
	mutex sync.Mutex
}

func NewFileHandle(filename string) *FileHandle {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	return &FileHandle{file: file}
}

func (f *FileHandle) Log(level logger.LogLevel, message string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	_, err := f.file.WriteString(fmt.Sprintf("%s: %s\n", level, message))
	if err != nil {
		panic(err)
	}
}

func (f *FileHandle) Close() error {
	return f.file.Close()
}
