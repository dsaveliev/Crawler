package main

import (
	"crawler/server"
	"flag"
)

var (
	port  = flag.String("port", "8080", "Define the TCP port to bind to.")
	debug = flag.Bool("debug", false, "Enable detailed logging.")
)

func main() {
	flag.Parse()
	server.Run(*port, *debug)
}
