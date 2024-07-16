package internal

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Service struct {
	S3 *S3 `yaml:"s3"`
	//SFTP SFTP `yaml:"sftp"`
}

type BCat struct {
	Services map[string]Service `yaml:"services"`
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
