package main

import (
	"log"
	"os"

         config "github.com/Amitgb14/configfile"
	"github.com/Amitgb14/httpapi/handler"
)

func main() {
	var test bool = false
	for _, file := range os.Args[1:] {
		data, err := config.Yaml(file)
		if err != nil {
			log.Fatalf("Read config: %v", err)
		}
		err = handler.NewRequests(data, test)
		if err != nil {
			log.Fatalf("%v", err)
		}

	}
}
