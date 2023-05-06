package api

import (
	"fmt"
	"github.com/vanerjkeee/go-service/service"
	"net/http"
)

type Controller struct {
	taskChannel   chan HttpRequestTask
	httpClient    service.HttpClient
	resultChannel chan service.HttpResponse
	storage       service.Storage
}

func Start(config *service.Config) {
	c := Controller{}
	c.taskChannel = make(chan HttpRequestTask, config.Tasks.ChannelSize)
	c.resultChannel = make(chan service.HttpResponse, config.Tasks.ChannelSize)
	c.httpClient = service.NewHttpClient(c.resultChannel, config.HttpClient.ChannelSize, config.HttpClient.WorkersCount)
	c.createTaskWorkers(config.Tasks.WorkersCount)
	c.createResultWorkers(config.Tasks.WorkersCount)

	http.HandleFunc("/add", c.AddRequest)
	http.HandleFunc("/status", c.StatusRequest)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}

type HttpRequestTask struct {
	Url   string
	Count int
}

func (c *Controller) createTaskWorkers(count int) {
	for i := 0; i < count; i++ {
		go c.taskWorker()
	}
}

func (c *Controller) taskWorker() {
	for {
		newTask := <-c.taskChannel
		for i := 0; i < newTask.Count; i++ {
			c.httpClient.Request(newTask.Url)
		}

		c.storage.IncTotalRequestCount(newTask.Count)
	}
}

func (c *Controller) createResultWorkers(count int) {
	for i := 0; i < count; i++ {
		go c.resultWorker()
	}
}

func (c *Controller) resultWorker() {
	for {
		result := <-c.resultChannel
		if result.Error != nil || result.Response.StatusCode >= 400 {
			c.storage.IncErrorRequestCount()
		} else {
			c.storage.IncSuccessRequestCount()
		}
	}
}
