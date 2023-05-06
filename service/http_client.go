package service

import (
	"net/http"
)

type HttpResponse struct {
	Request  string
	Response *http.Response
	Error    error
}

type HttpClient struct {
	internalChannel chan string
	ResultChannel   chan HttpResponse
}

func NewHttpClient(resultChannel chan HttpResponse, channelSize, workersCount int) HttpClient {
	hc := HttpClient{}
	hc.internalChannel = make(chan string, channelSize)
	hc.ResultChannel = resultChannel
	hc.createWorkers(workersCount)
	return hc
}

func (hc *HttpClient) Request(url string) {
	hc.internalChannel <- url
}

func (hc *HttpClient) createWorkers(count int) {
	for i := 0; i < count; i++ {
		go hc.worker()
	}
}

func (hc *HttpClient) worker() {
	for {
		url := <-hc.internalChannel
		hc.processTask(url)
	}
}

func (hc *HttpClient) processTask(url string) {
	resp, err := http.Get(url)
	hc.ResultChannel <- HttpResponse{url, resp, err}
}
