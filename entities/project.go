package entities

import (
	"time"
)

type ProjectRepository interface {
	Get(id int) Project
	Put(project Project)
	Remove(id int) error
	Scan(from, to int) []Project
}

type SampleInstruction struct {
	Duration    int
	Instruction string
}

type Project struct {
	Id      int
	Name    string
	Slug    string
	Created time.Time

	HighlevelDescription   string
	DetailedDescription    string
	PrivacyPolicyURL       string
	MinimumNumberOfSamples int
	MaximumNumberOfSamples int
	GeneralInstructions    string
	SampleInstructions     []SampleInstruction
	FormFields             []FormField
	Meta                   string
}

type FieldType string

type FormField struct {
	Label        string
	Slug         string
	Type         FieldType
	Required     bool
	Group, Order int
	Placeholder  string
	Description  string
	Meta         string
}
