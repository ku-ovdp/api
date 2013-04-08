// Package repository implements the storage interface for API entities
package repository

type RepositoryGroup map[string]interface{}

func NewRepositoryGroup() RepositoryGroup {
	return make(RepositoryGroup, 0)
}

type Repository interface {
	Get(id int) interface{}
	Put(entity interface{})
	Remove(id int) error
}
