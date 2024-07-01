package internal

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadedData struct {
	Data     io.Reader
	Bucket   string
	Key      string
	Uploader *s3manager.Uploader
}

func (u UploadedData) UploadData() (*s3manager.UploadOutput, error) {

	result, err := u.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(u.Key),
		Body:   u.Data,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func InitPartUploader(sess *session.Session) *s3manager.Uploader {

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // 2 mb
		u.Concurrency = 4
		u.LeavePartsOnError = true
	})

	return uploader
}
