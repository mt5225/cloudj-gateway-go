package gateway

import (
	"encoding/json"
	"testing"
)

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

func TestFlask01(t *testing.T) {
	box := Box{
		Width:  10,
		Height: 20,
		Color:  "blue",
		Open:   false,
	}
	b, _ := json.Marshal(box)
	opts := &gatewayOpts{
		body:     b,
		method:   "POST",
		endpoint: "http://localhost:5000/dummy",
	}
	_, err := Create(opts)
	if err != nil {
		t.Error("test flask dummy failed")
	}
}
