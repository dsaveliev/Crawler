package models

import (
	"strconv"
	"strings"
)

const (
	STATE_IN_PROGRESS TaskState = "inProgress"
	STATE_DONE        TaskState = "done"
)

var (
	BODY_DELIMITER = "\n"
	HEADERS        = []string{
		"StatusCode",
		"URL",
		"Title",
		"Description",
		"Keywords",
		"OGImage",
	}
)

type TaskState string

type Task struct {
	Id    uint64
	State TaskState
	Links []*Link
	Pages []*Page
	Done  chan struct{}
}

func NewTask(id uint64, body string) *Task {
	task := &Task{
		Id:    id,
		Links: []*Link{},
		State: STATE_IN_PROGRESS,
		Done:  make(chan struct{}, 1),
	}
	for _, url := range strings.Split(body, BODY_DELIMITER) {
		task.Links = append(task.Links, NewLink(url))
	}
	return task
}

func (t *Task) Complete() {
	t.State = STATE_DONE
	t.Done <- struct{}{}
}

func (t *Task) ToString() string {
	return strconv.FormatUint(t.Id, 10)
}

func (t *Task) ToSlice() (data [][]string) {
	data = [][]string{HEADERS}
	for _, page := range t.Pages {
		data = append(data, page.ToSlice())
	}
	return
}
