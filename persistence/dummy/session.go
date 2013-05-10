// Package dummy implements dummy storage for API entities
package dummy

import (
	"fmt"
	. "github.com/ku-ovdp/api/entities"
	. "github.com/ku-ovdp/api/repository"
	"time"
)

func (d dummyBackend) NewSessionRepository(repositories RepositoryGroup) SessionRepository {
	return sessionRepository{dummySessionData, repositories["projects"].(ProjectRepository)}
}

type sessionRepo map[int]Session

type sessionRepository struct {
	sessionRepo
	projects ProjectRepository
}

var dummySessionData = map[int]Session{
	1: {Id: 1, ProjectId: 1,
		Created: time.Now().Add(time.Hour * -24 * 14),
		FormValues: []FormFieldValue{
			{FieldSlug: "age", Value: 42},
			{FieldSlug: "gender", Value: "Male"},
			{FieldSlug: "parkinsons", Value: true},
		},
		Samples: []VoiceSample{
			{Created: time.Now().Add(time.Hour * -14),
				Length:   time.Second * 10,
				Bitrate:  24000,
				AudioURL: "http://s3.amazon.com/dopebeats.pcm",
			},
		},
	},
}

func (sr sessionRepository) Get(id int) (Session, error) {
	if obj, ok := sr.sessionRepo[id]; ok {
		return obj, nil
	} else {
		return Session{}, NotFound
	}
}

func (sr sessionRepository) Put(session Session) (Session, error) {
	session.Id = len(sr.sessionRepo) + 1
	sr.sessionRepo[session.Id] = session
	return session, nil
}

func (sr sessionRepository) Remove(id int) error {
	delete(sr.sessionRepo, id)
	return nil
}

func (sr sessionRepository) Scan(projectId int, from, to int) ([]Session, error) {
	results := []Session{}
	if _, err := sr.projects.Get(projectId); err != nil {
		if err == NotFound {
			err = fmt.Errorf("Project with id %d not found.", projectId)
		}
		return results, err
	}
	for id, value := range sr.sessionRepo {
		if id < from {
			continue
		}
		if id > to && to != 0 {
			continue
		}
		if value.ProjectId != projectId {
			continue
		}
		results = append(results, value)
	}
	return results, nil
}
