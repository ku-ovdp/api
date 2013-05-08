package main

import (
	"fmt"
	"github.com/traviscline/go-restful"
	"github.com/ku-ovdp/api/persistence/dummy"
	"github.com/ku-ovdp/api/projects"
	"github.com/ku-ovdp/api/repository"
	"github.com/ku-ovdp/api/sessions"
	"net/http"
	"time"
)

// Create application services and dependancies
func constructApplication() {
	repositories := repository.NewRepositoryGroup()

	// construct dummy Project repository
	projectRepository := dummy.NewProjectRepository(repositories)
	repositories["projects"] = projectRepository

	// construct dummy Session repository
	sessionRepository := dummy.NewSessionRepository(repositories)
	repositories["sessions"] = sessionRepository

	apiRoot := fmt.Sprintf("/v%d", API_VERSION)

	restful.Dispatch = func(w http.ResponseWriter, r *http.Request) {
		lwr := &loggedResponseWriter{w, 0}
		t1 := time.Now()
		restful.DefaultDispatch(lwr, r)
		fmt.Println(r.Method, r.URL, lwr.status, time.Now().Sub(t1))
	}
	restful.DefaultResponseMimeType = restful.MIME_JSON
	restful.Add(projects.NewProjectService(apiRoot, projectRepository))
	restful.Add(sessions.NewSessionService(apiRoot, sessionRepository))
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
		}
		resp.WriteAsJson(apiListing)
	}))
	ws.Route(ws.GET("/favicon.ico").To(func(req *restful.Request, resp *restful.Response) {
		resp.WriteHeader(http.StatusNotFound)
	}))
	return ws
}

type loggedResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *loggedResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *loggedResponseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}
