package internal

import (
	"errors"
	"os"

	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	ErrNoRow = errors.New("no row in struct")
)

type S3 struct {
	Url    string
	Access string
	Secret string
	Bucket string
	Region string
	Path   string
}

type BackUp struct {
	Path    string
	S3      S3
	Archive bool
	Name    string
}

func (b *BackUp) BackUpFiles() error {

	fileStat, err := os.Stat(filepath.Join(b.Path))
	if err != nil {
		return err
	}

	sess := AWSsession(&b.S3)
	uploader := InitPartUploader(sess)

	file, err := os.Open(filepath.Join(b.Path))
	if err != nil {
		return err
	}

	if fileStat.IsDir() {

	} else {

		UD := UploadedData{
			Data:     file,
			Bucket:   b.S3.Bucket,
			Key:      filepath.ToSlash(filepath.Join(b.S3.Path, fileStat.Name())),
			Uploader: uploader,
		}
		_, err = UD.UploadData()
		if err != nil {
			return err
		}

	}

	return nil
}

func AWSsession(s3 *S3) *session.Session {
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

	sess := session.Must(session.NewSession(awsConfig))
	return sess
}
