package database

import (
	"fmt"
)

type ErrObjNotFound struct {
}

func (e ErrObjNotFound) Error() string {
	return "object not found"
}

type ErrObjAlreadyExists struct {
	Id int
}

func (e ErrObjAlreadyExists) Error() string {
	return fmt.Sprintf("object with id %d already exists\n", e.Id)
}

type ErrUniqueConstraintFailed struct {
	Field, Value string
}

func (e ErrUniqueConstraintFailed) Error() string {
	return fmt.Sprintf("field(s) %s value(s) %s\n", e.Field, e.Value)
}

type ErrInternal struct {
}

func (e ErrInternal) Error() string {
	return "internal error"
}
