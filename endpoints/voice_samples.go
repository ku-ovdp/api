package endpoints

import (
	. "github.com/ku-ovdp/api/entities"
	. "github.com/ku-ovdp/api/repository"
	"github.com/ku-ovdp/api/stats"
	"github.com/traviscline/go-restful"
	"net/http"
	"strconv"
	"time"
)

type sampleService struct {
	*restful.WebService
	repository VoiceSampleRepository
}

func NewVoiceSampleService(apiRoot string, repository VoiceSampleRepository) *sampleService {
	s := new(sampleService)
	ws := new(restful.WebService)

	ws.Path(apiRoot + "/session/{session-id}/samples").
		Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").To(s.listVoiceSamples).
		Doc("List voice samples").
		//Param(ws.PathParameter("session-id", "identifier of the session").DataType("int")).
		Param(ws.QueryParameter("from", "minimum identifier of a project")).
		Param(ws.QueryParameter("to", "maximum identifier of a project")).
		Writes([]VoiceSample{}))

	ws.Route(ws.POST("").To(s.createVoiceSample).
		Doc("Create a voice sample").
		Param(ws.PathParameter("session-id", "identifier of the session").DataType("int")).
		Reads(VoiceSample{}))

	ws.Route(ws.GET("/{sample-index}").To(s.findVoiceSample).
		Doc("Get a voice sample").
		Param(ws.PathParameter("session-id", "identifier of the session").DataType("int")).
		Writes(VoiceSample{}))

	ws.Route(ws.PUT("/{sample-index}").To(s.updateVoiceSample).
		Doc("Update a voice sample").
		Param(ws.PathParameter("session-id", "identifier of the session").DataType("int")).
		Param(ws.PathParameter("sample-index", "identifier of the sample").DataType("int")).
		Param(ws.BodyParameter("VoiceSample", "the session entity").DataType("string")))

	ws.Route(ws.DELETE("/{sample-index}").To(s.removeVoiceSample).
		Doc("Delete a voice sample").
		Param(ws.PathParameter("session-id", "identifier of the session").DataType("int"))).
		Param(ws.PathParameter("sample-index", "identifier of the sample").DataType("int"))

	ws.Route(ws.GET("/{sample-index}/audio").To(s.streamVoiceSample).
		Doc("Get a voice sample's audio").
		Param(ws.PathParameter("session-id", "identifier of the session").DataType("int")).
		Param(ws.PathParameter("sample-index", "identifier of the sample").DataType("int")).
		Writes(VoiceSample{}))

	ws.Route(ws.PUT("/{sample-index}/audio").To(s.uploadVoiceSample).
		Doc("Attach audio to a voice sample").
		Param(ws.PathParameter("session-id", "identifier of the session").DataType("int")).
		Param(ws.BodyParameter("Audio", "the audio blob entity").DataType("string")))

	s.WebService = ws
	s.repository = repository

	return s
}

func (s *sampleService) listVoiceSamples(request *restful.Request, response *restful.Response) {
	sessionId, _ := strconv.Atoi(request.PathParameter("session-id"))

	from, _ := strconv.Atoi(request.QueryParameter("from"))
	to, _ := strconv.Atoi(request.QueryParameter("to"))

	if samples, err := s.repository.Scan(sessionId, from, to); err == nil {
		response.WriteEntity(samples)
	} else {
		response.WriteError(http.StatusBadRequest, err)
	}
}

func (s *sampleService) findVoiceSample(request *restful.Request, response *restful.Response) {
	sessionId, _ := strconv.Atoi(request.PathParameter("session-id"))
	sampleId, _ := strconv.Atoi(request.PathParameter("sample-index"))

	sample, _ := s.repository.Get(sessionId, sampleId)

	if sample.Id == 0 {
		response.WriteError(http.StatusNotFound, nil)
	} else {
		response.WriteEntity(sample)
	}
}

func (s *sampleService) createVoiceSample(request *restful.Request, response *restful.Response) {
	sessionId, _ := strconv.Atoi(request.PathParameter("session-id"))

	sample := VoiceSample{
		SessionId: sessionId,
		Created:   time.Now(),
	}
	var err error
	sample, err = s.repository.Put(sample)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	stats.ChangeStat("samples", 1)
	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(sample)
}

func (s *sampleService) updateVoiceSample(request *restful.Request, response *restful.Response) {
	sample := new(VoiceSample)
	err := request.ReadEntity(&sample)
	if err == nil {
		s.repository.Put(*sample)
		response.WriteEntity(sample)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (s *sampleService) removeVoiceSample(request *restful.Request, response *restful.Response) {
	sessionId, _ := strconv.Atoi(request.PathParameter("session-id"))
	sampleId, _ := strconv.Atoi(request.PathParameter("sample-index"))

	err := s.repository.Remove(sessionId, sampleId)
	if err == nil {
		response.WriteEntity("removed")
	} else {
		response.WriteError(http.StatusBadRequest, err)
	}
}

func (s *sampleService) streamVoiceSample(request *restful.Request, response *restful.Response) {
	// stream from s3
}

func (s *sampleService) uploadVoiceSample(request *restful.Request, response *restful.Response) {
	// upload to s3
}
