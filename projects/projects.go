// Package projects supplies an interface to the Project entity
package projects

import (
	"github.com/traviscline/go-restful"
	"github.com/ku-ovdp/api/entities"
	"net/http"
	"strconv"
)

type projectService struct {
	*restful.WebService
	repository entities.ProjectRepository
}

func NewProjectService(apiRoot string, repository entities.ProjectRepository) *projectService {
	ps := new(projectService)
	ws := new(restful.WebService)

	ws.Path(apiRoot + "/projects").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").To(ps.listProjects).
		Doc("List projects").
		Writes([]entities.Project{}))

	ws.Route(ws.GET("/{project-id}").To(ps.findProject).
		Doc("Get a project").
		Param(ws.PathParameter("project-id", "identifier of the project")).
		Writes(entities.Project{}))

	ws.Route(ws.POST("").To(ps.createProject).
		Doc("Create a project").
		Reads(entities.Project{}))

	ws.Route(ws.PUT("/{project-id}").To(ps.updateProject).
		Doc("Update a project").
		Param(ws.PathParameter("project-id", "identifier of the project")))

	ws.Route(ws.DELETE("/{project-id}").To(ps.removeProject).
		Doc("Delete a project").
		Param(ws.PathParameter("project-id", "identifier of the project")))

	ps.WebService = ws
	ps.repository = repository
	return ps
}

func (ps *projectService) listProjects(request *restful.Request, response *restful.Response) {
	if projects, err := ps.repository.Scan(0, 0); err == nil {
		response.WriteEntity(projects)
	} else {
		response.WriteError(http.StatusBadRequest, err)
	}
}

func (ps *projectService) findProject(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("project-id"))
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}
	project, err := ps.repository.Get(id)

	if project.Id == 0 {
		response.WriteError(http.StatusNotFound, nil)
	} else {
		response.WriteEntity(project)
	}
}

func (ps *projectService) updateProject(request *restful.Request, response *restful.Response) {
	project := new(entities.Project)
	err := request.ReadEntity(&project)
	if err == nil {
		ps.repository.Put(*project)
		response.WriteEntity(project)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (ps *projectService) createProject(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("project-id"))
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	project := entities.Project{Id: id}
	err = request.ReadEntity(&project)
	if err == nil {
		ps.repository.Put(project)
		response.WriteHeader(http.StatusCreated)
		response.WriteEntity(project)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (ps *projectService) removeProject(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("project-id"))
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	ps.repository.Remove(id)
}
