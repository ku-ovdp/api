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

	HighlevelDescription, DetailedDescription      string
	PrivacyPolicyURL                               string
	MinimumNumberOfSamples int
	MaximumNumberOfSamples int
	GeneralInstructions                            string
	SampleInstructions                             []SampleInstruction
	Meta                                           string
}

/*
class Project(models.Model):
    name = models.CharField(max_length=255)
    slug = models.SlugField()
    created = models.DateField(auto_now_add=True)
    high_level_description = models.TextField()
    detailed_description = models.TextField()
    privacy_policy_url = models.URLField()
    minimum_number_of_samples = models.IntegerField(null=True)
    maximum_number_of_samples = models.IntegerField(null=True)
    general_instructions = models.TextField()
    sample_instructions = JSONField(blank=True)
        # schema:
        # [{'duration': (integer),
        #   'instruction': (string)}*]
    meta = JSONField(blank=True)*/
