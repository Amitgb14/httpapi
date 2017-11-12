package handler

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	config "github.com/Amitgb14/httpapi/config"
)

// Set constant
const (
	USERAGENT        = "go-client/1.1 ..."
	contentTypePlain = "text/plain"
	contentTypeJson  = "application/json"
)

// Request at Client
type Request struct {
	httpClient *http.Client
	method     string
	url        string
}

// NewClient make new client
func (req *Request) NewClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	req.httpClient = &http.Client{Transport: tr}
	return req.httpClient
}

// NewRequests make new request
func (req *Request) NewRequests(data *config.Parameter) {

	var err error
	client := req.NewClient()

	host := "http://" + data.Host + ":" + strconv.Itoa(data.Port)
	fmt.Println(host)
	values := bytes.NewBuffer([]byte(``))
	for _, content := range data.Requests {
		var request *http.Request

		method := strings.ToUpper(content["Method"])
		contenttype := content["Content-Type"]
		if contenttype == "" {
			contenttype = contentTypePlain
		} else if contenttype == "application/json" {
			contenttype = contentTypeJson
		}

		if method == "GET" {
			request, err = http.NewRequest(method, host, nil)
		} else {
			request, err = http.NewRequest(method, host, values)
		}

		if err != nil {
			log.Fatal(err)
		}

		request.Header.Add("User-Agent", USERAGENT)
		request.Header.Set("Content-Type", contenttype)

		resp, err := client.Do(request)
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
		// datas, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(datas))
		// response.resp = resp

		// return &response, nil

	}
}
