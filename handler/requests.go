package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	config "github.com/Amitgb14/httpapi/config"
	"io/ioutil"
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
func NewRequests(data *config.Parameter) error{
	var req Request
	var err error

	var response = Response{}
	client := req.NewClient()

	host := "http://" + data.Host + ":" + strconv.Itoa(data.Port)
	fmt.Println(host)
	payloads :=  strings.NewReader("")

	for _, content := range data.Requests {
		var request *http.Request

		url := host + content["path"]
		method := strings.ToUpper(content["method"])
		contenttype := content["type"]
		if contenttype == "" {
			contenttype = contentTypePlain
		} else if contenttype == "application/json" {
			contenttype = contentTypeJson
		}

		fmt.Println(method, url)
		if method == "GET" {
			request, err = http.NewRequest(method, url, nil)

		} else {
			if content["Data"] != "" {
				payloads = strings.NewReader(content["data"])
			}
			request, err = http.NewRequest(method, url, payloads)
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

		response.Status = resp.StatusCode
		datas, _ := ioutil.ReadAll(resp.Body)
		response.Body = string(datas)

		fmt.Println("Status: ", response.Status)
		fmt.Println("Text: ", response.Body)
		fmt.Println("----------------------------------------------------")
		// return &response, nil

	}
	return nil
}
