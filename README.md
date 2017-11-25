# httpapi

HTTP API client for Go

### Download configfile
```
$ go get github.com/Amitgb14/configfile
```

### Build httpapi
```
$ git clone https://github.com/Amitgb14/httpapi.git
$ cd httpapi
$ go build
$ ./httpapi test.yaml
```

### test.yaml
```yaml
host: localhost
port: 8080
thread: 1
loop: 1
requests:
  - path: /status
    method: GET
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
    method: GET
    status: 200

  - path: /create
    method: POST
    status: 201
```

### b.yaml
```yaml
requests:
  - path: /health
    method: GET
    status: 200
```


