package http

import (
	"log"
	"net/http"
	"testing"
)

func TestDemoServer(t *testing.T) {
	// demo server start
	endpoint := "0.0.0.0:8080"
	handlerMap := map[string]http.HandlerFunc{
		"/test": func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write([]byte("test"))
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	cancel := make(chan struct{})
	go demoServer(endpoint, handlerMap, cancel)
	err, res := call(http.MethodGet, "http://localhost:8080/health", nil, nil, nil)
	if err != nil {
		t.Errorf("call failed %v", err)
	} else {
		if string(res) != "ok" {
			t.Errorf("call failed, want ok, but %s", string(res))
		}
	}
	cancel <- struct{}{}
}
