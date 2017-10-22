package server

import (
	"crawler/logger"

	"github.com/valyala/fasthttp"
)

func Run(port string, debug bool) {
	if debug {
		logger.EnableDebug()
		logger.Debug.Println("Debug mode enabled.")
	}
	listenString := "localhost:" + port
	logger.Info.Println("Serving http://" + listenString)
	logger.Error.Fatal(fasthttp.ListenAndServe(listenString, NewRouter()))
}
