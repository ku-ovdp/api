package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"os"
)

const API_VERSION = 1

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	listen()
}

func listen() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Println("attempting to listen on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(req *restful.Request, resp *restful.Response) {
	resp.Write([]byte("world"))
}
