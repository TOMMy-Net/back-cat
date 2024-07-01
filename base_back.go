package main

import (
	"errors"
	"os"

	"path/filepath"

	"github.com/TOMMy-Net/back-cat/internal"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	ErrNoRow = errors.New("no row in struct")
)

func (s Service) BackUpFiles(ServiceName string) error {
	pathSL := filepath.SplitList(s.Path)
	fileStat, err := os.Stat(filepath.Join(pathSL[0]))
	if err != nil {
		return err
	}

	sess := AWSsession(&s.S3)
	uploader := internal.InitPartUploader(sess)

	file, err := os.Open(filepath.Join(pathSL[0]))
	if err != nil {
		return err
	}

	if fileStat.IsDir() {

	} else {

		UD := internal.UploadedData{
			Data:     file,
			Bucket:   s.S3.Bucket,
			Key:      filepath.ToSlash(filepath.Join(pathSL[1], fileStat.Name())),
			Uploader: uploader,
		}
		_, err = UD.UploadData()
		if err != nil {
			return err
		}

	}

	return nil
}

func AWSsession(s3 *S3Serv) *session.Session {
	var awsConfig *aws.Config
	if s3.AccessKey == "" || s3.SecretKey == "" {
		//load default credentials
		awsConfig = &aws.Config{
			Endpoint: &s3.S3url,
			Region:   &s3.Region,
		}
	} else {
		awsConfig = &aws.Config{
			Endpoint:    &s3.S3url,
			Region:      &s3.Region,
			Credentials: credentials.NewStaticCredentials(s3.AccessKey, s3.SecretKey, ""),
		}
	}

	sess := session.Must(session.NewSession(awsConfig))
	return sess
}
