package entities

import (
	"time"
)

type VoiceSample struct {
	SessionId int
	Created   time.Time
	Length    time.Duration
	MimeType  string
	Bitrate   int
	AudioURL  string `json:"-"`
}
