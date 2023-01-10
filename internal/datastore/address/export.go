package address

import (
	"context"

	"be/internal/model"
)

func (s *Storage) UploadFileExportedAddresss(ctx context.Context, attachment model.Attachment) (*model.Attachment, error) {
	key, err := s.fileStorage.AddTemporaryFile(ctx, attachment.Source(), attachment.Filename, attachment.Mime)
	if err != nil {
		return nil, err
	}

	attachment.Key = key

	uri, err := s.fileStorage.Link(key, s.linkTimeLife)
	if err != nil {
		return nil, err
	}

	attachment.SetURL(uri)

	return &attachment, err
}
