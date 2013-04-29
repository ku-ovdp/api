package entities

type ProjectRepository interface {
	Get(id int) Project
	Put(project Project)
	Remove(id int) error
	Scan(from, to int) []Project
}

type Project struct {
	Id   int
	Name string
}
