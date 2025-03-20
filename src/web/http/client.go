package http

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	nurl "net/url"
	"strings"
	"time"
)

const (
	HttpClientTimeout = 30 * time.Second
)

func call(method, url string, h map[string]string, p map[string]string, body interface{}) (error, []byte) {
	reqStr, err := json.Marshal(body)
	if err != nil {
		return err, nil
	}
	reader := strings.NewReader(string(reqStr))
	u, err := nurl.Parse(url)
	if err != nil {
		return err, nil
	}
	var client *http.Client
	if u.Scheme == "http" {
		client = &http.Client{
			Timeout: HttpClientTimeout,
		}
	} else if u.Scheme == "https" {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: HttpClientTimeout,
		}
	} else {
		return fmt.Errorf("cann't suppot the scheme %s", u.Scheme), nil
	}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest(method, u.String(), reader)
	if err != nil {
		return err, nil
	}
	req.Header.Set("content-type", "application/json")
	for k, v := range h {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		return err, nil
	}
	respByteArr, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	// GET http requests don't log
	if method != http.MethodGet {
		log.Printf("url: %s\nmethod: %s\nheader: %s\nparams: %s\nbody: %s\nresponse: %s\n",
			url, method, h, p, string(reqStr), string(respByteArr))
	}
	return nil, respByteArr
}
