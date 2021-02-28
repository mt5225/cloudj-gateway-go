package gateway

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// Opts data struct
type Opts struct {
	Body     []byte
	Method   string
	Endpoint string
}

// ResponseMap from api gateway
type ResponseMap struct {
	ServerID string
	HostName string
	Success  bool
}

// Create resource via gateway
func Create(opts *Opts) (*ResponseMap, error) {
	// initialize endpoint
	endpoint := opts.Endpoint

	// initialize http client
	client := &http.Client{}

	// body into bytes.Buffer pointer
	var bodyBuffer *bytes.Buffer
	bodyBuffer = bytes.NewBuffer(opts.Body)

	// body always request body or nil for reads
	request, err := http.NewRequest(opts.Method, endpoint, bodyBuffer)
	request.Header.Set("Content-Type", "application/json")

	// code here to error handle

	// initiate request for response
	response, err := client.Do(request)

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// code here to error handle
	defer response.Body.Close()

	return fusioncloudGetServerID(body), err
}

func fusioncloudGetServerID(msg []byte) *ResponseMap {
	json := string(msg)
	resp := new(ResponseMap)
	resp.Success = gjson.Get(json, "isSuccess").Bool()
	resp.ServerID = gjson.Get(json, "resultObject.resultMap.servers.0.tenant_id").String()
	resp.HostName = gjson.Get(json, "resultObject.resultMap.servers.0.metadata.hostname").String()
	return resp
}
