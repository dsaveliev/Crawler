package server

import (
	"crawler/errors"
	"crawler/logger"
	"encoding/csv"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

const CSV_DELIMITER = '\t'

func readBody(ctx *fasthttp.RequestCtx) (string, error) {
	return string(ctx.PostBody()), nil
}

func validateBody(body string) (string, error) {
	body = strings.TrimSpace(body)
	if len(body) == 0 {
		return body, errors.EMPTY_URLS_LIST
	}
	return body, nil
}

func readId(ctx *fasthttp.RequestCtx) (uint64, error) {
	id, err := strconv.ParseUint(ctx.UserValue("id").(string), 10, 64)
	if err != nil {
		return id, errors.On(err, errors.INVALID_TASK_ID)
	}
	return id, nil
}

func encodeCSV(data [][]string, ctx *fasthttp.RequestCtx) error {
	ctx.Response.Header.Set("Content-Type", "text/csv; charset=UTF-8")

	csvWriter := csv.NewWriter(ctx)
	csvWriter.Comma = CSV_DELIMITER

	if err := csvWriter.WriteAll(data); err != nil {
		return errors.On(err, errors.INTERNAL_SERVER_ERROR)
	}

	csvWriter.Flush()

	return nil
}

func onError(err error, ctx *fasthttp.RequestCtx) {
	var statusCode int

	switch parsedError := err.(type) {
	case errors.BaseError:
		statusCode = parsedError.StatusCode
	default:
		statusCode = http.StatusInternalServerError
	}

	logger.Error.Printf("[%d] %s", statusCode, err.Error())

	ctx.SetStatusCode(statusCode)
	io.WriteString(ctx, err.Error())
}
