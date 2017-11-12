package main

import (
	"log"
	"os"

	"github.com/Amitgb14/httpapi/config"
	"github.com/Amitgb14/httpapi/handler"
)

func main() {

	for _, file := range os.Args[1:] {
		data, err := config.Yaml(file)
		if err != nil {
			log.Fatalf("Read config: %v", err)
		}
		// fmt.Println(data)
		req := handler.Request{}
		req.NewRequests(data)
		//if err != nil {
		//	log.Fatalf("%v", err)
		//}
		//fmt.Println(report)
	}
}
