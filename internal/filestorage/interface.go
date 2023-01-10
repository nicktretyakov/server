package filestorage

import (
	"context"
	"io"
	"time"
)

type IFileStorage interface {
	AddFile(ctx context.Context, r io.Reader, filename, mime string) (string, error)
	AddTemporaryFile(ctx context.Context, r io.Reader, filename, mime string) (string, error)
	RemoveFile(ctx context.Context, key string) error
	RenameFile(ctx context.Context, key, newFilename, mime string) error
	Link(key string, lifeTime time.Duration) (string, error)
}

type IFileLoader interface {
	DownloadFile(ctx context.Context, fileURL string) (IFile, error)
}

type IFile interface {
	Source() io.Reader
	Size() int
}
