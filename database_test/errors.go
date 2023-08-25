package databasetest

type ErrObjNotFound struct {
}

func (e ErrObjNotFound) Error() string {
	return "object not found"
}

type ErrObjAlreadyExists struct {
}

func (e ErrObjAlreadyExists) Error() string {
	return "object with such id already exists"
}
