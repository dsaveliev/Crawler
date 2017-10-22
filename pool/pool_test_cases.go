package pool

import (
	"crawler/errors"
	"crawler/models"
)

var testCasesStartTask = []struct {
	Body         string
	Id           uint64
	StorageSize  int
	ThrottleSize int
}{
	{
		"http://a.com",
		1,
		1,
		1,
	},
	{
		"http://b.com\nhttp://c.com",
		2,
		2,
		3,
	},
	{
		"http:/d.com\nhttp://e.com",
		3,
		3,
		4,
	},
}

var testCasesFindTask = []struct {
	Id    uint64
	Error error
	Setup func()
}{
	{
		1,
		errors.TASK_NOT_FOUND,
		func() {
			Storage = NewStorage()
		},
	},
	{
		1,
		errors.TASK_IN_PROGRESS,
		func() {
			Storage = NewStorage()
			Storage.Set(1, &models.Task{Id: 1, State: models.STATE_IN_PROGRESS})
		},
	},
	{
		1,
		nil,
		func() {
			Storage = NewStorage()
			Storage.Set(1, &models.Task{Id: 1, State: models.STATE_DONE})
		},
	},
}

var testCasesDeleteTask = []struct {
	Id          uint64
	Error       error
	StorageSize int
	Setup       func()
}{
	{
		1,
		errors.TASK_NOT_FOUND,
		0,
		func() {
			Storage = NewStorage()
		},
	},
	{
		1,
		errors.TASK_IN_PROGRESS,
		1,
		func() {
			Storage = NewStorage()
			Storage.Set(1, &models.Task{Id: 1, State: models.STATE_IN_PROGRESS})
		},
	},
	{
		1,
		nil,
		0,
		func() {
			Storage = NewStorage()
			Storage.Set(1, &models.Task{Id: 1, State: models.STATE_DONE})
		},
	},
}
