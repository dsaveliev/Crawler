package server

import (
	"crawler/logger"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func NewRouter() func(*fasthttp.RequestCtx) {
	router := fasthttprouter.New()
	router.POST("/tasks", makeHandler(saveHandler))
	router.GET("/tasks/:id", makeHandler(getHandler))
	return router.Handler
}

func makeHandler(fn func(ctx *fasthttp.RequestCtx)) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		logger.Info.Printf("%s %s\n", ctx.Method(), ctx.Path())
		fn(ctx)
	}
}
