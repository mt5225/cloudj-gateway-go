package gateway

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Opts data struct
type Opts struct {
	Body     []byte
	Method   string
	Endpoint string
}

// Create resource via gateway
func Create(opts *Opts) (map[string]interface{}, error) {
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

	// code here to read body of response and return it, and convert response json to map
	responseMap := make(map[string]interface{})
	json.Unmarshal(body, &responseMap)

	defer response.Body.Close()

	return responseMap, err
}
