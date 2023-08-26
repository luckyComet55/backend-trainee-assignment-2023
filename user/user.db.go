package user

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
)

type UserDatabase interface {
	db.Database[User]
}
