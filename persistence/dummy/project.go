// Package dummy implements dummy storage for API entities
package dummy

import (
	"github.com/ku-ovdp/api/entities"
	"time"
)

type projectRepository map[int]entities.Project

var dummyData = map[int]entities.Project{
	1: {Id: 1, Name: "Project Parkinson's",
		Slug:                   "parkinsons",
		HighlevelDescription:   "Project Parkinson's ... (high level)",
		DetailedDescription:    "Project Parkinson's ... (detailed)",
		PrivacyPolicyURL:       "http://openvoicedata.org/privacy.php",
		MinimumNumberOfSamples: 2,
		MaximumNumberOfSamples: 3,
		GeneralInstructions:    "(general instructions)",
		SampleInstructions: []entities.SampleInstruction{
			{Duration: 10, Instruction: "Produce an 'Ah' sound at a comfortable level."},
			{Duration: 10, Instruction: "Produce an 'Ah' sound with twice the previous effort."},
			{Duration: 0, Instruction: "Produce normal conversational speaking"},
		},
		FormFields: []entities.FormField{
			{Label: "Age", Slug: "age", Type: "int", Required: true, Description: "Your Age"},
			{Label: "Gender", Slug: "gender", Type: "string", Required: true, Description: "Your Gender",
				Meta: `{"options": ["Male", "Female", "Undisclosed"]}`},
		},
		Created: time.Now().Add(time.Hour * -24 * 14)},

	2: {Id: 2, Name: "Disphonia Foobar",
		Slug:    "foobar",
		Created: time.Now().Add(time.Hour * -24 * 10)},
}

func NewProjectRepository() projectRepository {
	return dummyData
}

func (pr projectRepository) Get(id int) entities.Project {
	return pr[id]
}

func (pr projectRepository) Put(project entities.Project) {
	pr[project.Id] = project
}

func (pr projectRepository) Remove(id int) error {
	delete(pr, id)
	return nil
}

func (pr projectRepository) Scan(from, to int) []entities.Project {
	results := make([]entities.Project, 0)
	for id, value := range pr {
		if id < from {
			continue
		}
		if id > to && to != 0 {
			continue
		}
		results = append(results, value)
	}
	return results
}
