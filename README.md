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
Host: localhost
Port: 8080
Thread: 1
Loop: 1
Requests:
  - Path: /status
    Method: Get
    Status: 200
```

### test2.yaml
```yaml
Host: localhost
Port: 8080
Thread: 1
Loop: 1
Include:
  - a.yaml
  - b.yaml
```

### a.yaml
```yaml
Requests:
  - Path: /status
    Method: Get
    Status: 200

  - Path: /create
    Method: POST
    Status: 201
```

### b.yaml
```yaml
Requests:
  - Path: /health
    Method: Get
    Status: 200
```


