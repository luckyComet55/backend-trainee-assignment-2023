package user

import (
	db_ "github.com/luckyComet55/backend-trainee-assignment-2023/database"
)

type UserDatabase interface {
	db_.Database[User]
	GetUserById(int) (User, error)
	GetRandomUsersByPercent(int) []User
}
