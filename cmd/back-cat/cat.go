package backcat

import (
	"github.com/TOMMy-Net/back-cat/internal"
	"log"
	"reflect"
)

func App() {
	data := internal.ReadCat()
	backUp := internal.NewBackup()

	for _, v := range data.Services {
		val := reflect.ValueOf(v)
		for i := 0; i < val.NumField(); i++ {
			v, ok := val.Field(i).Interface().(internal.Services)
			if ok && v != nil {
				backUp.Services = append(backUp.Services, v)
			}
		}
	}

	if err := backUp.BackUpFiles(); err != nil {
		log.Fatal(err)
	}

}
