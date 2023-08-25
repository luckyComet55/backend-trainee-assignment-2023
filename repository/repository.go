package repository

import "github.com/luckyComet55/backend-trainee-assignment-2023/databasetest"

type Repository[T databasetest.Identifiable] struct {
	Db Database[T]
}

func NewRepository[T databasetest.Identifiable](db Database[T]) Repository[T] {
	return Repository[T]{
		Db: db,
	}
}
