package filestorage

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"be/internal/lib"
)

func (s service) AddFile(ctx context.Context, r io.Reader, filename, mime string) (string, error) {
	key := lib.UUID().String()

	return key, s.upload(ctx, r, key, filename, mime)
}

func (s service) AddTemporaryFile(ctx context.Context, r io.Reader, filename, mime string) (string, error) {
	key := s.tmpFilesPrefix + lib.UUID().String()

	return key, s.upload(ctx, r, key, filename, mime)
}

func (s service) Link(key string, lifeTime time.Duration) (string, error) {
	sdkReq, _ := s3.New(s.s3Session).GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.filesBucket),
		Key:    aws.String(key),
	})

	uri, err := sdkReq.Presign(lifeTime)
	if err != nil {
		return "", err
	}

	return uri, nil
}

func (s service) RemoveFile(ctx context.Context, key string) error {
	_, err := s3.New(s.s3Session).DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.filesBucket),
		Key:    aws.String(key),
	})

	return err
}

func (s service) RenameFile(ctx context.Context, key, newFilename, mime string) error {
	_, err := s3.New(s.s3Session).CopyObjectWithContext(ctx, &s3.CopyObjectInput{
		CopySource:         aws.String(url.PathEscape(strings.Join([]string{s.filesBucket, key}, "/"))),
		Bucket:             aws.String(s.filesBucket),
		Key:                aws.String(key),
		ContentType:        aws.String(mime),
		MetadataDirective:  aws.String(s3.MetadataDirectiveReplace),
		ContentDisposition: aws.String(fmt.Sprintf(`filename="%s"`, url.QueryEscape(newFilename))),
	})

	return err
}

func (s service) upload(ctx context.Context, r io.Reader, key, filename, mime string) error {
	_, err := s3manager.NewUploader(s.s3Session).UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket:             aws.String(s.filesBucket),
		Key:                aws.String(key),
		Body:               r,
		ContentType:        aws.String(mime),
		ContentDisposition: aws.String(fmt.Sprintf(`filename="%s"`, url.QueryEscape(filename))),
	})

	return err
}
