package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Host, port string

func demoHandler(w http.ResponseWriter, req *http.Request) {
	structObj := Host
	s, _ := json.Marshal(structObj)
	_, err := w.Write(s)
	if err != nil {
		return
	}
}

func demoServer() {
	http.HandleFunc("/", demoHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}

func main() {
	port = os.Args[1]
	Host = os.Args[2]
	if len(Host) < 1 || len(port) < 1 {
		panic("host or port is empty")
	}
	demoServer()
}
