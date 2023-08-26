package database

type Database[T any] interface {
	GetObjectById(id int) (T, error)
	CreateObject(obj T) error
	UpdateObject(obj T) error
	DeleteObject(obj T) error
}
