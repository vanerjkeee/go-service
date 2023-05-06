package service

import (
	"testing"
)

func TestNewHttpClient(t *testing.T) {
	resultChannel := make(chan HttpResponse, 10)
	httpClient := NewHttpClient(resultChannel, 1, 1)

	if httpClient.ResultChannel != resultChannel {
		t.Fatalf(`NewHttpClient channel mismatch`)
	}
}

func TestHttpClientRequest(t *testing.T) {
	url := "https://google.com"
	resultChannel := make(chan HttpResponse, 10)
	httpClient := NewHttpClient(resultChannel, 1, 1)
	httpClient.Request(url)

	response := <-resultChannel
	if response.Request != url {
		t.Fatalf(`NewHttpClient url mismatch %s %s`, url, response.Request)
	}
}
