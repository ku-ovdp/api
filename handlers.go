package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/ku-ovdp/api/dummy"
	"github.com/ku-ovdp/api/projects"
	"github.com/ku-ovdp/api/repository"
	"net/http"
)

// Create application services and dependancies
func constructApplication() {
	repositories := repository.NewRepositoryGroup()
	projectRepository := dummy.NewProjectRepository()
	repositories["projects"] = projectRepository

	apiRoot := fmt.Sprintf("/v%d", API_VERSION)
	restful.Add(projects.NewProjectService(apiRoot, projectRepository))
	restful.Add(indexHandler(apiRoot))
}

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
