package entities

import (
	"time"
)

type Stringer interface {
	String() string
}

type SessionRepository interface {
	Get(id int) (Session, error)
	Put(session Session) (Session, error)
	Remove(id int) error
	Scan(projectId int, from, to int) ([]Session, error)
}

type Session struct {
	Id               int
	ProjectId        int
	RelatedSessionId int
	Created          time.Time
	Completed        bool
	FormValues       []FormFieldValue
	Samples          []VoiceSample
}

type VoiceSample struct {
	Created  time.Time
	Length   time.Duration
	Bitrate  int
	AudioURL string
}

type FormFieldValue struct {
	FieldSlug string
	Value     interface{}
}
