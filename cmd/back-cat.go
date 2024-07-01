package cmd

import (
	"log"
	"os"

	"github.com/TOMMy-Net/back-cat/internal"
	yaml "gopkg.in/yaml.v3"
)

type S3Serv struct {
	S3url     string `yaml:"s3-url"`
	AccessKey string `yaml:"s3-key"`
	SecretKey string `yaml:"s3-secret"`
	Bucket    string `yaml:"s3-bucket"`
	Region    string `yaml:"s3-region"`
	Path      string `yaml:"s3-path"`
}

type Service struct {
	Path    string `yaml:"path"`
	S3      S3Serv `yaml:"s3"`
	Time    string `yaml:"time"`
	Archive bool   `yaml:"archive"`
}

type BCat struct {
	Paths map[string]Service `yaml:"paths"`
}

func App() {
	data := ReadCat()
	for i, v := range data.Paths {
		backUp := internal.BackUp{
			Path: v.Path,
			S3: internal.S3{
				Url:    v.S3.S3url,
				Access: v.S3.AccessKey,
				Secret: v.S3.SecretKey,
				Bucket: v.S3.Bucket,
				Region: v.S3.Region,
				Path:   v.S3.Path,
			},
			Name:    i,
			Archive: v.Archive,
		}
		log.Println(backUp.BackUpFiles())
	}
	log.Println(data)
}

func ReadCat() BCat {
	file, err := os.ReadFile("backCat.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var data = BCat{}

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
