package main

import (
	"flag"
	"log"
	"os"

	spoe "github.com/criteo/haproxy-spoe-go"
)

/* global variables */
var Debug   *log.Logger
var Info    *log.Logger
var Warning *log.Logger
var Error   *log.Logger

var debug bool

func init() () {
	flag.BoolVar(&debug, "debug", false, "turn on debug messages")
	flag.Parse()

	Debug   = log.New(os.Stdout, "DEBUG:", log.Ldate|log.Ltime|log.Lshortfile)
	Info    = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stderr, "WARNING:", log.Ldate|log.Ltime|log.Lshortfile)
	Error   = log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() () {
	var agent *spoe.Agent
	var err error

	agent = spoe.New(doResolve)

	err = agent.ListenAndServe(":9000")
	if err != nil {
		Error.Fatal(err)
	}

	return
}
