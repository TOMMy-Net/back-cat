package main

import (
	"log"

	backcat "github.com/TOMMy-Net/back-cat/cmd/back-cat"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Fatalf("PANIC: %+v", e)
		}
	}()
	
	backcat.App()
}