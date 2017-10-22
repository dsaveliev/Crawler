package server

import (
	"crawler/models"
	"crawler/pool"
	"net/http"
)

var testCasesSaveHandler = []struct {
	ReqBody  string
	RespBody string
	Status   int
}{
	{
		"",
		"The URL's list is empty.",
		http.StatusBadRequest,
	},
	{
		"a.com/some/path",
		"1",
		http.StatusCreated,
	},
	{
		"b.com\nc.com\nd.com",
		"2",
		http.StatusCreated,
	},
}

var testCasesViewHandler = []struct {
	Id       string
	RespBody string
	Status   int
	Setup    func()
}{
	{
		"1",
		"Task not found.",
		http.StatusNotFound,
		func() {
			pool.Storage = pool.NewStorage()
		},
	},
	{
		"1",
		"StatusCode	URL	Title	Description	Keywords	OGImage\n",
		http.StatusOK,
		func() {
			pool.Storage = pool.NewStorage()
			pool.Storage.Set(1, &models.Task{Id: 1, State: models.STATE_DONE})
		},
	},
	{
		"2",
		"Task not found.",
		http.StatusNotFound,
		func() {
			pool.Storage = pool.NewStorage()
			pool.Storage.Set(1, &models.Task{Id: 1, State: models.STATE_DONE})
		},
	},
	{
		"1",
		"",
		http.StatusNoContent,
		func() {
			pool.Storage = pool.NewStorage()
			pool.Storage.Set(1, &models.Task{Id: 1, State: models.STATE_IN_PROGRESS})
		},
	},
}

var testCasesDeleteHandler = []struct {
	Id       string
	RespBody string
	Status   int
	Setup    func()
}{
	{
		"1",
		"Task not found.",
		http.StatusNotFound,
		func() {
			pool.Storage = pool.NewStorage()
		},
	},
	{
		"2",
		"Task not found.",
		http.StatusNotFound,
		func() {
			pool.Storage = pool.NewStorage()
			pool.Storage.Set(1, &models.Task{Id: 1, State: models.STATE_DONE})
		},
	},
	{
		"1",
		"",
		http.StatusNoContent,
		func() {
			pool.Storage = pool.NewStorage()
			pool.Storage.Set(1, &models.Task{Id: 1, State: models.STATE_IN_PROGRESS})
		},
	},
	{
		"1",
		"StatusCode	URL	Title	Description	Keywords	OGImage\n",
		http.StatusOK,
		func() {
			pool.Storage = pool.NewStorage()
			pool.Storage.Set(1, &models.Task{Id: 1, State: models.STATE_DONE})
		},
	},
	{
		"1",
		"Task not found.",
		http.StatusNotFound,
		func() {},
	},
}
