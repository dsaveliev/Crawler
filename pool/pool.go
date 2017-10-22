package pool

import (
	"crawler/errors"
	"crawler/logger"
	"crawler/models"
	"net/http"
)

var transport = &http.Transport{DisableKeepAlives: true}
var client = &http.Client{Transport: transport}

type ChannelPages chan *models.Page

var (
	Throttle = NewThrottle()
	Storage  = NewStorage()
)

func StartTask(body string) (task *models.Task) {
	id := Storage.NextIndex()
	task = models.NewTask(id, body)
	Storage.Set(id, task)
	go runTask(task)
	return
}

func FindTask(id uint64) (*models.Task, error) {
	task, ok := Storage.Get(id)
	if !ok {
		return task, errors.TASK_NOT_FOUND
	}
	if task.State != models.STATE_DONE {
		return task, errors.TASK_IN_PROGRESS
	}
	return task, nil
}

func DeleteTask(id uint64) (*models.Task, error) {
	task, ok := Storage.Get(id)
	if !ok {
		return task, errors.TASK_NOT_FOUND
	}
	if task.State != models.STATE_DONE {
		return task, errors.TASK_IN_PROGRESS
	}
	Storage.Delete(id)
	return task, nil
}

func runTask(task *models.Task) {
	counter := 0
	chPages := make(ChannelPages, len(task.Links))
	for _, link := range task.Links {
		go getPage(task.Id, link, chPages)
	}
	for counter < len(task.Links) {
		page := <-chPages
		task.Pages = append(task.Pages, page)
		counter++
	}
	task.Complete()
	logger.Debug.Printf("[Task #%d] Completed.\n", task.Id)
}

func getPage(id uint64, link *models.Link, chPages ChannelPages) {
	if link.Error == nil {
		<-Throttle.Get(link.Hostname)
	}
	logger.Debug.Printf("[Task #%d] Get page: %s\n", id, link.URL)
	chPages <- models.GetPage(client, link)
}
