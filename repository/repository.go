package repository

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
)

type Repository[T db.Identifiable] struct {
	Db db.Database[T]
}

func NewRepository[T db.Identifiable](db db.Database[T]) Repository[T] {
	return Repository[T]{
		Db: db,
	}
}
