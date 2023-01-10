package filestorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type service struct {
	s3Session      *session.Session
	filesBucket    string
	tmpFilesPrefix string
}

type Config struct {
	Region    string
	Endpoint  string
	AccessKey string
	SecretKey string

	Bucket               string
	TemporaryFilesPrefix string
}

func New(cfg Config) (IFileStorage, error) {
	s3Session, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
		Endpoint:    aws.String(cfg.Endpoint),
		Region:      aws.String(cfg.Region),
	})
	if err != nil {
		return nil, err
	}

	return &service{
		s3Session:      s3Session,
		filesBucket:    cfg.Bucket,
		tmpFilesPrefix: cfg.TemporaryFilesPrefix,
	}, nil
}
