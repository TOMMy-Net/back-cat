package main

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type S3Serv struct {
	S3url     string `yaml:"s3-url"`
	AccessKey string `yaml:"s3-key"`
	SecretKey string `yaml:"s3-secret"`
	Bucket    string `yaml:"s3-bucket"`
	Region    string `yaml:"s3-region"`
}

type Service struct {
	Path    string `yaml:"path"` // path:/s3-path
	S3      S3Serv `yaml:"s3"`
	Time    string `yaml:"time"`
	Archive bool   `yaml:"archive"`
}

type BCat struct {
	Paths map[string]Service `yaml:"paths"`
}

func main() {
	data := ReadCat()
	for i, v := range data.Paths {
		d:= v
		go d.BackUpFiles(i)
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

