package backcat

import (
	"log"
	"reflect"

	"github.com/TOMMy-Net/back-cat/internal"
)

func App() {
	data := internal.ReadCat()
	backUp := internal.NewBackup()
	val := reflect.ValueOf(data.Services)
	for i := 0; i < val.NumField(); i++ {
		v, ok := val.Field(i).Interface().(internal.Services) 
		if ok && v != nil{
			backUp.Services = append(backUp.Services, v)
		}
	}

	if err := backUp.BackUpFiles(); err != nil {
		log.Fatal(err)
	}

}
