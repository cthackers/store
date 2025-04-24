package fs

import (
	"io"
	"os"
)

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	Slice(start, end int64) ([]byte, error)
}

type Filesystem interface {
	Open(name string, flags int, perm os.FileMode) (File, error)
	Stat(name string) (os.FileInfo, error)
	Remove(name string) error
	Rename(oldname, newname string) error
	ReadDir(name string) ([]os.DirEntry, error)
	MakeDir(name string, perm os.FileMode) error
}
