package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Amitgb14/httpapi/config"
)

func main() {

	for _, file := range os.Args[1:] {
		data, err := config.Yaml(file)
		if err != nil {
			log.Fatalf("Read config: %v", err)
		}
		fmt.Println(*data)
	}
}
