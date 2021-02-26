package gateway

import (
	"bytes"
	"net/http"
)

type gatewayOpts struct {
	body   []byte
	method string
}

// Invoke gateway with message
func Invoke(opts *gatewayOpts) (map[string][]string, error) {
	// initialize endpoint
	endpoint := "https://api.pizza.com/request/create"

	// initialize http client
	client := &http.Client{}

	// body into bytes.Buffer pointer
	var bodyBuffer *bytes.Buffer
	bodyBuffer = bytes.NewBuffer(opts.body)

	// body always request body or nil for reads
	request, err := http.NewRequest(opts.method, endpoint, bodyBuffer)

	// code here to error handle

	// initiate request for response
	response, err := client.Do(request)

	// code here to error handle

	// code here to read body of response and return it, and convert response json to map
	responseMap := make(map[string][]string)

	defer response.Body.Close()

	return responseMap, err
}
