package internal

import (
	"context"
	"errors"
	"log"
	"sync"
)

var (
	ErrNoRow = errors.New("no row in struct")
)

type BackUp struct {
	Services []Services
}

type Config struct {
	ctx context.Context
	wg  *sync.WaitGroup
}

type Services interface {
	UpdateSettings(Settings)
	GetSettings() Settings
	Run(*Config) error
}

func (b *BackUp) BackUpFiles() error {
	config := NewConfig()
	for i := 0; i < len(b.Services); i++ {
		config.wg.Add(1)
		go func(s Services, wg *sync.WaitGroup) {
			err := s.Run(config) // launching backup one of the protocol services
			if err != nil {
				log.Println(err)
			}
			wg.Done()

		}(b.Services[i], config.wg)
	}

	config.wg.Wait()
	return nil
}

func NewBackup() *BackUp {
	return new(BackUp)
}

func NewConfig() *Config {
	c := new(Config)
	c.ctx = context.Background()
	c.wg = new(sync.WaitGroup)
	return c
}
