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

func TestReponse(resp *Response, status [2]int, response string) {
	var checkStatus bool = false
	for _, code := range status {
		if resp.Status == code {
			checkStatus = true
		}
	}
	if !checkStatus {
		log.Printf("\tFailed: %s\n", strconv.Itoa(resp.Status) + " != " + strings.Trim(strings.Replace(fmt.Sprint(status), " ", ", ", -1), "[]"))
		return
	}
	if response != "" {
		if resp.Body != response {
			log.Printf("\tFailed: %s\n", string(resp.Body) + " != " + response)
			return
		}
	}


}
// NewRequests make new request
func NewRequests(data *config.Parameter, test bool) error{
	var req Request
	var err error

	var response = Response{}
	client := req.NewClient()

	host := "http://" + data.Host + ":" + strconv.Itoa(data.Port)
	log.Printf("%s\n", host)
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

		log.Printf("%s %s", method, url)
		if method == "GET" {
			request, err = http.NewRequest(method, url, nil)
		} else {
			if content["data"] != "" {
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

		if test {
			status := strings.Split(content["status"], ",")
			status_ := [2]int{}
			one, _ := strconv.Atoi(status[0])
			status_[0] = one
			if len(status) > 1 {
				two, _ := strconv.Atoi(status[1])
				status_[1] = two
			}
			TestReponse(&response, status_, content["resps"])
		} else {
			log.Printf("Status: %d", response.Status)
			log.Printf("Text: %s", response.Body)
			log.Printf("----------------------------------------------------")

		}

	}
	return nil
}
