package internal

import (
	"context"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type Uploader interface {
	UploadData(context.Context, io.Reader) error
	SetPath(...string)
	GetPath() string
}

type Settings struct {
	Time    string `yaml:"time"`
	Archive bool   `yaml:"archive"`
	Path    string `yaml:"path" validate:"required"`
	Name    string `yaml:"-"`
}

func (s Settings) WalkandUpload(uploader Uploader, c *Config) error {
	err := filepath.Walk(s.Path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Printf("%s ~ %s : %s", s.Name, path, err)
			return nil
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				log.Printf("%s ~ %s : %s", s.Name, path, err)
				return nil
			}
			defer file.Close()
			uploader.SetPath(path)
			err = uploader.UploadData(c.ctx, file)
			if err != nil {
				log.Printf("%s ~ %s : err uploading (%s)", s.Name, path, err)
			} else {
				log.Printf("%s ~ %s : done uploading", s.Name, path)
			}
		}

		return nil
	})

	return err
}
