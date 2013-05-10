// Package repository implements the storage interface for API entities
package repository

import (
	. "github.com/ku-ovdp/api/entities"
)

type RepositoryGroup map[string]interface{}

func NewRepositoryGroup() RepositoryGroup {
	return make(RepositoryGroup, 0)
}

type ProjectRepository interface {
	Get(id int) (Project, error)
	Put(project Project) (Project, error)
	Remove(id int) error
	Scan(from, to int) ([]Project, error)
}

type SessionRepository interface {
	Get(id int) (Session, error)
	Put(session Session) (Session, error)
	Remove(id int) error
	Scan(projectId int, from, to int) ([]Session, error)
}
