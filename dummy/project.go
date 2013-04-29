// Package dummy implements dummy storage for API entities
package dummy

import (
	"github.com/ku-ovdp/api/entities"
)

type projectRepository map[int]entities.Project

func NewProjectRepository() projectRepository {
	return map[int]entities.Project{
		1: {1, "Project Parkinson's"},
		2: {2, "Disphonia Foo"},
		3: {3, "Disphonia Bar"},
		4: {4, "Disphonia Baz"},
	}
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
