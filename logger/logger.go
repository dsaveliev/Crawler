package logger

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug = log.New(ioutil.Discard, "[DEBUG] ", log.Ldate|log.Ltime)
	Info  = log.New(os.Stdout, "[ INFO] ", log.Ldate|log.Ltime)
	Error = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime)
)

func EnableDebug() {
	Debug = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime)
}

func Mute() {
	for _, l := range []*log.Logger{Debug, Info, Error} {
		l.SetOutput(ioutil.Discard)
	}
}
