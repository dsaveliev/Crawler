package models

import (
	"crawler/errors"
	"net/http"
)

var testCasesGetPage = []struct {
	Link       *Link
	StatusCode int
	URL        string
}{
	{
		&Link{URL: "http://ya.ru", Hostname: "ya.ru", Error: nil},
		http.StatusOK,
		"http://ya.ru",
	},
	{
		&Link{URL: "http:/ya.ru", Hostname: "", Error: errors.INVALID_URL},
		STATUS_INVALID_URL,
		"http:/ya.ru",
	},
	{
		&Link{URL: "http://ya.ru/error", Hostname: "ya.ru", Error: nil},
		STATUS_NETWORK_ERROR,
		"http://ya.ru/error",
	},
}
