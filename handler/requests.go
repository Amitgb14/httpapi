package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	config "github.com/Amitgb14/httpapi/config"
	"github.com/Amitgb14/httpapi/reports"
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

func TestReponse(resp *Response, status []int, response string) bool {
	var checkStatus bool = false
	for _, code := range status {
		if resp.Status == code {
			checkStatus = true
		}
	}
	if !checkStatus {
		var msg = " != "
		if len(status) > 1 {
			msg = " not in " + fmt.Sprintf("%v", status)
		} else {
			msg += fmt.Sprintf("%v", status)[1:4]
		}
		log.Printf("\tFailed: %s\n", strconv.Itoa(resp.Status) + msg)
		return checkStatus
	}
	if response != "" {
		if resp.Body != response {
			log.Printf("\tFailed: %s\n", string(resp.Body) + " != " + response)
			return checkStatus
		}
	}

	return checkStatus
}
// NewRequests make new request
func NewRequests(data *config.Parameter, test bool) error{
	var req Request
	var err error

	var response = Response{}
	var report = reports.Reports{}

	client := req.NewClient()
	protocol := "http://"
	if data.SSL {
		protocol = "https://"
	}
	host := protocol + data.Host
	if data.Port != 0 {
		host += ":" + strconv.Itoa(data.Port)
	}
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
		if data.Token != "" {
			request.Header.Add("Authorization", data.Token)
		}

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
			status_ := []int{}
			var i int
			for i < len(status) {
				if status[i] != "" {
					sts, _ := strconv.Atoi(status[i])
					status_ = append(status_, sts)
				}
				i++
			}
			repo := make(map[string][]string)
			testStatus := TestReponse(&response, status_, content["resps"])
			report.TotalCount++
			var flgs = "Failed"
			if testStatus {
				report.Pass++
				flgs = "Passed"
			}

			t := []string{flgs, method}
			repo[url] = t
			report.TestName = append(report.TestName, repo)
		} else {
			log.Printf("Status: %d", response.Status)
			log.Printf("Text: %s", response.Body)
			log.Printf("----------------------------------------------------")

		}

	}
	if test{
		reports.GeneratesReport(&report)
	}
	return nil
}
