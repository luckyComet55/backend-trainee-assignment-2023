package databasetest

type Identifiable interface {
	GetId() int
}

type Database[T any] interface {
	GetObjectById(id int) (T, error)
	GetObjectByName(name string) (T, error)
	CreateObject(obj T) error
	UpdateObject(obj T) error
	DeleteObject(obj T) error
}
