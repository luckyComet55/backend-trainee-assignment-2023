package databasetest

type Database[T Identifiable] struct {
	storage map[int]T
}

func NewDatabase[T Identifiable]() Database[T] {
	return Database[T]{
		storage: make(map[int]T),
	}
}

func (d Database[T]) GetObjectById(id int) (T, error) {
	if v, ok := d.storage[id]; !ok {
		return v, ErrObjNotFound{}
	} else {
		return v, nil
	}
}

func (d Database[T]) CreateObject(obj T) error {
	if _, ok := d.storage[obj.GetId()]; ok {
		return ErrObjAlreadyExists{}
	} else {
		d.storage[obj.GetId()] = obj
		return nil
	}
}

func (d Database[T]) UpdateObject(obj T) error {
	if _, ok := d.storage[obj.GetId()]; !ok {
		return ErrObjNotFound{}
	} else {
		d.storage[obj.GetId()] = obj
		// the only possible error here may be network error
		return nil
	}
}

func (d Database[T]) DeleteObject(obj T) error {
	if _, ok := d.storage[obj.GetId()]; !ok {
		return ErrObjNotFound{}
	} else {
		delete(d.storage, obj.GetId())
		return nil
	}
}
