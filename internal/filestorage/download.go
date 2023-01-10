package filestorage

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

var ErrDownloadFailed = errors.New("file download failed")

type FileLoader struct{}

func (FileLoader) DownloadFile(ctx context.Context, fileURL string) (IFile, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)

	if _, err = io.Copy(buffer, resp.Body); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(ErrDownloadFailed, "s3 respond: %s", resp.Status)
	}

	defer func() { resp.Body.Close() }()

	f := File{}
	f.SetSource(buffer)

	return &f, nil
}

type File struct {
	source *bytes.Buffer
}

func (f *File) Source() io.Reader {
	return f.source
}

func (f *File) SetSource(source *bytes.Buffer) {
	f.source = source
}

func (f *File) Size() int {
	return f.source.Len()
}
