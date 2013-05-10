package main

import (
	"encoding/json"
	"fmt"
	"github.com/ku-ovdp/api/persistence"
	_ "github.com/ku-ovdp/api/persistence/dummy"
	_ "github.com/ku-ovdp/api/persistence/mgo"
	"github.com/ku-ovdp/api/projects"
	"github.com/ku-ovdp/api/repository"
	"github.com/ku-ovdp/api/sessions"
	"github.com/ku-ovdp/api/sockgroup"
	"github.com/ku-ovdp/api/stats"
	"github.com/traviscline/go-restful"
	"net/http"
	"time"
)

// Create application services and dependancies
func constructApplication() {

	sg := sockgroup.NewGroup()
	sg.Start()
	stats.Destination = sg

	// construct repositories

	backend := persistence.Get(*persistenceBackend)
	if backend == nil {
		logger.Fatalln("Invalid repository backend.", *persistenceBackend)
	}
	backend.Init()

	repositories := repository.NewRepositoryGroup()
	projectRepository := backend.NewProjectRepository(repositories)
	repositories["projects"] = projectRepository
	sessionRepository := backend.NewSessionRepository(repositories)
	repositories["sessions"] = sessionRepository

	restful.Dispatch = func(w http.ResponseWriter, r *http.Request) {
		lwr := &loggedResponseWriter{w, 0}
		t1 := time.Now()
		restful.DefaultDispatch(lwr, r)
		fmt.Println(r.Method, r.URL, lwr.status, time.Now().Sub(t1))
	}
	restful.DefaultResponseMimeType = restful.MIME_JSON

	apiRoot := fmt.Sprintf("/v%d", API_VERSION)
	restful.Add(projects.NewProjectService(apiRoot, projectRepository))
	restful.Add(sessions.NewSessionService(apiRoot, sessionRepository))

	http.Handle("/stats/", sg.Handler("/stats"))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", indexHandler(apiRoot))
}

func indexHandler(apiRoot string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		b, _ := json.MarshalIndent([]struct {
			Url string `json:"endpoint_url"`
		}{
			{apiRoot + "/projects"},
		}, "", "  ")
		fmt.Fprintln(w, string(b))
	}
}

type loggedResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *loggedResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
	stats.ChangeStat("requests", 1)
}

func (w *loggedResponseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}
