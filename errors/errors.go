package errors

import (
	"crawler/logger"
	"net/http"
)

var EMPTY_URLS_LIST = BaseError{
	http.StatusBadRequest,
	"The URL's list is empty.",
}

var INVALID_URL = BaseError{
	http.StatusBadRequest,
	"Invalid URL.",
}

var INVALID_TASK_ID = BaseError{
	http.StatusBadRequest,
	"Invalid task id.",
}

var TASK_IN_PROGRESS = BaseError{
	http.StatusNoContent,
	"Task in progress.",
}

var TASK_NOT_FOUND = BaseError{
	http.StatusNotFound,
	"Task not found.",
}

var INTERNAL_SERVER_ERROR = BaseError{
	http.StatusInternalServerError,
	"Something went wrong.",
}

type BaseError struct {
	StatusCode int
	Message    string
}

func (e BaseError) Error() string {
	return e.Message
}

func On(err error, berr BaseError) error {
	logger.Error.Println(err.Error())
	if err != nil {
		return berr
	} else {
		return nil
	}
}
