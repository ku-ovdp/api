package main

import (
	"github.com/emicklei/go-restful"
	"net/http"
)

// Redirects / requests to the url provided
func indexHandler(apiRoot string) *restful.WebService {
	ws := new(restful.WebService)
	ws.Produces(restful.MIME_JSON).Consumes(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(func(req *restful.Request, resp *restful.Response) {
		http.Redirect(resp.ResponseWriter, req.Request, apiRoot, http.StatusFound)
	}))
	ws.Route(ws.GET(apiRoot).To(func(req *restful.Request, resp *restful.Response) {
		apiListing := []struct {
			Url string `json:"endpoint_url"`
		}{
			{apiRoot + "/projects"},
			{apiRoot + "/sessions"},
			{apiRoot + "/voice_samples"},
		}
		resp.WriteAsJson(apiListing)
	}))
	ws.Route(ws.GET("/favicon.ico").To(func(req *restful.Request, resp *restful.Response) {
		resp.WriteHeader(http.StatusNotFound)
	}))
	return ws
}
