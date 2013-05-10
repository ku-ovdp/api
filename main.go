// Package api implements the REST API for api.openvoicedata.org
package main

import (
	"flag"
	"github.com/inhies/go-utils/log"
	_ "log"
	"net/http"
	"os"
)

// Constants
const API_VERSION = 1

// Globals
var logger *log.Logger

// Flags
var logLevel = flag.String("logLevel", "info", "Log level (warn, info, debug)")
var persistenceBackend = flag.String("persistenceBackend", "mgo", "Persistence backend (dummy, mgo)")

func main() {
	flag.Parse()
	setupLogger()

	constructApplication()

	listen()
}

func listen() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	logger.Infoln("attempting to listen on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Critln("ListenAndServe:", err)
		os.Exit(1)
	}
}

func setupLogger() {
	var (
		err error
		ll  log.LogLevel
	)
	if ll, err = log.ParseLevel(*logLevel); err != nil {
		ll, _ = log.ParseLevel("debug")
	}
	logger, err = log.NewLevel(ll, true, os.Stdout, "", 0)
	logger.Debugln("Configured logger with logLevel", ll.String())
}
