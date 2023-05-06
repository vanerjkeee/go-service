package api

import (
	"encoding/json"
	"net/http"
)

type RequestAddTask struct {
	Url            string `json:"url"`
	RequestsNumber int    `json:"number_of_requests"`
}

type ResponseStatus struct {
	Total     int   `json:"total"`
	Success   int   `json:"success"`
	Error     int   `json:"error"`
}

func (c *Controller) AddRequest(w http.ResponseWriter, r *http.Request) {
	var addTaskRequests []RequestAddTask
	err := json.NewDecoder(r.Body).Decode(&addTaskRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var newTasks []HttpRequestTask
	for _, v := range addTaskRequests {
		newTasks = append(newTasks, HttpRequestTask{Url: v.Url, Count: v.RequestsNumber})
	}

	distinct := make(map[string]HttpRequestTask)
	for _, task := range newTasks {
		if curTask, ok := distinct[task.Url]; ok {
			curTask.Count = task.Count + curTask.Count
			distinct[task.Url] = curTask
		} else {
			distinct[task.Url] = task
		}
	}
	for _, v := range distinct {
		c.taskChannel <- v
	}

	w.WriteHeader(200)
}

func (c *Controller) StatusRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp := ResponseStatus{c.storage.GetTotalRequestCount(), c.storage.GetSuccessRequestCount(), c.storage.GetErrorRequestCount()}
	json.NewEncoder(w).Encode(resp)
}
