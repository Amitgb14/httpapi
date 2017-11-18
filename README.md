# httpapi

HTTP API client for Go


### Example
```go
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
```

### test.yaml
```yaml
host: localhost
port: 8080
thread: 1
loop: 1
requests:
  - path: /status
    method: Get
    status: 200
```

### test2.yaml
```yaml
host: localhost
port: 8080
thread: 1
loop: 1
include:
  - a.yaml
  - b.yaml
```

### a.yaml
```yaml
requests:
  - path: /status
    method: Get
    status: 200

  - path: /create
    method: POST
    status: 201
```

### b.yaml
```yaml
requests:
  - path: /health
    method: Get
    status: 200
```


