// Package sessions supplies an interface to the Session entity
package sessions

import (
	"github.com/ku-ovdp/api/entities"
	"github.com/traviscline/go-restful"
	"net/http"
	"strconv"
	"time"
)

type sessionService struct {
	*restful.WebService
	repository entities.SessionRepository
}

func NewSessionService(apiRoot string, repository entities.SessionRepository) *sessionService {
	s := new(sessionService)
	ws := new(restful.WebService)

	ws.Path(apiRoot + "/project/{project-id}/sessions").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").To(s.listSessions).
		Doc("List sessions").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("int")).
		Writes([]entities.Session{}))

	ws.Route(ws.GET("/{session-id}").To(s.findSession).
		Doc("Get a session").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("int")).
		Param(ws.PathParameter("session-id", "identifier of the session").DataType("int")).
		Writes(entities.Session{}))

	ws.Route(ws.POST("").To(s.createSession).
		Doc("Create a session").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("int")).
		Reads(entities.Session{}))

	ws.Route(ws.PUT("/{session-id}").To(s.updateSession).
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("int")).
		Param(ws.PathParameter("session-id", "identifier of the session").DataType("int")).
		Doc("Update a session").
		Param(ws.PathParameter("session-id", "identifier of the session")))

	ws.Route(ws.DELETE("/{session-id}").To(s.removeSession).
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("int")).
		Param(ws.PathParameter("session-id", "identifier of the session").DataType("int")).
		Doc("Delete a session").
		Param(ws.PathParameter("session-id", "identifier of the session")))

	s.WebService = ws
	s.repository = repository
	return s
}

func (s *sessionService) listSessions(request *restful.Request, response *restful.Response) {
	projectId, err := strconv.Atoi(request.PathParameter("project-id"))
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	if sessions, err := s.repository.Scan(projectId, 0, 0); err == nil {
		response.WriteEntity(sessions)
	} else {
		response.WriteError(http.StatusBadRequest, err)
	}
}

func (s *sessionService) findSession(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("session-id"))
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}
	session, _ := s.repository.Get(id)

	if session.Id == 0 {
		response.WriteError(http.StatusNotFound, nil)
	} else {
		response.WriteEntity(session)
	}
}

func (s *sessionService) updateSession(request *restful.Request, response *restful.Response) {
	session := new(entities.Session)
	err := request.ReadEntity(&session)
	if err == nil {
		s.repository.Put(*session)
		response.WriteEntity(session)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (s *sessionService) createSession(request *restful.Request, response *restful.Response) {
	projectId, err := strconv.Atoi(request.PathParameter("project-id"))
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	session := entities.Session{
		ProjectId: projectId,
		Created:   time.Now(),
	}
	session, err = s.repository.Put(session)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(session)
}

func (s *sessionService) removeSession(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("session-id"))
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	s.repository.Remove(id)
}
