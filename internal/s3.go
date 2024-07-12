package internal

import (
	"context"
	"io"

	"path/filepath"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws"
)

type S3 struct {
	Url      string   `yaml:"s3-url" validate:"required"`
	Access   string   `yaml:"s3-key" validate:"required"`
	Secret   string   `yaml:"s3-secret" validate:"required"`
	Bucket   string   `yaml:"s3-bucket" validate:"required"`
	Region   string   `yaml:"s3-region" validate:"required"`
	Path     string   `yaml:"s3-path" validate:"required"`
	Settings Settings `yaml:"settings"`
}

type S3UploadedData struct {
	Data     io.Reader
	Bucket   string
	Key      string
	Uploader *s3manager.Uploader
}

func (u *S3UploadedData) UploadData(ctx context.Context) error {

	_, err := u.Uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(u.Key),
		Body:   u.Data,
	})

	if err != nil {
		return err
	}
	return nil
}

func NewS3UploadedData() *S3UploadedData {
	return &S3UploadedData{}
}

func (s3 S3) awsSession() (*session.Session, error) {
	var awsConfig *aws.Config
	if s3.Access == "" || s3.Secret == "" {
		//load default credentials
		awsConfig = &aws.Config{
			Endpoint: &s3.Url,
			Region:   &s3.Region,
		}
	} else {
		awsConfig = &aws.Config{
			Endpoint:    &s3.Url,
			Region:      &s3.Region,
			Credentials: credentials.NewStaticCredentials(s3.Access, s3.Secret, ""),
		}
	}
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return &session.Session{}, err
	}
	must := session.Must(sess, err)
	return must, err
}

func (s3 S3) GetUploader() (*S3UploadedData, error) {
	sess, err := s3.awsSession()
	if err != nil {
		return &S3UploadedData{}, err
	}
	up := InitS3PartUploader(sess)
	data := NewS3UploadedData()
	data.Uploader = up
	data.Bucket = s3.Bucket
	return data, nil
}

func (s3 S3) GetSettings() Settings {
	return s3.Settings
}

func (u *S3UploadedData) SetData(data io.Reader) {
	u.Data = data
}

func (u *S3UploadedData) SetPath(path ...string) {
	u.Key = filepath.ToSlash(filepath.Join(path...))
}

func (u *S3UploadedData) GetPath() string {
	return u.Key
}

func InitS3PartUploader(sess *session.Session) *s3manager.Uploader {

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // 5 mb
		u.Concurrency = 50
		u.LeavePartsOnError = true
	})

	return uploader
}

func (s S3) Run(c *Config) error {
	if err := validate(s); err != nil {
		return err
	} else {
		up, err := s.GetUploader()
		if err != nil {
			return err
		}
		settings := s.GetSettings()
		settings.WalkandUpload(up, c)
		return nil
	}
}
