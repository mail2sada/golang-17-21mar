package main

import (
	"os"
)

const (
	INFO  = 0
	DEBUG = 1
	ERROR = 2
)

type Logger struct {
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	ErrorLogger *log.Logger
}

var logger Logger
var logLevel = ERROR

func init() {
	f, _ := os.Create("log.file")
	logger.InfoLogger = log.New(f, "[INFO]", log.Ldate|log.Lshortfile)
	logger.DebugLogger = log.New(f, "[DEBUG]", log.Ldate|log.Lshortfile)
	logger.ErrorLogger = log.New(f, "[ERROR]", log.Ldate|log.Lshortfile)
}
func log(logline string, level int) {
	if level >= logLevel {
		switch level {
		case INFO:
			logger.InfoLogger(logline)
		case ERROR:
			logger.ErrorLogger(logline)
		case DEBUG:
			logger.DebugLogger(logline)
		}
	}
}
func main() {
	log("printingh", ERROR)
	log("ajfajkdfha", INFO)
}
