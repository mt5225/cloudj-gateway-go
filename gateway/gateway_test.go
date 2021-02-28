package gateway

import (
	"encoding/json"
	"testing"
)

type Msg struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

func TestFlask01(t *testing.T) {
	msg := Msg{
		Width:  10,
		Height: 20,
		Color:  "blue",
		Open:   false,
	}
	b, _ := json.Marshal(msg)
	opts := &Opts{
		Body:     b,
		Method:   "POST",
		Endpoint: "http://localhost:5000",
	}
	resp, err := Create(opts)
	if err != nil {
		t.Error("test with mock flask failed")
	}

	if resp.HostName != "10-225-135-226" {
		t.Error(resp)
	}

	if resp.Success == false {
		t.Error("test ifSuccess failed")
	}
}
