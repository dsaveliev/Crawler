package server

import (
	"crawler/pool"
	"io"
	"net/http"

	"github.com/valyala/fasthttp"
)

func saveHandler(ctx *fasthttp.RequestCtx) {
	body, err := readBody(ctx)
	if err != nil {
		onError(err, ctx)
		return
	}

	body, err = validateBody(body)
	if err != nil {
		onError(err, ctx)
		return
	}

	task := pool.StartTask(body)

	ctx.SetStatusCode(http.StatusCreated)
	io.WriteString(ctx, task.ToString())
}

func getHandler(ctx *fasthttp.RequestCtx) {
	delete := ctx.QueryArgs().GetBool("delete")
	if delete {
		deleteHandler(ctx)
	} else {
		viewHandler(ctx)
	}
}

func viewHandler(ctx *fasthttp.RequestCtx) {
	id, err := readId(ctx)
	if err != nil {
		onError(err, ctx)
		return
	}

	task, err := pool.FindTask(id)
	if err != nil {
		onError(err, ctx)
		return
	}

	err = encodeCSV(task.ToSlice(), ctx)
	if err != nil {
		onError(err, ctx)
		return
	}
}

func deleteHandler(ctx *fasthttp.RequestCtx) {
	id, err := readId(ctx)
	if err != nil {
		onError(err, ctx)
		return
	}

	task, err := pool.DeleteTask(id)
	if err != nil {
		onError(err, ctx)
		return
	}

	err = encodeCSV(task.ToSlice(), ctx)
	if err != nil {
		onError(err, ctx)
		return
	}
}
