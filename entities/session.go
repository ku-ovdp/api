package entities

import (
	"time"
)

type Stringer interface {
	String() string
}

type Session struct {
	Id               int
	ProjectId        int
	RelatedSessionId int
	Created          time.Time
	Completed        bool
	UserAgent        string
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
