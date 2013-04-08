// Package api implements the REST API for api.openvoicedata.org
package main

import (
	"flag"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/inhies/go-utils/log"
	"github.com/ku-ovdp/api/dummy"
	"github.com/ku-ovdp/api/projects"
	"github.com/ku-ovdp/api/repository"
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

func main() {
	flag.Parse()
	setupLogger()

	repositories := repository.NewRepositoryGroup()
	projectRepository := dummy.NewProjectRepository()

	repositories["projects"] = projectRepository
	fmt.Printf("%#v\n", repositories)

	apiRoot := fmt.Sprintf("/v%d", API_VERSION)
	restful.Add(projects.NewProjectService(apiRoot, projectRepository))
	restful.Add(indexHandler(apiRoot))
	//restful.Add(NewSessionService())
	//restful.Add(NewVoiceSampleService())

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
