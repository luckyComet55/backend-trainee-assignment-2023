package user

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
)

type UserMockDatabase struct {
	storage map[int]User
}

func NewUserMockDatabase() *UserMockDatabase {
	return &UserMockDatabase{
		storage: make(map[int]User),
	}
}

func (d *UserMockDatabase) GetObjectById(id int) (User, error) {
	if v, ok := d.storage[id]; !ok {
		return v, db.ErrObjNotFound{}
	} else {
		return v, nil
	}
}

func (d *UserMockDatabase) GetObjectByName(name string) (User, error) {
	return User{}, db.ErrUnsupportedMethod{}
}

func (d *UserMockDatabase) CreateObject(user User) error {
	if _, ok := d.storage[user.GetId()]; ok {
		return db.ErrObjAlreadyExists{Id: user.GetId()}
	}
	d.storage[user.GetId()] = user
	return nil
}

func (d *UserMockDatabase) UpdateObject(user User) error {
	if _, ok := d.storage[user.GetId()]; !ok {
		return db.ErrObjNotFound{}
	}
	d.storage[user.GetId()] = user
	return nil
}

func (d *UserMockDatabase) DeleteObject(user User) error {
	if _, ok := d.storage[user.GetId()]; !ok {
		return db.ErrObjNotFound{}
	}
	delete(d.storage, user.GetId())
	return nil
}
